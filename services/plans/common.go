package plans

import (
	"github.com/hjkelly/zbbapi/models"
)

// Make sure this Plan has input sufficient enough to be saved.
func getValidated(input models.Plan) (models.Plan, error) {
	return input.GetValidated()
}

// Returns the updated Plan, which is the current Plan updated with the input data for the update.
func getUpdated(current, input models.Plan) models.Plan {
	return current
}
