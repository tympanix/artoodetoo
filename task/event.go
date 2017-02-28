package task

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// IEvent interface describes an event
type IEvent interface {
	Output() interface{}
	SetOutput(out interface{})
}

// Event describes an object which can be triggered by some incident.
type Event struct {
	id     string
	output interface{}
	Name   string
	Test   interface{}
}

// Trigger triggers the event
func (e *Event) Trigger() {
	log.Fatal("Trigger not implemented!")
}

// Output returns the output from the event
func (e *Event) Output() interface{} {
	return e.output
}

// SetOutput sets the reference to the outout of the event
func (e *Event) SetOutput(out interface{}) {
	e.output = out
}

// SetID sets the identification for the event
func (e *Event) SetID(id string) {
	e.id = id
}

// GetID returns the identification of the event
func (e *Event) GetID() string {
	return e.id
}

func (e *Event) String() string {
	return e.GetID()
}

// MarshalJSON returns a json representation of an event
func (e *Event) MarshalJSON() ([]byte, error) {
	fmt.Println("Calling marshal")
	m := make(map[string]interface{})
	m["id"] = e.GetID()

	output := make(map[string]string)

	s := reflect.ValueOf(e.Output())

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
