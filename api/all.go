package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tympanix/automato/event"
	"github.com/Tympanix/automato/task"
	"github.com/Tympanix/automato/unit"
	"github.com/Tympanix/automato/util"
)

func init() {
	API.Handle("/all", auth(all)).Methods("GET")
}

// All is a struct which contains all information in the application
type All struct {
	Events    []*event.Event `json:"events"`
	EventTemp []*event.Event `json:"eventtemplates"`
	Tasks     []*task.Task   `json:"tasks"`
	Actions   []*unit.Unit   `json:"actions"`
}

func all(w http.ResponseWriter, r *http.Request) {
	SetJSON(w)

	all := All{
		Events:    util.AllEvents(),
		EventTemp: util.AllEventTemplates(),
		Tasks:     util.AllTasks(),
		Actions:   util.AllUnits(),
	}

	json.NewEncoder(w).Encode(&all)
}
