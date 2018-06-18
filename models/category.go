package models

import (
	"strconv"
	"strings"

	"github.com/hjkelly/zbbapi/common"
	uuid "github.com/satori/go.uuid"
)

// Category lets you build an expected budget and categorize your actual expenses.
type Category struct {
	ID   uuid.UUID `json:"id" bson:"_id"`
	Name string    `json:"name"`
	Timestamped
}

type NameAndAmount struct {
	Name string `json:"name"`
	Amount
}

func (cram NameAndAmount) GetValidated() (NameAndAmount, error) {
	var nameErr error
	cleanName := strings.TrimSpace(cram.Name)
	if len(cleanName) == 0 {
		nameErr = common.NewValidationError("name", common.MissingCode, "You must priovide a name.")
	}
	cleanAmount, amountErr := cram.Amount.GetValidated()

	err := common.CombineErrors(nameErr, amountErr)
	if err != nil {
		return NameAndAmount{}, err
	}

	cram.Name = cleanName
	cram.Amount = cleanAmount
	return cram, nil
}

type NamesAndAmounts []NameAndAmount

func (items NamesAndAmounts) GetValidated() (NamesAndAmounts, error) {
	errs := make([]error, 0)
	var itemErr error
	for i, item := range items {
		items[i], itemErr = item.GetValidated()
		errs = append(errs, common.AddValidationContext(itemErr, strconv.Itoa(i)))
	}
	err := common.CombineErrors(errs...)
	if err != nil {
		return NamesAndAmounts{}, err
	}
	return items, nil
}

func (items NamesAndAmounts) AsMap() map[string]Amount {
	itemMap := map[string]Amount{}
	for _, item := range items {
		itemMap[item.Name] = item.Amount
	}
	return itemMap
}
