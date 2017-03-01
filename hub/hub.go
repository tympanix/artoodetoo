package hub

import (
	"log"

	"github.com/Tympanix/automato/task"
)

// Components contains all available components in the aplication
var Components []*task.Component

// Events contains all available events in the application
var Events []task.Event

// Converters contains all available converters in the application
var Converters []task.Converter

// Actions contains all available actions in the application
var Actions []task.Action

// Register is called to register a new component in the hub thus
// to make it public for use by the web app
func Register(comp interface{}) {
	switch t := comp.(type) {
	case task.Event:
		Events = append(Events, t)
	case task.Converter:
		Converters = append(Converters, t)
	case task.Action:
		Actions = append(Actions, t)
	default:
		log.Fatalf("Could not register component of type %T", t)
	}

	Components = append(Components, task.NewComponent(comp))
}
