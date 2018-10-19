package dao

import (
	. "../models"
	"gopkg.in/mgo.v2/bson"
)

// Find a CED by its id
func (c *DAO) FindCEDById(id string) (CED, error) {
	var ced CED
	err := db.C("CarrierEntity").FindId(bson.ObjectIdHex(id)).One(&ced)
	return ced, err
}
