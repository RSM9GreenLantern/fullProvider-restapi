package dao

import (
	. "../models"
	"gopkg.in/mgo.v2/bson"
)

// Find a movie by its id
func (p *DAO) FindPracById(id string) (Prac, error) {
	var prac Prac
	err := db.C("prac").FindId(bson.ObjectIdHex(id)).One(&prac)
	return prac, err
}
