package util

import (
	"github.com/Tympanix/automato/event"
	"github.com/Tympanix/automato/storage"
)

// AddEvent adds a new event to application and saves it to storage
func AddEvent(t *event.Event) error {
	if err := event.AddEvent(t); err != nil {
		return err
	}
	if err := storage.SaveEvent(t); err != nil {
		event.RemoveEvent(t)
		return err
	}
	return nil
}

// AllEventTemplates returns all available events in the application
func AllEventTemplates() []*event.Event {
	events := make([]*event.Event, len(event.Templates))
	idx := 0
	for _, v := range event.Templates {
		events[idx] = v
		idx++
	}
	return events
}

// AllEvents returns all user created events in the application
func AllEvents() []*event.Event {
	events := make([]*event.Event, len(event.Events))
	idx := 0
	for _, v := range event.Events {
		events[idx] = v
		idx++
	}
	return events
}
