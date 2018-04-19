package budgets

import (
	"github.com/hjkelly/zbbapi/models"
)

// Make sure this Budget has input sufficient enough to be saved.
func getValidated(input models.Budget) (models.Budget, error) {
	return input.GetValidated()
}

// Returns the updated Budget, which is the current Budget updated with the input data for the update.
func getUpdated(current, input models.Budget) models.Budget {
	return current
}
