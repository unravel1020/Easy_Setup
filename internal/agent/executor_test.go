package agent

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/unravel1020/Easy_Setup/internal/catalog"
)

func TestReadJobAcceptsUTF8BOM(t *testing.T) {
	jobDir := t.TempDir()
	job := Job{
		ID:        "bomjob",
		Status:    "launched",
		CreatedAt: time.Now().UTC(),
		ItemIDs:   []string{"template.fullstack-web"},
	}
	data, err := json.Marshal(job)
	if err != nil {
		t.Fatal(err)
	}
	data = append([]byte{0xef, 0xbb, 0xbf}, data...)
	if err := os.WriteFile(filepath.Join(jobDir, "bomjob.json"), data, 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := ReadJob(jobDir, "bomjob")
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != "bomjob" {
		t.Fatalf("job id = %q", got.ID)
	}

	jobs, err := ListJobs(jobDir)
	if err != nil {
		t.Fatal(err)
	}
	if len(jobs) != 1 || jobs[0].ID != "bomjob" {
		t.Fatalf("jobs = %#v", jobs)
	}
}

func TestRenderedPowerShellJobUpdatesStatus(t *testing.T) {
	job := &Job{
		ID:      "statusjob",
		JobPath: filepath.Join("jobs", "statusjob.json"),
		LogPath: filepath.Join("jobs", "statusjob.log"),
		Actions: []catalog.Action{
			{
				ID:      "sample",
				Command: "Write-Host sample",
				Verify:  []string{"Write-Host verify"},
			},
		},
	}

	script := renderPowerShellJob(job)
	for _, want := range []string{
		`Set-EasySetupJobStatus "running"`,
		`Test-EasySetupCanceled`,
		`Update-EasySetupExitCode $easySetupLastSucceeded`,
		`Set-EasySetupJobStatus "completed" 0`,
		`Set-EasySetupJobStatus "failed" $script:EasySetupExitCode`,
		`Set-EasySetupJobStatus "canceled" 130`,
		`completedAt`,
		`exitCode`,
	} {
		if !strings.Contains(script, want) {
			t.Fatalf("script missing %q:\n%s", want, script)
		}
	}
}

func TestCancelJobWritesCancelRequest(t *testing.T) {
	jobDir := t.TempDir()
	job := Job{
		ID:         "canceljob",
		Status:     "running",
		JobPath:    filepath.Join(jobDir, "canceljob.json"),
		CancelPath: filepath.Join(jobDir, "canceljob.cancel"),
		CreatedAt:  time.Now().UTC(),
	}
	if err := writeJob(job.JobPath, &job); err != nil {
		t.Fatal(err)
	}

	got, err := CancelJob(jobDir, "canceljob")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != "cancel-requested" {
		t.Fatalf("status = %q", got.Status)
	}
	if _, err := os.Stat(job.CancelPath); err != nil {
		t.Fatal(err)
	}

	stored, err := ReadJob(jobDir, "canceljob")
	if err != nil {
		t.Fatal(err)
	}
	if stored.Status != "cancel-requested" {
		t.Fatalf("stored status = %q", stored.Status)
	}
}

func TestPowerShellJobScriptCompletesStatus(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("PowerShell job script integration test runs on Windows")
	}
	if _, err := exec.LookPath("powershell"); err != nil {
		t.Skip("powershell is not available")
	}

	jobDir := t.TempDir()
	job := &Job{
		ID:         "completejob",
		Status:     "launched",
		JobPath:    filepath.Join(jobDir, "completejob.json"),
		ScriptPath: filepath.Join(jobDir, "completejob.ps1"),
		LogPath:    filepath.Join(jobDir, "completejob.log"),
		CreatedAt:  time.Now().UTC(),
		ItemIDs:    []string{"test"},
		Actions: []catalog.Action{
			{
				ID:      "safe.echo",
				Command: "Write-Host ok",
				Verify:  []string{"Write-Host verified"},
			},
		},
	}
	if err := writeJob(job.JobPath, job); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(job.ScriptPath, []byte(renderPowerShellJob(job)), 0o644); err != nil {
		t.Fatal(err)
	}

	command := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-File", job.ScriptPath)
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("script failed: %v\n%s", err, string(output))
	}

	got, err := ReadJob(jobDir, "completejob")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != "completed" {
		t.Fatalf("status = %q", got.Status)
	}
	if got.ExitCode == nil || *got.ExitCode != 0 {
		t.Fatalf("exit code = %#v", got.ExitCode)
	}
	if got.CompletedAt == nil {
		t.Fatal("expected completedAt")
	}
}

func TestPowerShellJobScriptCancelsAtStepBoundary(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("PowerShell job script integration test runs on Windows")
	}
	if _, err := exec.LookPath("powershell"); err != nil {
		t.Skip("powershell is not available")
	}

	jobDir := t.TempDir()
	job := &Job{
		ID:         "cancelpsjob",
		Status:     "launched",
		JobPath:    filepath.Join(jobDir, "cancelpsjob.json"),
		ScriptPath: filepath.Join(jobDir, "cancelpsjob.ps1"),
		LogPath:    filepath.Join(jobDir, "cancelpsjob.log"),
		CancelPath: filepath.Join(jobDir, "cancelpsjob.cancel"),
		CreatedAt:  time.Now().UTC(),
		ItemIDs:    []string{"test"},
		Actions: []catalog.Action{
			{
				ID:      "safe.echo",
				Command: "Write-Host should-not-run",
			},
		},
	}
	if err := writeJob(job.JobPath, job); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(job.CancelPath, []byte("cancel"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(job.ScriptPath, []byte(renderPowerShellJob(job)), 0o644); err != nil {
		t.Fatal(err)
	}

	command := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-File", job.ScriptPath)
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("script failed: %v\n%s", err, string(output))
	}

	got, err := ReadJob(jobDir, "cancelpsjob")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != "canceled" {
		t.Fatalf("status = %q", got.Status)
	}
	if got.ExitCode == nil || *got.ExitCode != 130 {
		t.Fatalf("exit code = %#v", got.ExitCode)
	}
}
