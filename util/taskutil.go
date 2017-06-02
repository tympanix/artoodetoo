package util

import (
	"github.com/Tympanix/artoodetoo/storage"
	"github.com/Tympanix/artoodetoo/task"
)

// AddTask adds a new tasks to application and saves it to storage
func AddTask(t *task.Task) error {
	if err := task.AddTask(t); err != nil {
		return err
	}
	if err := storage.SaveTask(t); err != nil {
		task.RemoveTask(t)
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
	if err := task.RemoveTask(t); err != nil {
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
