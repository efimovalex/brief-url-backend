package app

import (
	"net/http"

	"github.com/efimovalex/brief_url/adaptor/db"
	"github.com/efimovalex/brief_url/client"
)

// SampleEndpoint exists as an example
type SampleEndpoint struct {
	toggler db.Toggleable
}

func (ep *SampleEndpoint) Get(w http.ResponseWriter, r *http.Request) {
	/*
		Do what ever logic is needed
	*/

	w.WriteHeader(http.StatusTeapot)
	w.Write([]byte("Sample Endpoint Response"))
}
