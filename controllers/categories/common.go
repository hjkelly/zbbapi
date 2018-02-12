package categories

import (
	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
)

func validate(input models.Category) error {
	if common.StringIsEmpty(input.Name) {
		return common.NewValidationError("name", "REQUIRED_TEXT", "You must provide a name.")
	}
	return nil
}

func getUpdated(current, input models.Category) models.Category {
	current.Name = input.Name
	return current
}
