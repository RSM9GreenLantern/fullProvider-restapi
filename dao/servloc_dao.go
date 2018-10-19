package dao

import (
	. "../models"
	"gopkg.in/mgo.v2/bson"
)

// Find a movie by its id
func (s *DAO) FindServLocById(id string) (ServLoc, error) {
	var servloc ServLoc
	err := db.C("serviceLocation").FindId(bson.ObjectIdHex(id)).One(&servloc)
	return servloc, err
}