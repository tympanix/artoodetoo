package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/Tympanix/automato/generate"
	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/subject"
	"github.com/Tympanix/automato/types"
)

// Core is an interfaces which is the trigger mechanism for an event
type Core interface {
	types.Triggerable
	Bind(e types.Triggerable)
	Describe() string
}

// Event is a type which is used to trigger tasks
type Event struct {
	subject.Subject
	Core      `json:"-"`
	Observers []types.Runnable `json:"-"`
	UUID      string           `json:"uuid"`
	Desc      string           `json:"description"`
	stop      chan struct{}    `json:"-"`
	running   bool             `json:"-"`
}

// New takes an event and applies its type. The same event is returned.
func New(trigger Core) *Event {
	e := &Event{
		Subject:   *subject.New(trigger, new(eventResolver)),
		Core:      trigger,
		Observers: make([]types.Runnable, 0),
		UUID:      generate.NewUUID(12),
		Desc:      trigger.Describe(),
	}
	e.init()
	return e
}

// ID returns the unique id of the event
func (e *Event) ID() string {
	return e.UUID
}

func (e *Event) init() {
	e.Bind(e)
	e.stop = make(chan struct{})
}

// Trigger returns the trigger for the event
func (e *Event) Trigger() {
	log.Printf("Triggered event %s\n", e.Type())
	for _, task := range e.Observers {
		state := state.New()
		e.StoreOutput(state)
		if err := task.Run(state); err != nil {
			log.Println(err)
		}
	}
}

func (e *Event) Start() {
	if e.running {
		return
	}
	go e.Listen(e.stop)
	e.running = true
}

func (e *Event) Stop() {
	if !e.running {
		return
	}
	e.stop <- struct{}{}
	e.running = false
}

// Validate checks certain safety proper
func (e *Event) Validate() error {
	for _, input := range e.In {
		for _, ingr := range input.Recipe {
			if ingr.IsVariable() {
				return errors.New("Events are not allowed to have variables as input")
			}
		}
	}
	return nil
}

// UnmarshalJSON serialized an event fron json encoding
func (e *Event) UnmarshalJSON(data []byte) error {
	type event Event
	newEvent := event(*e)
	if err := json.Unmarshal(data, &newEvent); err != nil {
		return err
	}

	*e = Event(newEvent)

	err := e.RebuildSubject(data, new(eventResolver))
	if err != nil {
		return err
	}

	newTrigger, ok := e.GetSubject().(Core)
	if !ok {
		return fmt.Errorf("Internal error while parsing event")
	}
	e.Core = newTrigger
	e.init()

	if err := e.Validate(); err != nil {
		return err
	}

	// Try to assign input to state, this should fail if any variables are used
	if err := e.AssignInput(state.New()); err != nil {
		return err
	}

	return nil
}

// GenerateUUDI creates a new UUID for the event
func (e *Event) GenerateUUDI() {
	e.UUID = generate.NewUUID(12)
}

// Subscribe adds a new task to this event's observers
func (e *Event) Subscribe(task types.Runnable) error {
	_, err := e.findObserverIndex(task)
	if err == nil {
		return errors.New("Runnable already subscribed to event")
	}
	e.Observers = append(e.Observers, task)
	return nil
}

func (e *Event) findObserverIndex(task types.Runnable) (int, error) {
	for i, t := range e.Observers {
		if task == t {
			return i, nil
		}
	}
	return -1, errors.New("Observer not found")
}

func (e *Event) removeObserver(index int) {
	e.Observers = append(e.Observers[:index], e.Observers[index+1:]...)
}

// Unsubscribe removes a task from this events observables
func (e *Event) Unsubscribe(task types.Runnable) error {
	i, err := e.findObserverIndex(task)
	if err != nil {
		return err
	}
	e.removeObserver(i)
	return nil
}
