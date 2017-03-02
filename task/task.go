package task

import (
	"fmt"
	"log"
)

// Task is an object that processes data based on events, converters and actions
type Task struct {
	Name    string
	Event   *Component
	Actions []*Component
}

// Describe prints our information about the action to the console
func (t *Task) Describe() {
	fmt.Printf("Task: %v\n", t.Name)
	fmt.Printf("Event: %v\n", t.Event)
	fmt.Printf("Actions:\n")

	for _, a := range t.Actions {
		fmt.Printf(" %s %v\n", "-", a)
	}

}

// Run starts the task
func (t *Task) Run() {
	state := make(State)

	t.Event.Execute()
	state.AddOutput(t.Event)

	for _, action := range t.Actions {
		log.Printf("Runnig action %v\n", action)
		state.GetInput(action)
		action.Execute()
		state.AddOutput(action)
	}

}
