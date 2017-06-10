package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tympanix/artoodetoo/event"
	"github.com/Tympanix/artoodetoo/util"
	"github.com/gorilla/mux"
)

func init() {
	API.Handle("/all_events", auth(listEventTemplates)).Methods("GET")
	API.Handle("/events", auth(listEvents)).Methods("GET")
	API.Handle("/events", auth(newEvent)).Methods("POST")
	API.Handle("/events/{event}/stop", auth(stopEvent)).Methods("POST")
	API.Handle("/events/{event}/start", auth(startEvent)).Methods("POST")
}

func stopEvent(w http.ResponseWriter, r *http.Request) {
	e, ok := getEvent(w, r)
	if !ok {
		return
	}
	if err := e.Stop(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func startEvent(w http.ResponseWriter, r *http.Request) {
	e, ok := getEvent(w, r)
	if !ok {
		return
	}
	if err := e.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getEvent(w http.ResponseWriter, r *http.Request) (*event.Event, bool) {
	vars := mux.Vars(r)
	eventid, ok := vars["event"]
	if !ok {
		http.Error(w, "Unspecified event", http.StatusInternalServerError)
		return nil, false
	}
	e, err := event.GetEventByID(eventid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, false
	}
	return e, true
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
