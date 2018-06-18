package budgets

import (
	"github.com/hjkelly/zbbapi/models"
)

// Make sure this Budget has input sufficient enough to be saved.
func getValidated(input models.Budget) (models.Budget, error) {
	validated, err := input.GetValidated()
	if err != nil {
		return models.Budget{}, err
	}
	return validated, nil
}

// Returns the updated Budget, which is the current Budget updated with the input data for the update.
func getUpdated(current, input models.Budget) models.Budget {
	current.StartDate = input.StartDate
	current.EndDate = input.EndDate
	current.Incomes = input.Incomes
	current.Expenses = input.Expenses
	current.Bills = input.Bills
	current.Savings = input.Savings
	current.Checklist = input.Checklist
	return current
}
