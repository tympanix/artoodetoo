package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Tympanix/automato/event"
	"github.com/Tympanix/automato/util"
)

func init() {
	API.HandleFunc("/all_events", listEventTemplates).Methods("GET")
	API.HandleFunc("/events", listEvents).Methods("GET")
	API.HandleFunc("/events", newEvent).Methods("POST")
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
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	event, err := event.UnmarshalJSON(data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := util.AddEvent(event); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
