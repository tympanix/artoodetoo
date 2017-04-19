package event

import "errors"

// Templates containes sample events as a preview to the user
var Templates []Event

// Events contains the registered events in the application
var Events map[string]Event

// Register registers events as templates for the user
func Register(event Event) {
	Templates = append(Templates, event)
}

// AddEvent adds an event to the application
func AddEvent(event Event) error {
	_, found := Events[event.ID()]
	if !found {
		return errors.New("Event with that id already exists")
	}
	Events[event.ID()] = event
	return nil
}

// RemoveEvent removes an evenet from the application
func RemoveEvent(event Event) {
	delete(Events, event.ID())
}
