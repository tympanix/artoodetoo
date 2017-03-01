package hub

import (
	"log"

	"github.com/Tympanix/automato/task"
	"github.com/Tympanix/automato/task/action"
	"github.com/Tympanix/automato/task/component"
	"github.com/Tympanix/automato/task/event"
)

// Components contains all available components in the aplication
var Components []*component.Component

// Events contains all available events in the application
var Events []event.Event

// Converters contains all available converters in the application
var Converters []task.Converter

// Actions contains all available actions in the application
var Actions []action.Action

// Register is called to register a new component in the hub thus
// to make it public for use by the web app
func Register(comp interface{}) {
	switch t := comp.(type) {
	case event.Event:
		Events = append(Events, event.Factory(t))
	case task.Converter:
		Converters = append(Converters, t)
	case action.Action:
		Actions = append(Actions, t)
	default:
		log.Fatalf("Could not register component of type %T", t)
	}

	Components = append(Components, component.New(comp))
}
