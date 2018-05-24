package plans

import (
	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Retrieve fetches a single Plan from the database, if its ID exists.
func Retrieve(id string) (*models.Plan, error) {
	ds := newDatastore()
	result := new(models.Plan)
	err := ds.C().Find(bson.M{
		"_id": id,
	}).One(&result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, common.NotFoundErr
		}
		return nil, err
	}
	return result, nil
}
