package db

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Stat struct {
	ID           bson.ObjectId `bson:"_id"`
	UniqueClicks int           `bson:"unique_clicks"`
	IP           int           `bson:"ip"`
	Clicks       int           `bson:"clicks"`
	Day          time.Time     `bson:"day"`
}

func GetStatCollection(DB *mgo.Database) *mgo.Collection {
	return DB.C("stat")
}
