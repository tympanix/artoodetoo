package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tympanix/artoodetoo/event"
	"github.com/Tympanix/artoodetoo/task"
	"github.com/Tympanix/artoodetoo/unit"
	"github.com/Tympanix/artoodetoo/util"
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
