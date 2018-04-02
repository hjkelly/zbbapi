package budgets

import (
	"github.com/hjkelly/zbbapi/models"
)

// Do any trimming, cleanup before validation.
func sanitize(input models.Budget) models.Budget {
	return input
}

// Make sure this Budget has input sufficient enough to be saved.
func validate(input models.Budget) error {
	return nil
}

// Returns the updated Budget, which is the current Budget updated with the input data for the update.
func getUpdated(current, input models.Budget) models.Budget {
	return current
}
