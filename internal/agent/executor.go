package agent

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/unravel1020/Easy_Setup/internal/catalog"
)

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
	if jobDir == "" {
		jobDir = filepath.Join(".easy-setup", "logs")
	}
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
