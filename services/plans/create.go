package plans

import (
	"github.com/hjkelly/zbbapi/models"
)

// Create validates and preps a Plan, then saves it via the controller's datastore.
func Create(input models.Plan) (*models.Plan, error) {
	// Did they give us enough to save?
	var err error
	input, err = getValidated(input)
	if err != nil {
		return nil, err
	}

	// prepare the rest of the resource
	input.ID = models.NewSafeUUID()
	input.SetCreationTimestamp()

	// save
	ds := newDatastore()
	err = ds.C().Insert(input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}
