package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"titanium-rails/internal/ledger"
	"titanium-rails/internal/stms"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	Health(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response["status"] != "ok" {
		t.Errorf("expected status 'ok', got %s", response["status"])
	}
}

func TestSTMSAuth(t *testing.T) {
	l := ledger.NewStub()
	h := New(l)

	reqBody := stms.AuthRequest{
		Amount: 1000,
		CardId: "card123",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/stms/auth", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.STMSAuth(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var response stms.AuthResponse
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if !response.Approved {
		t.Errorf("expected approved=true, got false")
	}
}

