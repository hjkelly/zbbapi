package categories

import (
	"github.com/hjkelly/zbbapi/models"
	"gopkg.in/mgo.v2/bson"
)

// List returns all Categories from the database.
func List() ([]models.Category, error) {
	ds := newDatastore()
	results := make([]models.Category, 0)
	err := ds.C().Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
