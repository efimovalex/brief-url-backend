package app

import (
	"errors"
	"log"

	"github.com/codegangsta/negroni"
	"github.com/efimovalex/brief_url/middlewares"
	"github.com/gorilla/mux"
)

// REST contains the rest based router and allows a listener to be set for easier, race-free testing
type REST struct {
	Addr   string
	Router *mux.Router
}

// StartHTTP listens on the configured ports for the REST application
func (r *REST) StartHTTP() error {
	if r.Addr == "" {
		return errors.New("address is required")
	}

	log.Printf("started http server: %s", r.Addr)

	n := negroni.New()
	n.Use(middlewares.NewCORS(r.Router))
	n.UseHandler(r.Router)

	n.Run(r.Addr)

	return nil
}
