package ledger

import "testing"

func TestStub_CheckBalance(t *testing.T) {
	s := NewStub()
	balance, err := s.CheckBalance("account123")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if balance != 10000 {
		t.Errorf("expected balance 10000, got %d", balance)
	}
}

func TestStub_PlaceHold(t *testing.T) {
	s := NewStub()
	err := s.PlaceHold("account123", 1000)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

