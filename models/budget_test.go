package models

import (
	"testing"

	"github.com/hjkelly/zbbapi/common"
)

func TestBudgetGetValidatedMinimal(t *testing.T) {
	b := Budget{
		StartDate: common.Date{Year: 2018, Month: 5, Day: 12},
		EndDate:   common.Date{Year: 2018, Month: 5, Day: 26},
	}
	validated, err := b.GetValidated()
	if err != nil {
		t.Errorf("Minimal valid budget wasn't allowed!")
	}
	if validated.Balance.AmountCents != 0 {
		t.Errorf("Balance wasn't calculated to be 0 when 0 items are in the budget.")
	}
}

func TestBudgetGetValidated(t *testing.T) {
	b := Budget{
		StartDate: common.Date{Year: 2018, Month: 5, Day: 12},
		EndDate:   common.Date{Year: 2018, Month: 5, Day: 26},
		Incomes: NamesAndAmounts{
			{Name: "X", Amount: Amount{AmountCents: 111111}},
			{Name: "X", Amount: Amount{AmountCents: 222222}},
		},
		Bills: NamesAndAmounts{
			{Name: "X", Amount: Amount{AmountCents: 33333}},
			{Name: "X", Amount: Amount{AmountCents: 44444}},
			{Name: "X", Amount: Amount{AmountCents: 5555}},
			{Name: "X", Amount: Amount{AmountCents: 6666}},
		},
		Savings: NamesAndAmounts{
			{Name: "X", Amount: Amount{AmountCents: 7777}},
		},
		Expenses: NamesAndAmounts{
			{Name: "X", Amount: Amount{AmountCents: 88888}},
			{Name: "X", Amount: Amount{AmountCents: 9999}},
		},
	}
	validated, err := b.GetValidated()
	if err != nil {
		t.Errorf("Got errors during validation! %+v", err)
	}
	expectedBalance := 111111 + 222222 - 33333 - 44444 - 5555 - 6666 - 7777 - 88888 - 9999
	if validated.Balance.AmountCents != expectedBalance {
		t.Errorf("Didn't get the expected balance of %d; instead, got: %d", expectedBalance, validated.Balance.AmountCents)
	}
}

func TestBudgetGetValidatedNegativeAmount(t *testing.T) {
	b := Budget{
		StartDate: common.Date{Year: 2018, Month: 5, Day: 12},
		EndDate:   common.Date{Year: 2018, Month: 5, Day: 26},
		Bills: NamesAndAmounts{
			{Name: "X", Amount: Amount{AmountCents: 1}},
			{Name: "X", Amount: Amount{AmountCents: 1}},
			{Name: "X", Amount: Amount{AmountCents: 1}},
		},
		Savings: NamesAndAmounts{
			{Name: "X", Amount: Amount{AmountCents: 1}},
		},
		Expenses: NamesAndAmounts{
			{Name: "X", Amount: Amount{AmountCents: 1}},
			{Name: "X", Amount: Amount{AmountCents: 1}},
		},
	}
	validated, err := b.GetValidated()
	if err != nil {
		t.Errorf("Got errors during validation! %+v", err)
	}
	expectedBalance := -6
	if validated.Balance.AmountCents != expectedBalance {
		t.Errorf("Didn't get the expected balance of %d; instead, got: %d", expectedBalance, validated.Balance.AmountCents)
	}
}
