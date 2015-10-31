package middlewares

import (
	"net/http"

	"github.com/gorilla/mux"
)

// CORS maps the stats client used and the router
type CORS struct {
	router *mux.Router
}

// NewStatsDTiming return a instance of statsd timing middleware
func NewCORS(router *mux.Router) *CORS {
	return &CORS{router: router}
}

// ServeHTTP overwrites the http method in order to send stats
func (c CORS) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if origin := r.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	next(rw, r)
}
