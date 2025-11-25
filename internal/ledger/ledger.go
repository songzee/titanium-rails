package ledger

// Ledger defines the interface for account operations
type Ledger interface {
	CheckBalance(accountId string) (int, error)
	PlaceHold(accountId string, amount int) error
}

// Stub is a minimal implementation for prototyping
type Stub struct{}

func NewStub() *Stub {
	return &Stub{}
}

func (s *Stub) CheckBalance(accountId string) (int, error) {
	// Stub: return fake balance
	return 10000, nil
}

func (s *Stub) PlaceHold(accountId string, amount int) error {
	// Stub: no-op
	return nil
}

