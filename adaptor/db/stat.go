package db

import (
	"time"

	"gopkg.in/mgo.v2"
)

type Stat struct {
	unique_ip_access int
	ip_access        int
	clicks           int
	day              time.Time
	URLID            mgo.ObjectId
}

func GetStatCollection(DB *mgo.Database) *mgo.Collection {
	return DB.C("stat")
}
