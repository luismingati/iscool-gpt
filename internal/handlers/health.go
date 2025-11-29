package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	OK bool `json:"ok"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(HealthResponse{OK: true}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
