package models

import (
	"github.com/hjkelly/zbbapi/common"
)

type Budget struct {
	ID        SafeUUID        `json:"id" bson:"_id"`
	StartDate common.Date     `json:"startDate"`
	EndDate   common.Date     `json:"endDate"`
	Incomes   NamesAndAmounts `json:"incomes"`
	Bills     NamesAndAmounts `json:"bills"`
	Expenses  NamesAndAmounts `json:"expenses"`
	Savings   NamesAndAmounts `json:"savings"`
	Checklist []ChecklistItem `json:"checklist"`
	Balance   Amount
	Timestamped
}

func (budget Budget) GetValidated() (Budget, error) {
	// this will hold each error as we validate
	var err error
	// this will hold error results
	errs := make([]error, 0)

	// dates
	errs = append(errs, common.AddValidationContext(budget.StartDate.ValidateNonZero(), "startDate"))
	errs = append(errs, common.AddValidationContext(budget.EndDate.ValidateNonZero(), "endDate"))

	// income, expenses, bills
	budget.Incomes, err = budget.Incomes.GetValidated()
	errs = append(errs, common.AddValidationContext(err, "incomes"))
	budget.Expenses, err = budget.Expenses.GetValidated()
	errs = append(errs, common.AddValidationContext(err, "expenses"))
	budget.Bills, err = budget.Bills.GetValidated()
	errs = append(errs, common.AddValidationContext(err, "bills"))
	budget.Savings, err = budget.Savings.GetValidated()
	errs = append(errs, common.AddValidationContext(err, "savings"))

	// checklist items
	for i, item := range budget.Checklist {
		budget.Checklist[i], err = item.GetValidated()
		errs = append(errs, err)
	}

	// TODO: make sure they didn't try to provide read-only/protected fields

	// Calculate the balance.
	balance := 0
	for _, income := range budget.Incomes {
		balance += income.Amount.AmountCents
	}
	for _, expense := range budget.Expenses {
		balance -= expense.Amount.AmountCents
	}
	for _, bill := range budget.Bills {
		balance -= bill.Amount.AmountCents
	}
	for _, saving := range budget.Savings {
		balance -= saving.Amount.AmountCents
	}
	budget.Balance.AmountCents = balance

	// Finalize the errors, if there were any.
	err = common.CombineErrors(errs...)
	if err != nil {
		return Budget{}, err
	}

	// Otherwise, add any sanitized values.
	return budget, nil
}

func (budget Budget) AddPlanData(plan Plan) Budget {
	for _, income := range plan.Incomes {
		budget.Incomes = append(budget.Incomes, income.NameAndAmount)
	}
	for _, expense := range plan.Expenses {
		budget.Expenses = append(budget.Expenses, expense.NameAndAmount)
	}
	for _, bill := range plan.Bills {
		budget.Bills = append(budget.Bills, bill.NameAndAmount)
	}
	for _, saving := range plan.Savings {
		budget.Savings = append(budget.Savings, saving.NameAndAmount)
	}
	return budget
}

type ChecklistItem struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func (item ChecklistItem) GetValidated() (ChecklistItem, error) {
	// TODO
	return item, nil
}
