package util

import (
	"github.com/Tympanix/automato/storage"
	"github.com/Tympanix/automato/task"
)

// AddTask adds a new tasks to application and saves it to storage
func AddTask(t *task.Task) error {
	if err := task.Register(t); err != nil {
		return err
	}
	if err := storage.SaveTask(t); err != nil {
		task.Unregister(t)
		return err
	}
	return nil
}

// AllTasks return all tasks in a list
func AllTasks() []*task.Task {
	return task.All()
}
