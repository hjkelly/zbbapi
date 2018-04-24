package models

import "github.com/hjkelly/zbbapi/common"

// CategoryRefAndAmount holds a category reference/ID and an amount. This is a pretty common data structure, so it may be used in budgeting/planning or in tracking actual expenses.
type CategoryRefAndAmount struct {
	CategoryID SafeUUID `json:"categoryID"`
	Amount
}

// GetValidated returns a sanitized copy, if the ID and amount are valid; otherwise, returns an error.
func (cram CategoryRefAndAmount) GetValidated() (CategoryRefAndAmount, error) {
	cleanCatID, catIDErr := cram.CategoryID.GetValidated()
	cleanAmount, amountErr := cram.Amount.GetValidated()

	err := common.CombineErrors(catIDErr, amountErr)
	if err != nil {
		return CategoryRefAndAmount{}, err
	}

	cram.CategoryID = cleanCatID
	cram.Amount = cleanAmount
	return cram, nil
}
