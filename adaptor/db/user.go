package db

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserCollection struct {
	mgo.Collection
}

type User struct {
	ID         bson.ObjectId
	Email      string
	password   string
	Active     bool
	Domains    []Domain.ID
	URLs       []URL.ID
	APIKeys    []APIKey.ID
	created_at time.Time
	updated_at time.Time
}

func GetUserCollection(DB *mgo.Database) *UserCollection {
	return DB.C("user")
}

func (uc *UserCollection) Create(user *User) error {
	err := uc.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Find(user *User) {
	c.Find(user).One(&User)
}
