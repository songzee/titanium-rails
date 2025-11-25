package handlers

import (
	"encoding/json"
	"net/http"
	"titanium-rails/internal/ledger"
	"titanium-rails/internal/stms"
)

type Handler struct {
	ledger ledger.Ledger
}

func New(l ledger.Ledger) *Handler {
	return &Handler{ledger: l}
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) STMSAuth(w http.ResponseWriter, r *http.Request) {
	var req stms.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Stub: always approve
	resp := stms.AuthResponse{
		Approved: true,
		Reason:   "approved",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

