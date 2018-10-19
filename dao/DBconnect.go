package dao

import (
	"log"
	mgo "gopkg.in/mgo.v2"
)

type DAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// Establish a connection to database
func (d *DAO) Connect() {
	session, err := mgo.Dial(d.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(d.Database)
}