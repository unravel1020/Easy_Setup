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
