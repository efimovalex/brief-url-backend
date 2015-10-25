package app

import (
	"net/http"
)

// DomainEndpoint exists as an example
type DomainEndpoint struct {
}

func (ep *DomainEndpoint) Get(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user_id"]
	if userID == "" {
		errs := []clientresponses.ClientErrorItem{{Message: "missing user_id", Field: ":user_id"}}
		handleErrs(w, http.StatusBadRequest, errs)

		return
	}

	domainID := mux.Vars(r)["domain_id"]
	if domainID == "" {

	} else {

	}

	w.WriteHeader(http.StatusOK)
}
