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

	ln.Info("started http server", ln.Map{"addr": a.Listener.Addr()})

	return http.Serve(a.Listener, a.Router)
}
