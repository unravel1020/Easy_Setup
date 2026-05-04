package agent

import (
	"encoding/json"
	"net/http"

	"github.com/unravel1020/Easy_Setup/internal/catalog"
)

type Server struct {
	Catalog      *catalog.Catalog
	AllowExecute bool
	Platform     string
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
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/catalog", s.handleCatalog)
	mux.HandleFunc("/plan", s.handlePlan)
	mux.HandleFunc("/execute", s.handleExecute)
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
			"reason": "Go Agent execution is not enabled yet.",
		})
		return
	}
	writeError(w, http.StatusNotImplemented, "execution will be implemented after job logging and confirmation are ported")
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
