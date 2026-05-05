package agent

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/unravel1020/Easy_Setup/internal/catalog"
)

type Server struct {
	Catalog      *catalog.Catalog
	AllowExecute bool
	Platform     string
	JobDir       string
	Launch       Launcher
}

type planRequest struct {
	ItemIDs []string `json:"itemIds"`
}

func New(catalogData *catalog.Catalog, platform string, allowExecute bool) *Server {
	if platform == "" {
		platform = "windows"
	}
	return &Server{
		Catalog:      catalogData,
		AllowExecute: allowExecute,
		Platform:     platform,
		JobDir:       ".easy-setup/logs",
		Launch:       DefaultLauncher,
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/catalog", s.handleCatalog)
	mux.HandleFunc("/plan", s.handlePlan)
	mux.HandleFunc("/execute", s.handleExecute)
	mux.HandleFunc("/jobs", s.handleJobs)
	mux.HandleFunc("/jobs/", s.handleJob)
	return cors(mux)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":           true,
		"name":         "Easy_Setup Go Agent",
		"version":      "0.1.0",
		"allowExecute": s.AllowExecute,
	})
}

func (s *Server) handleCatalog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	writeJSON(w, http.StatusOK, s.Catalog)
}

func (s *Server) handlePlan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var request planRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	plan, err := s.Catalog.Plan(s.Platform, request.ItemIDs)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, plan)
}

func (s *Server) handleExecute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	if !s.AllowExecute {
		writeJSON(w, http.StatusForbidden, map[string]any{
			"ok":     false,
			"reason": "Agent was not started with -allow-execute.",
		})
		return
	}

	var request planRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	plan, err := s.Catalog.Plan(s.Platform, request.ItemIDs)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	job, err := CreateJob(plan, s.JobDir, s.Launch)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]any{
			"ok":     false,
			"reason": err.Error(),
			"job":    job,
		})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":  true,
		"job": job,
	})
}

func (s *Server) handleJobs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	jobs, err := ListJobs(s.JobDir)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"jobs": jobs,
	})
}

func (s *Server) handleJob(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/jobs/")
	if strings.HasSuffix(id, "/log") {
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		s.handleJobLog(w, r, strings.TrimSuffix(id, "/log"))
		return
	}
	if strings.HasSuffix(id, "/cancel") {
		if r.Method != http.MethodPost {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		s.handleJobCancel(w, r, strings.TrimSuffix(id, "/cancel"))
		return
	}
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	if id == "" || strings.Contains(id, "/") {
		writeError(w, http.StatusNotFound, "job not found")
		return
	}
	job, err := ReadJob(s.JobDir, id)
	if err != nil {
		if errors.Is(err, ErrJobNotFound) {
			writeError(w, http.StatusNotFound, "job not found")
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, job)
}

func (s *Server) handleJobLog(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" || strings.Contains(id, "/") {
		writeError(w, http.StatusNotFound, "job not found")
		return
	}
	logText, err := ReadJobLog(s.JobDir, id, 32*1024)
	if err != nil {
		if errors.Is(err, ErrJobNotFound) {
			writeError(w, http.StatusNotFound, "job not found")
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"id":  id,
		"log": logText,
	})
}

func (s *Server) handleJobCancel(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" || strings.Contains(id, "/") {
		writeError(w, http.StatusNotFound, "job not found")
		return
	}
	job, err := CancelJob(s.JobDir, id)
	if err != nil {
		if errors.Is(err, ErrJobNotFound) {
			writeError(w, http.StatusNotFound, "job not found")
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":  true,
		"job": job,
	})
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, reason string) {
	writeJSON(w, status, map[string]any{
		"ok":     false,
		"reason": reason,
	})
}
