package categories

import (
	"strings"

	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
)

// Do any trimming, cleanup before validation.
func sanitize(input models.Category) models.Category {
	input.Name = strings.TrimSpace(input.Name)
	return input
}

// Make sure this category has input sufficient enough to be saved.
func validate(input models.Category) error {
	if common.StringIsEmpty(input.Name) {
		return common.NewValidationError("name", "REQUIRED_TEXT", "You must provide a name.")
	}
	return nil
}

// Returns the updated category, which is the current category updated with the input data for the update.
func getUpdated(current, input models.Category) models.Category {
	current.Name = input.Name
	return current
}
