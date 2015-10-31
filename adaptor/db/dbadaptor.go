package db

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2"
)

type Adaptor struct {
	DB       *mgo.Database
	User     *mgo.Collection
	Url      *URLCollection
	Domain   *DomainCollection
	ApiToken *mgo.Collection
	Stat     *mgo.Collection
}

func New() (*Adaptor, error) {
	maxWait := time.Duration(5 * time.Second)
	session, err := mgo.DialWithTimeout("localhost:27017", maxWait)
	if err != nil {
		err = errors.New("MongoDB Error - " + err.Error())
		return &Adaptor{}, err
	}
	DB := session.DB("brief_url_db")

	return &Adaptor{
		DB:       DB,
		User:     GetUserCollection(DB),
		Url:      GetURLCollection(DB),
		Domain:   GetDomainCollection(DB),
		ApiToken: GetAPITokenCollection(DB),
		Stat:     GetStatCollection(DB),
	}, nil
}
