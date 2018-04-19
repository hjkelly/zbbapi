package models

import "github.com/hjkelly/zbbapi/common"

type CategoryRefAndAmount struct {
	CategoryID SafeUUID `json:"categoryID"`
	Amount
}

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
