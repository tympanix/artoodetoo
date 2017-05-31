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
	UUID    string                `json:"uuid"`
	Name    string                `json:"name"`
	Event   *event.Proxy          `json:"event"`
	Actions []*unit.Unit          `json:"actions"`
	Queue   chan types.TupleSpace `json:"-"`
	running *sync.Mutex           `json:"-"`
	once    *sync.Once            `json:"-"`
}

func (t *Task) init() {
	t.Queue = make(chan types.TupleSpace, 1<<12)
	t.running = new(sync.Mutex)
	t.once = new(sync.Once)
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

func (t *Task) Validate() error {
	if err := t.detectCycles(); err != nil {
		return err
	}
	return nil
}

// Subscribe subscribes the task to its event
func (t *Task) Subscribe() error {
	t.Queue = make(chan types.TupleSpace, 1<<12)
	t.once = new(sync.Once)
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
	if err := t.Event.Unsubscribe(t); err != nil {
		return err
	}
	close(t.Queue)
	return nil
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
	t.Queue <- ts
	t.once.Do(t.startWorker)
	return nil
}

func (t *Task) startWorker() {
	go func() {
		for ts := range t.Queue {
			t.run(ts)
		}
	}()
}

func (t *Task) run(ts types.TupleSpace) {
	log.Printf("Running task %s\n", t.Name)

	t.running.Lock()
	defer t.running.Unlock()

	waitgroup := new(sync.WaitGroup)
	waitgroup.Add(len(t.Actions))

	for _, action := range t.Actions {
		action.RunAsync(waitgroup, ts)
	}

	waitgroup.Wait()
	log.Printf("Finished %s", t.Name)
}

func (t *Task) UnmarshalJSON(data []byte) error {
	type task Task
	var _task task
	if err := json.Unmarshal(data, &_task); err != nil {
		return err
	}

	*t = Task(_task)
	t.init()

	if err := t.Validate(); err != nil {
		return err
	}

	return nil
}
