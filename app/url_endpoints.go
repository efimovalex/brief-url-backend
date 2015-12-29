package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/efimovalex/brief_url/adaptor/db"
	"github.com/efimovalex/brief_url/client"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// URLEndpoints exists as an example
type URLEndpoints struct {
	DB     *db.Adaptor
	Logger *log.Logger
}

func (ue *URLEndpoints) Get(w http.ResponseWriter, r *http.Request) {
	URLID := mux.Vars(r)["url_id"]
	if URLID == "" {
		URLs, err := ue.DB.Url.GetAll()
		if err != nil {
			errs := []client.Error{{Message: err.Error()}}
			handleErrs(w, http.StatusInternalServerError, errs)

			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(URLs); err != nil {
			ue.Logger.Println("problem encoding urls")
			errs := []client.Error{{Message: err.Error()}}
			handleErrs(w, http.StatusInternalServerError, errs)

			return
		}

		return
	} else {

		objID := bson.ObjectIdHex(URLID)
		URL, err := ue.DB.Url.Get(objID)
		if err != nil {
			errs := []client.Error{{Message: err.Error()}}
			handleErrs(w, http.StatusInternalServerError, errs)

			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(URL); err != nil {
			ue.Logger.Println("problem encoding url")
			errs := []client.Error{{Message: err.Error()}}
			handleErrs(w, http.StatusInternalServerError, errs)

			return
		}

		return
	}
}

func (ue *URLEndpoints) Post(w http.ResponseWriter, r *http.Request) {
	body, ioErr := ioutil.ReadAll(r.Body)
	if ioErr != nil {
		// attempt to read the body passed in
		captured := make([]byte, 0, 500)
		_, err := r.Body.Read(captured)
		if err != nil {
			ue.Logger.Println("io util error while parsing request body")
		}
		errs := []client.Error{{Message: "invalid body"}}
		handleErrs(w, http.StatusBadRequest, errs)

		return
	}

	url := db.URL{}
	err := json.Unmarshal(body, &url)
	if err != nil {
		errs := []client.Error{{Message: err.Error()}}
		handleErrs(w, http.StatusInternalServerError, errs)

		return
	}

	err = ue.DB.Url.Add(&url)
	if err != nil {
		errs := []client.Error{{Message: err.Error()}}
		handleErrs(w, http.StatusInternalServerError, errs)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ue *URLEndpoints) Delete(w http.ResponseWriter, r *http.Request) {
	URLID := mux.Vars(r)["url_id"]
	if URLID == "" {
		errs := []client.Error{{Message: "missing url_id", Field: ":url_id"}}
		handleErrs(w, http.StatusBadRequest, errs)

		return
	}

	err := ue.DB.Url.Delete(URLID)
	if err != nil {
		errs := []client.Error{{Message: err.Error()}}
		handleErrs(w, http.StatusInternalServerError, errs)

		return
	}

	w.WriteHeader(http.StatusOK)

	return
}

func (ue *URLEndpoints) Patch(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user_id"]
	if userID == "" {
		errs := []client.Error{{Message: "missing user_id", Field: ":user_id"}}
		handleErrs(w, http.StatusBadRequest, errs)

		return
	}
}
