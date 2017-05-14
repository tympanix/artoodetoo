package task

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/Tympanix/automato/event"
	"github.com/Tympanix/automato/generate"
	"github.com/Tympanix/automato/types"
	"github.com/Tympanix/automato/unit"
)

// Task is an object that processes data based on events, converters and actions
type Task struct {
	UUID    string       `json:"uuid"`
	Name    string       `json:"name"`
	Event   *event.Proxy `json:"event"`
	Actions []*unit.Unit `json:"actions"`
	lock    *sync.Mutex  `json:"-"`
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
		return fmt.Errorf("Task %s has no event", t.Name)
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
func (t *Task) Run(ts types.TupleSpace) error {
	log.Printf("Running task %s\n", t.Name)

	t.lock.Lock()
	defer t.lock.Unlock()

	waitgroup := new(sync.WaitGroup)
	waitgroup.Add(len(t.Actions))

	for _, action := range t.Actions {
		go func(action *unit.Unit) {
			if err := action.AssignInput(ts); err != nil {
				log.Println(err)
			}
			action.Execute()
			if err := action.StoreOutput(ts); err != nil {
				log.Println(err)
			}
			waitgroup.Done()
		}(action)
	}

	waitgroup.Wait()
	log.Printf("Finished %s", t.Name)
	return nil
}

func (t *Task) UnmarshalJSON(data []byte) error {
	type task Task
	var _task task
	if err := json.Unmarshal(data, &_task); err != nil {
		return err
	}

	_task.lock = new(sync.Mutex)
	*t = Task(_task)
	return nil
}
