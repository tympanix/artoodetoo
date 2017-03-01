package hub

import (
	"log"

	"github.com/Tympanix/automato/task"
	"github.com/Tympanix/automato/task/event"
)

// Events contains all available events in the application
var Events []event.Event

// Converters contains all available converters in the application
var Converters []task.Converter

// Actions contains all available actions in the application
var Actions []task.IAction

// Register is called to register a new component in the hub thus
// to make it public for use by the web app
func Register(component interface{}) {
	switch t := component.(type) {
	case event.Event:
		Events = append(Events, event.Factory(t))
	case task.Converter:
		Converters = append(Converters, t)
	case task.IAction:
		Actions = append(Actions, t)
	default:
		log.Fatalf("Could not register component of type %T", t)
	}
}
