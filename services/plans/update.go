package plans

import (
	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UpdateID finds the current Plan by ID, updates all its user-updatable fields, and saves it again.
func UpdateID(id string, input models.Plan) (*models.Plan, error) {
	ds := newDatastore()

	// Make sure the one we're updating exists.
	current := models.Plan{}
	err := ds.C().Find(bson.M{
		"_id": uuid.FromStringOrNil(id),
	}).One(&current)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, common.NotFoundErr
		}
		return nil, err
	}

	// Validate the input and use it to update the current data.
	input, err = getValidated(input)
	if err != nil {
		return nil, err
	}
	result := getUpdated(current, input)
	result.SetModificationTimestamp()

	// Update the database with our new result.
	err = ds.C().UpdateId(result.ID, result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
