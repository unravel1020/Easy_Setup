package agent

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/unravel1020/Easy_Setup/internal/catalog"
)

func TestPlanEndpoint(t *testing.T) {
	catalogData, err := catalog.Load("../../src/catalog/catalog.json")
	if err != nil {
		t.Fatal(err)
	}
	server := New(catalogData, "windows", false)

	body := bytes.NewBufferString(`{"itemIds":["template.fullstack-web"]}`)
	request := httptest.NewRequest(http.MethodPost, "/plan", body)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	server.Handler().ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", response.Code, response.Body.String())
	}

	var plan catalog.Plan
	if err := json.Unmarshal(response.Body.Bytes(), &plan); err != nil {
		t.Fatal(err)
	}
	if len(plan.Actions) == 0 {
		t.Fatal("expected plan actions")
	}
	if plan.Actions[0].ID != "git" {
		t.Fatalf("first action = %q, want git", plan.Actions[0].ID)
	}
}

func TestExecuteDisabled(t *testing.T) {
	catalogData, err := catalog.Load("../../src/catalog/catalog.json")
	if err != nil {
		t.Fatal(err)
	}
	server := New(catalogData, "windows", false)

	request := httptest.NewRequest(http.MethodPost, "/execute", bytes.NewBufferString(`{}`))
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, request)

	if response.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusForbidden)
	}
}

func TestExecuteCreatesJob(t *testing.T) {
	catalogData, err := catalog.Load("../../src/catalog/catalog.json")
	if err != nil {
		t.Fatal(err)
	}
	jobDir := t.TempDir()
	launched := ""
	server := New(catalogData, "windows", true)
	server.JobDir = jobDir
	server.Launch = func(scriptPath string) error {
		launched = scriptPath
		return nil
	}

	body := bytes.NewBufferString(`{"itemIds":["template.fullstack-web"]}`)
	request := httptest.NewRequest(http.MethodPost, "/execute", body)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", response.Code, response.Body.String())
	}
	if launched == "" {
		t.Fatal("expected launcher to be called")
	}
	if _, err := os.Stat(launched); err != nil {
		t.Fatal(err)
	}
	if filepath.Dir(launched) != jobDir {
		t.Fatalf("script dir = %q, want %q", filepath.Dir(launched), jobDir)
	}
}

func TestJobsEndpoints(t *testing.T) {
	catalogData, err := catalog.Load("../../src/catalog/catalog.json")
	if err != nil {
		t.Fatal(err)
	}
	jobDir := t.TempDir()
	server := New(catalogData, "windows", true)
	server.JobDir = jobDir
	server.Launch = func(scriptPath string) error { return nil }

	executeBody := bytes.NewBufferString(`{"itemIds":["template.fullstack-web"]}`)
	executeRequest := httptest.NewRequest(http.MethodPost, "/execute", executeBody)
	executeResponse := httptest.NewRecorder()
	server.Handler().ServeHTTP(executeResponse, executeRequest)
	if executeResponse.Code != http.StatusOK {
		t.Fatalf("execute status = %d", executeResponse.Code)
	}

	var executeResult struct {
		Job Job `json:"job"`
	}
	if err := json.Unmarshal(executeResponse.Body.Bytes(), &executeResult); err != nil {
		t.Fatal(err)
	}

	listRequest := httptest.NewRequest(http.MethodGet, "/jobs", nil)
	listResponse := httptest.NewRecorder()
	server.Handler().ServeHTTP(listResponse, listRequest)
	if listResponse.Code != http.StatusOK {
		t.Fatalf("list status = %d", listResponse.Code)
	}
	var listResult struct {
		Jobs []Job `json:"jobs"`
	}
	if err := json.Unmarshal(listResponse.Body.Bytes(), &listResult); err != nil {
		t.Fatal(err)
	}
	if len(listResult.Jobs) != 1 {
		t.Fatalf("jobs = %d, want 1", len(listResult.Jobs))
	}

	getRequest := httptest.NewRequest(http.MethodGet, "/jobs/"+executeResult.Job.ID, nil)
	getResponse := httptest.NewRecorder()
	server.Handler().ServeHTTP(getResponse, getRequest)
	if getResponse.Code != http.StatusOK {
		t.Fatalf("get status = %d", getResponse.Code)
	}
	var got Job
	if err := json.Unmarshal(getResponse.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.ID != executeResult.Job.ID {
		t.Fatalf("job id = %q, want %q", got.ID, executeResult.Job.ID)
	}
}

func TestJobNotFound(t *testing.T) {
	catalogData, err := catalog.Load("../../src/catalog/catalog.json")
	if err != nil {
		t.Fatal(err)
	}
	server := New(catalogData, "windows", false)
	server.JobDir = t.TempDir()

	request := httptest.NewRequest(http.MethodGet, "/jobs/missing", nil)
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, request)
	if response.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusNotFound)
	}
}
