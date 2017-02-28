package hub

import (
	"fmt"
	"log"
	"reflect"

	"github.com/Tympanix/automato/task"
)

// Events contains all available events in the application
var Events []task.IEvent
var converters []task.Converter

// Actions contains all available actions in the application
var Actions []task.IAction

// Register is called to register a new component in the hub thus
// to make it public for use by the web app
func Register(component interface{}) {
	applyID(component)
	switch t := component.(type) {
	case task.IEvent:
		Events = append(Events, t)
	case task.Converter:
		converters = append(converters, t)
	case task.IAction:
		Actions = append(Actions, t)
	default:
		log.Fatalf("Could not register component of type %T", t)
	}
}

func applyID(component interface{}) {
	if t, ok := component.(task.Identifier); ok {
		t.SetID(getNameOfStruct(t))
	}
}

func getNameOfStruct(component interface{}) string {
	t := reflect.TypeOf(component)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return fmt.Sprintf("%s", t.String())
}
