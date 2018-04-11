package budgets

import (
	"github.com/hjkelly/zbbapi/common"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Delete uses the datastore to remove this ID, if it exists.
func Delete(id string) error {
	ds := newDatastore()
	err := ds.C().Remove(bson.M{
		"_id": uuid.FromStringOrNil(id),
	})
	if err != nil {
		if err == mgo.ErrNotFound {
			return common.NotFoundErr
		}
		return err
	}
	return nil
}
