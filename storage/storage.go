package storage

import (
	"errors"
	"log"

	"github.com/Tympanix/automato/task"
)

// Driver holds the active storage manager for the application
var Driver Store

// Store interface is an object which can store tasks for consistency and Could
// be a driver for a database, fileserver or any other storage mechanism
type Store interface {
	SaveTask(*task.Task) error
	GetAllTasks() ([]*task.Task, error)
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
func Load() int {
	tasks, err := GetAllTasks()
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tasks {
		err := task.Register(t)
		if err != nil {
			log.Fatal(err)
		}
	}
	return len(tasks)
}

// GetAllTasks returns all tasks saved in the storage manager
func GetAllTasks() ([]*task.Task, error) {
	if err := hasDriver(); err != nil {
		return nil, err
	}
	return Driver.GetAllTasks()
}

func hasDriver() error {
	if Driver == nil {
		return errors.New("No storage drive specified")
	}
	return nil
}
