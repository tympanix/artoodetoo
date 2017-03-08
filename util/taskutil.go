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

// DeleteTask removes the task from the storage manager and unregisters the task
func DeleteTask(t *task.Task) error {
	if err := storage.DeleteTask(t); err != nil {
		return err
	}
	if err := task.Unregister(t); err != nil {
		return err
	}
	return nil
}

// UpdateTask updates the task in the storage manager and in memory
func UpdateTask(t *task.Task) error {
	if err := storage.UpdateTask(t); err != nil {
		return err
	}
	if err := task.Update(t); err != nil {
		return err
	}
	return nil
}
