package db

import (
	"gopkg.in/mgo.v2"
)

type Adaptor struct {
	user       *mgo.Collection
	url        *mgo.Collection
	domains    *mgo.Collection
	api_tokens *mgo.Collection
	stats      *mgo.Collection
}

func New(hosts string) (*Adaptor, *mgo.Session) {
	session := mgo.Dial(hosts)
	DB := session.DB("brief_url_db")

	return &Adaptor{
		user:      getUserCollection(DB),
		url:       getURLCollection(DB),
		domain:    getDomainCollection(DB),
		api_token: getAPITokenCollection(DB),
		stat:      getStatCollection(DB),
	}
}
