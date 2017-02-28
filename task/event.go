package task

import "log"

// IEvent interface describes an event
type IEvent interface {
	GetOutput() *interface{}
}

// Event describes an object which can be triggered by some incident.
type Event struct {
	ID     string
	Output interface{}
}

// Trigger triggers the event
func (e *Event) Trigger() {
	log.Fatal("Trigger not implemented!")
}

// GetOutput returns the output from the event
func (e *Event) GetOutput() *interface{} {
	return &e.Output
}

// SetID sets the identification for the event
func (e *Event) SetID(id string) {
	e.ID = id
}

// GetID returns the identification of the event
func (e *Event) GetID() string {
	return e.ID
}
