package db

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID         bson.ObjectId
	Email      string
	password   string
	Active     bool
	Domains    []bson.ObjectId
	URLs       []bson.ObjectId
	APIKeys    []bson.ObjectId
	created_at time.Time
	updated_at time.Time
}

func GetUserCollection(DB *mgo.Database) *mgo.Collection {
	return DB.C("user")
}
