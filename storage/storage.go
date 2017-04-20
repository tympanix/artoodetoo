package storage

import (
	"errors"
	"log"

	"github.com/Tympanix/automato/event"
	"github.com/Tympanix/automato/task"
)

// Driver holds the active storage manager for the application
var Driver Store

// Store interface is an object which can store tasks for consistency and Could
// be a driver for a database, fileserver or any other storage mechanism
type Store interface {
	SaveTask(*task.Task) error
	DeleteTask(*task.Task) error
	UpdateTask(*task.Task) error
	GetAllTasks() ([]*task.Task, error)
	GetAllEvents() ([]event.Event, error)
	SaveEvent(event.Event) error
}

// Register sets a new Store as the current storage method
func Register(store Store) {
	Driver = store
}

// SaveTask saves a new task in the storage manager
func SaveTask(task *task.Task) error {
	if err := hasDriver(); err != nil {
		return err
	}
	return Driver.SaveTask(task)
}

// Load loads the saved tasks and registers them into the application
func Load() (int, int) {
	tasks, err := GetAllTasks()
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tasks {
		err = task.Register(t)
		if err != nil {
			log.Fatal(err)
		}
	}

	events, err := GetAllEvents()
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range events {
		err := event.AddEvent(e)
		if err != nil {
			log.Fatal(err)
		}
	}
	return len(tasks), len(events)
}

// SaveEvent saves an event in the storage manager
func SaveEvent(event event.Event) error {
	if err := hasDriver(); err != nil {
		return err
	}
	return Driver.SaveEvent(event)
}

// GetAllTasks returns all tasks saved in the storage manager
func GetAllTasks() ([]*task.Task, error) {
	if err := hasDriver(); err != nil {
		return nil, err
	}
	return Driver.GetAllTasks()
}

// GetAllEvents returns all events saved in the storage manager
func GetAllEvents() ([]event.Event, error) {
	if err := hasDriver(); err != nil {
		return nil, err
	}
	return Driver.GetAllEvents()
}

// DeleteTask uses the current storage manager to delete a task
func DeleteTask(task *task.Task) error {
	if err := hasDriver(); err != nil {
		return err
	}
	if err := Driver.DeleteTask(task); err != nil {
		return err
	}
	return nil
}

// UpdateTask uses the current storage manager to update a task
func UpdateTask(task *task.Task) error {
	if err := hasDriver(); err != nil {
		return err
	}
	if err := Driver.UpdateTask(task); err != nil {
		return err
	}
	return nil
}

func hasDriver() error {
	if Driver == nil {
		return errors.New("No storage drive specified")
	}
	return nil
}
