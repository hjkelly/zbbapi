package models

import "github.com/hjkelly/zbbapi/common"

type Amount struct {
	AmountCents int64 `json:"amount"`
}

func (a Amount) GetValidated() (Amount, error) {
	if a.AmountCents < 0 {
		return Amount{}, common.NewValidationError("amount", common.NumOutOfRangeCode, "The amount cannot be negative.")
	}
	return a, nil
}
