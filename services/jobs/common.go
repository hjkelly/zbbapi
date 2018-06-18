package jobs

import (
	"github.com/hjkelly/zbbapi/models"
)

// Make sure this Job has input sufficient enough to be saved.
func getValidated(input models.Job) (models.Job, error) {
	return input.GetValidated()
}

// Returns the updated Job, which is the current Job updated with the input data for the update.
func getUpdated(current, input models.Job) models.Job {
	current.Status = input.Status
	current.Result = input.Result
	return current
}
