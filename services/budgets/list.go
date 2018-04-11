package budgets

import (
	"github.com/hjkelly/zbbapi/models"
	"gopkg.in/mgo.v2/bson"
)

// List returns all Budgets from the database.
func List() ([]models.Budget, error) {
	ds := newDatastore()
	results := make([]models.Budget, 0)
	err := ds.C().Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
