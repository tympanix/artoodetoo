package task

import (
	"fmt"
	"log"

	"github.com/Tympanix/automato/event"
	"github.com/Tympanix/automato/generate"
	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/unit"
)

// Task is an object that processes data based on events, converters and actions
type Task struct {
	UUID    string       `json:"uuid"`
	Name    string       `json:"name"`
	Event   *event.Proxy `json:"event"`
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

// Subscribe subscribes the task to its event
func (t *Task) Subscribe() error {
	if t.Event == nil {
		return fmt.Errorf("Task %s has no event to subscribe to", t.Name)
	}
	return t.Event.Subscribe(t)
}

// Unsubscribe removed this task as an observer for its event
func (t *Task) Unsubscribe() error {
	if t.Event == nil {
		return fmt.Errorf("Task %s has no event to subscribe to", t.Name)
	}
	return t.Event.Unsubscribe(t)
}

func (t *Task) GenerateUUID() {
	t.UUID = generate.NewUUID(12)
}

// GetUnitByName retrieves a unit in the actions list and returns it
func (t *Task) GetUnitByName(name string) (unit *unit.Unit, err error) {
	for _, u := range t.Actions {
		if u.Name == name {
			return u, nil
		}
	}
	err = fmt.Errorf("Task does not have unit with name '%s'", name)
	return
}

// Run starts the task
func (t *Task) Run(state state.State) error {
	log.Printf("Running task %s\n", t.Name)

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
