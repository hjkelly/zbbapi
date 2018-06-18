package jobs

import (
	"github.com/hjkelly/zbbapi/models"
)

// Create validates and preps a Job, then saves it via the controller's datastore.
func Create(input models.Job) (*models.Job, error) {
	var err error
	input, err = getValidated(input)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	// prepare the rest of the resource
	input.SetCreationTimestamp()
	// fill in the UUID if they didn't provide one
	if len(input.ID) == 0 {
		input.ID = string(models.NewSafeUUID())
	}

	// save
	ds := newDatastore()
	err = ds.C().Insert(input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}
