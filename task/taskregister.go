package task

import "errors"

var tasks map[string]*Task

func init() {
	tasks = make(map[string]*Task)
}

// AddTask registers a new task into the system
func AddTask(task *Task) error {
	if len(task.UUID) == 0 {
		return errors.New("Task has no UUID")
	}
	found, _ := GetTaskByID(task.UUID)
	if found != nil {
		return errors.New("Task with that name already exists")
	}
	tasks[task.Name] = task
	task.Subscribe()
	return nil
}

// RemoveTask removes a task from the application
func RemoveTask(task *Task) error {
	if _, err := GetTaskByID(task.UUID); err != nil {
		return errors.New("Could not unregister task because it was not found")
	}
	delete(tasks, task.Name)
	return nil
}

// Update updates the task
func Update(task *Task) error {
	_, err := GetTaskByID(task.UUID)
	if err != nil {
		return errors.New("Could not update task because it was not found")
	}
	tasks[task.Name] = task
	return nil
}

// GetTaskByName returns the task among the registered tasks where the name matches
func GetTaskByID(id string) (*Task, error) {
	task, ok := tasks[id]
	if ok {
		return task, nil
	}
	return nil, errors.New("Could not find task")
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
