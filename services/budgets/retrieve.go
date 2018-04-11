package budgets

import (
	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Retrieve fetches a single Budget from the database, if its ID exists.
func Retrieve(id string) (*models.Budget, error) {
	ds := newDatastore()
	result := new(models.Budget)
	err := ds.C().Find(bson.M{
		"_id": uuid.FromStringOrNil(id),
	}).One(&result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, common.NotFoundErr
		}
		return nil, err
	}
	return result, nil
}
