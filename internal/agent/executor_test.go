package agent

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
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
