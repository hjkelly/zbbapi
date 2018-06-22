package models

import (
	"github.com/hjkelly/zbbapi/common"
)

// PlanToBudgetConversion represents the data collected when turning a plan into a budget.
type Conversion struct {
	ID            SafeUUID        `json:"id" bson:"_id"` // auto-populated
	BudgetID      SafeUUID        `json:"budgetID"`      // auto-populated
	PlanID        SafeUUID        `json:"planID"`
	StartDate     common.Date     `json:"startDate"`
	EndDate       common.Date     `json:"endDate"`
	ExactIncomes  NamesAndAmounts `json:"exactIncomes"`
	ExactExpenses NamesAndAmounts `json:"exactExpenses"`
	ExactBills    NamesAndAmounts `json:"exactBills"`
	Timestamped
}

func (conversion Conversion) GetValidated() (Conversion, error) {
	var err error
	errs := []error{}
	// start/end dates valid?
	errs = append(errs, common.AddValidationContext(conversion.StartDate.ValidateNonZero(), "startDate"))
	errs = append(errs, common.AddValidationContext(conversion.EndDate.ValidateNonZero(), "endDate"))
	// plan ID valid?
	conversion.PlanID, err = conversion.PlanID.GetValidated()
	if err != nil {
		errs = append(errs, err)
	}
	// incomes, expenses, bills valid?
	conversion.ExactIncomes, err = conversion.ExactIncomes.GetValidated()
	if err != nil {
		errs = append(errs, common.AddValidationContext(err, "exactIncomes"))
	}
	conversion.ExactExpenses, err = conversion.ExactExpenses.GetValidated()
	if err != nil {
		errs = append(errs, common.AddValidationContext(err, "exactExpenses"))
	}
	conversion.ExactBills, err = conversion.ExactBills.GetValidated()
	if err != nil {
		errs = append(errs, common.AddValidationContext(err, "exactBills"))
	}
	return conversion, nil
}

func (conversion Conversion) MakeBudget(plan Plan) Budget {
	budget := Budget{
		StartDate: conversion.StartDate,
		EndDate:   conversion.EndDate,
	}
	daysSince := conversion.EndDate.DaysSince(conversion.StartDate)
	partialMonthMultiplier := daysSince / 30.0

	// incorporate exact incomes
	exactIncomeMap := conversion.ExactIncomes.AsMap()
	budget.Incomes = make(NamesAndAmounts, 0, len(plan.Incomes))
	for _, planItem := range plan.Incomes {
		var budgetAmount Amount
		if exactAmount, found := exactIncomeMap[planItem.Name]; found {
			budgetAmount = exactAmount
			// remove the item so we know not to add it as a custom income
			delete(exactIncomeMap, planItem.Name)
		} else {
			budgetAmount = planItem.Amount
		}
		budget.Incomes = append(budget.Incomes, NameAndAmount{
			Name:   planItem.Name,
			Amount: budgetAmount,
		})
	}
	// if we still have exact incomes in the map, that means they're custom and should be added
	for name, exactAmount := range exactIncomeMap {
		budget.Incomes = append(budget.Incomes, NameAndAmount{
			Name:   name,
			Amount: exactAmount,
		})
	}

	// incorporate exact bills
	exactBillMap := conversion.ExactExpenses.AsMap()
	budget.Bills = make(NamesAndAmounts, 0, len(plan.Bills))
	for _, planItem := range plan.Bills {
		var budgetAmount Amount
		if exactAmount, found := exactBillMap[planItem.Name]; found {
			budgetAmount = exactAmount
			// remove the item so we know not to add it as a custom income
			delete(exactBillMap, planItem.Name)
		} else {
			// if it falls in this date range, include it.
			budgetAmount = planItem.Amount
			// TODO
		}
		budget.Bills = append(budget.Bills, NameAndAmount{
			Name:   planItem.Name,
			Amount: budgetAmount,
		})
	}
	// if we still have exact incomes in the map, that means they're custom and should be added
	for name, exactAmount := range exactBillMap {
		budget.Bills = append(budget.Bills, NameAndAmount{
			Name:   name,
			Amount: exactAmount,
		})
	}

	// incorporate exact expenses
	exactExpenseMap := conversion.ExactExpenses.AsMap()
	budget.Expenses = make(NamesAndAmounts, 0, len(plan.Expenses))
	for _, planItem := range plan.Expenses {
		var budgetAmount Amount
		if exactAmount, found := exactExpenseMap[planItem.Name]; found {
			budgetAmount = exactAmount
			// remove the item so we know not to add it as a custom income
			delete(exactExpenseMap, planItem.Name)
		} else {
			budgetAmount = Amount{planItem.Amount.AmountCents * partialMonthMultiplier}
		}
		budget.Expenses = append(budget.Expenses, NameAndAmount{
			Name:   planItem.Name,
			Amount: budgetAmount,
		})
	}
	// if we still have exact incomes in the map, that means they're custom and should be added
	for name, exactAmount := range exactExpenseMap {
		budget.Expenses = append(budget.Expenses, NameAndAmount{
			Name:   name,
			Amount: exactAmount,
		})
	}

	return budget
}
