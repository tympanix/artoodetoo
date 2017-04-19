package util

import (
	"github.com/Tympanix/automato/event"
	"github.com/Tympanix/automato/storage"
)

// AddEvent adds a new event to application and saves it to storage
func AddEvent(t event.Event) error {
	if err := event.AddEvent(t); err != nil {
		return err
	}
	if err := storage.SaveEvent(t); err != nil {
		event.RemoveEvent(t)
		return err
	}
	return nil
}
