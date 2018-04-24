package models

import "github.com/hjkelly/zbbapi/common"

// Amount is the central model for putting a dollar amount on anything.
type Amount struct {
	AmountCents int64 `json:"amount"`
}

// GetValidated returns a sanitized copy, or an error if something isn't right.
func (a Amount) GetValidated() (Amount, error) {
	if a.AmountCents < 0 {
		return Amount{}, common.NewValidationError("amount", common.NumOutOfRangeCode, "The amount cannot be negative.")
	}
	return a, nil
}
