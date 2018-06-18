package jobs

import (
	"github.com/hjkelly/zbbapi/models"
	"gopkg.in/mgo.v2/bson"
)

// List returns all Jobs from the database.
func List() ([]models.Job, error) {
	ds := newDatastore()
	results := make([]models.Job, 0)
	err := ds.C().Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
