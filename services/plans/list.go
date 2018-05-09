package plans

import (
	"github.com/hjkelly/zbbapi/models"
	"gopkg.in/mgo.v2/bson"
)

// List returns all Plans from the database.
func List() ([]models.Plan, error) {
	ds := newDatastore()
	results := make([]models.Plan, 0)
	err := ds.C().Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
