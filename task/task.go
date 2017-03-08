package task

import (
	"fmt"
	"log"

	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/unit"
)

// Task is an object that processes data based on events, converters and actions
type Task struct {
	Name    string
	Event   *unit.Unit
	Actions []*unit.Unit
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
	state := state.New()

	t.Event.Execute()
	state.StoreStruct(t.Event.Name, t.Event.Output())

	for _, action := range t.Actions {
		log.Printf("Runnig action %v\n", action)
		action.AssignInput(state)
		action.Execute()
		action.StoreOutput(state)
	}

}
