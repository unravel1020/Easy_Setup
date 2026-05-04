package agent

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/unravel1020/Easy_Setup/internal/catalog"
)

var ErrJobNotFound = errors.New("job not found")

type Launcher func(scriptPath string) error

type Job struct {
	ID         string           `json:"id"`
	Status     string           `json:"status"`
	ScriptPath string           `json:"scriptPath"`
	LogPath    string           `json:"logPath"`
	CreatedAt  time.Time        `json:"createdAt"`
	ItemIDs    []string         `json:"itemIds"`
	Actions    []catalog.Action `json:"actions"`
}

func DefaultLauncher(scriptPath string) error {
	if runtime.GOOS == "windows" {
		return exec.Command("powershell", "-NoExit", "-ExecutionPolicy", "Bypass", "-File", scriptPath).Start()
	}
	return exec.Command("sh", scriptPath).Start()
}

func CreateJob(plan *catalog.Plan, jobDir string, launch Launcher) (*Job, error) {
	jobDir = defaultJobDir(jobDir)
	if launch == nil {
		launch = DefaultLauncher
	}
	if err := os.MkdirAll(jobDir, 0o755); err != nil {
		return nil, err
	}

	id, err := newJobID()
	if err != nil {
		return nil, err
	}

	scriptPath := filepath.Join(jobDir, id+".ps1")
	logPath := filepath.Join(jobDir, id+".log")
	jobPath := filepath.Join(jobDir, id+".json")

	job := &Job{
		ID:         id,
		Status:     "created",
		ScriptPath: scriptPath,
		LogPath:    logPath,
		CreatedAt:  time.Now().UTC(),
		ItemIDs:    append([]string(nil), plan.ItemIDs...),
		Actions:    append([]catalog.Action(nil), plan.Actions...),
	}

	if err := os.WriteFile(scriptPath, []byte(renderPowerShellJob(job)), 0o644); err != nil {
		return nil, err
	}
	if err := writeJob(jobPath, job); err != nil {
		return nil, err
	}
	if err := launch(scriptPath); err != nil {
		job.Status = "launch-failed"
		_ = writeJob(jobPath, job)
		return job, err
	}
	job.Status = "launched"
	if err := writeJob(jobPath, job); err != nil {
		return nil, err
	}
	return job, nil
}

func renderPowerShellJob(job *Job) string {
	lines := []string{
		`$ErrorActionPreference = "Continue"`,
		fmt.Sprintf(`Start-Transcript -Path "%s" -Force`, escapePowerShellString(job.LogPath)),
		fmt.Sprintf(`Write-Host "Easy_Setup job %s"`, job.ID),
	}
	for _, action := range job.Actions {
		lines = append(lines,
			fmt.Sprintf(`Write-Host ""`),
			fmt.Sprintf(`Write-Host "Running: %s"`, escapePowerShellString(action.ID)),
			action.Command,
		)
		for _, verify := range action.Verify {
			lines = append(lines,
				fmt.Sprintf(`Write-Host "Verify: %s"`, escapePowerShellString(verify)),
				verify,
			)
		}
	}
	lines = append(lines, `Stop-Transcript`)
	return strings.Join(lines, "\r\n") + "\r\n"
}

func writeJob(path string, job *Job) error {
	data, err := json.MarshalIndent(job, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func ReadJob(jobDir string, id string) (*Job, error) {
	if id == "" {
		return nil, ErrJobNotFound
	}
	jobPath := filepath.Join(defaultJobDir(jobDir), id+".json")
	data, err := os.ReadFile(jobPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrJobNotFound
		}
		return nil, err
	}
	var job Job
	if err := json.Unmarshal(trimUTF8BOM(data), &job); err != nil {
		return nil, err
	}
	return &job, nil
}

func ListJobs(jobDir string) ([]Job, error) {
	entries, err := os.ReadDir(defaultJobDir(jobDir))
	if err != nil {
		if os.IsNotExist(err) {
			return []Job{}, nil
		}
		return nil, err
	}

	jobs := []Job{}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(defaultJobDir(jobDir), entry.Name()))
		if err != nil {
			return nil, err
		}
		var job Job
		if err := json.Unmarshal(trimUTF8BOM(data), &job); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	sortJobs(jobs)
	return jobs, nil
}

func ReadJobLog(jobDir string, id string, maxBytes int64) (string, error) {
	job, err := ReadJob(jobDir, id)
	if err != nil {
		return "", err
	}
	if maxBytes <= 0 {
		maxBytes = 32 * 1024
	}
	data, err := os.ReadFile(job.LogPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	if int64(len(data)) > maxBytes {
		data = data[int64(len(data))-maxBytes:]
	}
	return string(data), nil
}

func sortJobs(jobs []Job) {
	for i := 0; i < len(jobs); i++ {
		for j := i + 1; j < len(jobs); j++ {
			if jobs[j].CreatedAt.After(jobs[i].CreatedAt) {
				jobs[i], jobs[j] = jobs[j], jobs[i]
			}
		}
	}
}

func defaultJobDir(jobDir string) string {
	if jobDir == "" {
		return filepath.Join(".easy-setup", "logs")
	}
	return jobDir
}

func newJobID() (string, error) {
	var bytes [16]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes[:]), nil
}

func escapePowerShellString(value string) string {
	return strings.ReplaceAll(value, `"`, "`\"")
}

func trimUTF8BOM(data []byte) []byte {
	if len(data) >= 3 && data[0] == 0xef && data[1] == 0xbb && data[2] == 0xbf {
		return data[3:]
	}
	return data
}
