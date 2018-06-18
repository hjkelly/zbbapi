package jobs

import (
	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Retrieve fetches a single Job from the database, if its ID exists.
func Retrieve(id string) (*models.Job, error) {
	ds := newDatastore()
	result := new(models.Job)
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
