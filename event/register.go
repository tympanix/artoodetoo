package event

import "errors"

// Templates containes sample events as a preview to the user
var Templates map[string]*Event

// Events contains the registered events in the application
var Events map[string]*Event

func init() {
	Templates = make(map[string]*Event)
	Events = make(map[string]*Event)
}

// Register registers events as templates for the user
func Register(trigger Trigger) {
	newEvent := New(trigger)
	newEvent.UUID = ""
	Templates[newEvent.Type()] = newEvent
}

// AddEvent adds an event to the application
func AddEvent(event *Event) error {
	_, found := Events[event.ID()]
	if found {
		return errors.New("Event with that id already exists")
	}
	Events[event.ID()] = event
	go event.Listen()
	return nil
}

// RemoveEvent removes an evenet from the application
func RemoveEvent(event *Event) {
	delete(Events, event.ID())
}
