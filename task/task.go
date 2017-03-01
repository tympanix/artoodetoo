package task

import (
	"fmt"
)

const (
	variable = iota
	static   = iota
	finish   = iota
)

// Task is an object that processes data based on events, converters and actions
type Task struct {
	Event   *Component
	Actions []Component
}

// Run starts the given task by connecting all components by pipes
// such that they cna comminucate with each other as a single workflow
func (t *Task) Run() {
	fmt.Println("Runnings task")
}
