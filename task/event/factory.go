package event

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Factory function that creates a new event wrapper based on the event
// passed into the factory. This will bootstrap the event, make it cloneable,
// give it unique identifier. The wrapper can afterwards be passed to the hub
func Factory(e Event) *Wrapper {
	return &Wrapper{
		id:    fmt.Sprintf("event%d", 1),
		event: e,
	}
}

// Wrapper wraps an event and extends it's functionality
type Wrapper struct {
	id    string
	event Event
}

// Output returns the output from the underlying event represented by the wrapper
func (w *Wrapper) Output() interface{} {
	return w.event.Output()
}

// ID returns the id for the events which is being wrapped
func (w *Wrapper) ID() string {
	return w.id
}

func (w *Wrapper) String() string {
	return w.ID()
}

// MarshalJSON returns a json representation of an event
func (w *Wrapper) MarshalJSON() ([]byte, error) {
	fmt.Println("Calling marshal")
	m := make(map[string]interface{})
	m["id"] = w.ID()

	output := make(map[string]string)

	s := reflect.ValueOf(w.Output())

	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		output[typeOfT.Field(i).Name] = f.Type().String()
	}

	m["output"] = output
	return json.Marshal(&m)
}
