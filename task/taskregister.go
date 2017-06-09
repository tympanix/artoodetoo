package task

import (
	"errors"
	"fmt"
)

var tasks map[string]*Task

func init() {
	tasks = make(map[string]*Task)
}

// AddTask registers a new task into the system
func AddTask(task *Task) error {
	if len(task.UUID) == 0 {
		return errors.New("Task has no UUID")
	}
	if found, _ := GetTaskByID(task.UUID); found != nil {
		return errors.New("Task with that id already exists")
	}
	if _, err := GetTaskByName(task.Name); err == nil {
		return errors.New("Task with that name already exists")
	}
	tasks[task.UUID] = task
	task.Subscribe()
	return nil
}

// RemoveTask removes a task from the application
func RemoveTask(task *Task) error {
	if _, err := GetTaskByID(task.UUID); err != nil {
		return errors.New("Could not unregister task because it was not found")
	}
	task.Unsubscribe()
	delete(tasks, task.UUID)
	return nil
}

// Update updates the task
func Update(task *Task) error {
	old, err := GetTaskByID(task.UUID)
	if err != nil {
		return errors.New("Could not update task because it was not found")
	}
	if err := old.Unsubscribe(); err != nil {
		return err
	}
	if err := task.Subscribe(); err != nil {
		return err
	}
	tasks[task.UUID] = task
	return nil
}

// GetTaskByID returns the task among the registered tasks where the id matches
func GetTaskByID(id string) (*Task, error) {
	task, ok := tasks[id]
	if ok {
		return task, nil
	}
	return nil, errors.New("Could not find task")
}

// GetTaskByName retrieves a task by its name
func GetTaskByName(name string) (task *Task, err error) {
	for _, task = range tasks {
		if task.Name == name {
			return
		}
	}
	err = fmt.Errorf("Task with name %s not found", name)
	return
}

// All return all registered tasks in a list
func All() []*Task {
	all := make([]*Task, len(tasks))
	idx := 0
	for _, v := range tasks {
		all[idx] = v
		idx++
	}
	return all
}
