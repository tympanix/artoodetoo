package task

import (
	"fmt"
	"log"

	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/unit"
)

// Task is an object that processes data based on events, converters and actions
type Task struct {
	Name    string       `json:"name"`
	Event   *unit.Unit   `json:"event"`
	Actions []*unit.Unit `json:"actions"`
}

// Describe prints our information about the action to the console
func (t *Task) Describe() {
	log.Printf("Task: %v\n", t.Name)
	log.Printf("Event: %v\n", t.Event)
	log.Printf("Actions:\n")

	for _, a := range t.Actions {
		log.Printf(" %s %v\n", "-", a)
	}
}

// GetUnitByName retrieves a unit in the actions list and returns it
func (t *Task) GetUnitByName(name string) (unit *unit.Unit, err error) {
	if t.Event.Name == name {
		return t.Event, nil
	}
	for _, u := range t.Actions {
		if u.Name == name {
			return u, nil
		}
	}
	err = fmt.Errorf("Task does not have unit with name '%s'", name)
	return
}

// Run starts the task
func (t *Task) Run() error {
	state := state.New()

	log.Printf("Running task %s\n", t.Name)

	t.Event.Execute()
	if err := t.Event.StoreOutput(state); err != nil {
		return err
	}

	for _, action := range t.Actions {
		if err := action.AssignInput(state); err != nil {
			return err
		}
		action.Execute()
		if err := action.StoreOutput(state); err != nil {
			return err
		}
	}
	return nil
}
