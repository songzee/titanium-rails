package stms

// AuthRequest represents an STMS authorization request
type AuthRequest struct {
	Amount int    `json:"amount"`
	CardId string `json:"cardId"`
}

// AuthResponse represents an STMS authorization response
type AuthResponse struct {
	Approved bool   `json:"approved"`
	Reason   string `json:"reason"`
}

