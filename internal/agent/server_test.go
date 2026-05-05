package agent

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
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

func TestExecuteRequiresHighRiskConfirmation(t *testing.T) {
	catalogData := highRiskCatalog()
	server := New(catalogData, "windows", true)
	server.JobDir = t.TempDir()
	server.Launch = func(scriptPath string) error { return nil }

	request := httptest.NewRequest(http.MethodPost, "/execute", bytes.NewBufferString(`{"itemIds":["danger.remote"]}`))
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, request)
	if response.Code != http.StatusConflict {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusConflict)
	}
	var blocked struct {
		Reason  string           `json:"reason"`
		Actions []catalog.Action `json:"actions"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &blocked); err != nil {
		t.Fatal(err)
	}
	if blocked.Reason != "high risk confirmation required" || len(blocked.Actions) != 1 {
		t.Fatalf("blocked = %#v", blocked)
	}

	confirmed := httptest.NewRequest(http.MethodPost, "/execute", bytes.NewBufferString(`{"itemIds":["danger.remote"],"confirmHighRisk":true}`))
	confirmedResponse := httptest.NewRecorder()
	server.Handler().ServeHTTP(confirmedResponse, confirmed)
	if confirmedResponse.Code != http.StatusOK {
		t.Fatalf("confirmed status = %d, body = %s", confirmedResponse.Code, confirmedResponse.Body.String())
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

func TestJobLogEndpoint(t *testing.T) {
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
	if err := os.WriteFile(executeResult.Job.LogPath, []byte("hello log"), 0o644); err != nil {
		t.Fatal(err)
	}

	request := httptest.NewRequest(http.MethodGet, "/jobs/"+executeResult.Job.ID+"/log", nil)
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d", response.Code)
	}
	var result struct {
		Log string `json:"log"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.Log != "hello log" {
		t.Fatalf("log = %q", result.Log)
	}
}

func TestJobEventsEndpoint(t *testing.T) {
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
	if err := os.WriteFile(executeResult.Job.LogPath, []byte("hello stream"), 0o644); err != nil {
		t.Fatal(err)
	}

	request := httptest.NewRequest(http.MethodGet, "/jobs/"+executeResult.Job.ID+"/events?once=1", nil)
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d", response.Code)
	}
	if contentType := response.Header().Get("Content-Type"); !strings.Contains(contentType, "text/event-stream") {
		t.Fatalf("content type = %q", contentType)
	}
	body := response.Body.String()
	if !strings.Contains(body, "event: job") || !strings.Contains(body, "hello stream") {
		t.Fatalf("unexpected event body: %s", body)
	}
}

func TestJobCancelEndpoint(t *testing.T) {
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

	request := httptest.NewRequest(http.MethodPost, "/jobs/"+executeResult.Job.ID+"/cancel", nil)
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d", response.Code)
	}
	var result struct {
		Job Job `json:"job"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.Job.Status != "cancel-requested" {
		t.Fatalf("status = %q", result.Job.Status)
	}
	if _, err := os.Stat(executeResult.Job.CancelPath); err != nil {
		t.Fatal(err)
	}
}

func TestJobRetryEndpoint(t *testing.T) {
	catalogData, err := catalog.Load("../../src/catalog/catalog.json")
	if err != nil {
		t.Fatal(err)
	}
	jobDir := t.TempDir()
	launchCount := 0
	server := New(catalogData, "windows", true)
	server.JobDir = jobDir
	server.Launch = func(scriptPath string) error {
		launchCount++
		return nil
	}

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

	request := httptest.NewRequest(http.MethodPost, "/jobs/"+executeResult.Job.ID+"/retry", nil)
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d", response.Code)
	}
	var result struct {
		Previous Job `json:"previous"`
		Job      Job `json:"job"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.Previous.ID != executeResult.Job.ID {
		t.Fatalf("previous id = %q", result.Previous.ID)
	}
	if result.Job.ID == "" || result.Job.ID == executeResult.Job.ID {
		t.Fatalf("retry job id = %q, previous = %q", result.Job.ID, executeResult.Job.ID)
	}
	if launchCount != 2 {
		t.Fatalf("launch count = %d", launchCount)
	}
	if len(result.Job.ItemIDs) != len(executeResult.Job.ItemIDs) {
		t.Fatalf("item ids = %#v", result.Job.ItemIDs)
	}
}

func TestJobRetryRequiresHighRiskConfirmation(t *testing.T) {
	catalogData := highRiskCatalog()
	jobDir := t.TempDir()
	job := Job{
		ID:      "retrydanger",
		Status:  "failed",
		JobPath: filepath.Join(jobDir, "retrydanger.json"),
		ItemIDs: []string{"danger.remote"},
	}
	if err := writeJob(job.JobPath, &job); err != nil {
		t.Fatal(err)
	}
	server := New(catalogData, "windows", true)
	server.JobDir = jobDir
	server.Launch = func(scriptPath string) error { return nil }

	request := httptest.NewRequest(http.MethodPost, "/jobs/retrydanger/retry", nil)
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, request)
	if response.Code != http.StatusConflict {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusConflict)
	}

	confirmed := httptest.NewRequest(http.MethodPost, "/jobs/retrydanger/retry", bytes.NewBufferString(`{"confirmHighRisk":true}`))
	confirmedResponse := httptest.NewRecorder()
	server.Handler().ServeHTTP(confirmedResponse, confirmed)
	if confirmedResponse.Code != http.StatusOK {
		t.Fatalf("confirmed status = %d, body = %s", confirmedResponse.Code, confirmedResponse.Body.String())
	}
}

func TestJobRetryRequiresExecutePermission(t *testing.T) {
	catalogData, err := catalog.Load("../../src/catalog/catalog.json")
	if err != nil {
		t.Fatal(err)
	}
	jobDir := t.TempDir()
	job := Job{
		ID:      "retryjob",
		Status:  "failed",
		JobPath: filepath.Join(jobDir, "retryjob.json"),
		ItemIDs: []string{"template.fullstack-web"},
	}
	if err := writeJob(job.JobPath, &job); err != nil {
		t.Fatal(err)
	}
	server := New(catalogData, "windows", false)
	server.JobDir = jobDir

	request := httptest.NewRequest(http.MethodPost, "/jobs/retryjob/retry", nil)
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, request)
	if response.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusForbidden)
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

func highRiskCatalog() *catalog.Catalog {
	return &catalog.Catalog{
		Version: "test",
		Recipes: []catalog.Recipe{
			{
				ID:       "danger.remote",
				Name:     "Remote Script",
				Category: "testing",
				Install: map[string]string{
					"windows": "irm https://example.com/install.ps1 | iex",
				},
			},
		},
	}
}
