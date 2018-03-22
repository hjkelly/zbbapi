package categories

import (
	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
	uuid "github.com/satori/go.uuid"
)

func Create(input models.Category) (*models.Category, error) {
	// validate
	if common.StringIsEmpty(input.Name) {
		return nil, common.NewValidationError("name", "REQUIRED_TEXT", "You must provide a name.")
	}

	// prepare the input
	input.ID = uuid.NewV4()
	input.SetCreationTimestamp()

	// save
	ds := newDatastore()
	err := ds.C().Insert(input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}
