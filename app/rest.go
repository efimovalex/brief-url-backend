package app

import (
	"errors"
	"log"
	"net"
	"net/http"
)

// REST contains the rest based router and allows a listener to be set for easier, race-free testing
type REST struct {
	Router   http.Handler
	Listener net.Listener
}

// StartHTTP listens on the configured ports for the REST application
func (a *REST) StartHTTP() error {
	if a.Listener == nil {
		return errors.New("listener is required")
	}

	log.Printf("started http server: %s", a.Listener.Addr().String())

	return http.Serve(a.Listener, a.Router)
}
