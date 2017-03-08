package storage

import "github.com/Tympanix/automato/task"

// Driver holds the active storage options for the application
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
