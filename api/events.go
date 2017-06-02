package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tympanix/artoodetoo/event"
	"github.com/Tympanix/artoodetoo/util"
)

func init() {
	API.Handle("/all_events", auth(listEventTemplates)).Methods("GET")
	API.Handle("/events", auth(listEvents)).Methods("GET")
	API.Handle("/events", auth(newEvent)).Methods("POST")
}

func listEventTemplates(w http.ResponseWriter, r *http.Request) {
	SetJSON(w)
	json.NewEncoder(w).Encode(util.AllEventTemplates())
}

func listEvents(w http.ResponseWriter, r *http.Request) {
	SetJSON(w)
	json.NewEncoder(w).Encode(util.AllEvents())
}

func newEvent(w http.ResponseWriter, r *http.Request) {
	var event event.Event
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	event.GenerateUUDI()
	if err := util.AddEvent(&event); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
