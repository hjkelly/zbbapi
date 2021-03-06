package categories

import (
	"github.com/hjkelly/zbbapi/models"
	uuid "github.com/satori/go.uuid"
)

// Create validates and preps a Category, then saves it via the controller's datastore.
func Create(input models.Category) (*models.Category, error) {
	input = sanitize(input)
	// Did they give us enough to save?
	err := validate(input)
	if err != nil {
		return nil, err
	}

	// prepare the rest of the resource
	input.ID = uuid.NewV4()
	input.SetCreationTimestamp()

	// save
	ds := newDatastore()
	err = ds.C().Insert(input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}
