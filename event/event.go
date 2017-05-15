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

// Trigger is an interfaces which describes the implementations needed for an event
type Trigger interface {
	types.Triggerable
	Describe() string
}

// Event is a type which is used to trigger tasks
type Event struct {
	subject.Subject
	trigger   Trigger
	Observers []types.Runnable `json:"-"`
	UUID      string           `json:"uuid"`
	Desc      string           `json:"description"`
}

// New takes an event and applies its type. The same event is returned.
func New(trigger Trigger) *Event {
	return &Event{
		Subject:   *subject.New(trigger, new(eventResolver)),
		trigger:   trigger,
		Observers: make([]types.Runnable, 0),
		UUID:      generate.NewUUID(12),
		Desc:      trigger.Describe(),
	}
}

// Listen starts listening for events
func (e *Event) Listen() error {
	go func() {
		if err := e.trigger.Listen(); err != nil {
			log.Println(err)
		}
	}()
	for goon := range e.Trigger() {
		if goon {
			e.Fire()
		} else {
			break
		}
	}
	return nil
}

// ID returns the unique id of the event
func (e *Event) ID() string {
	return e.UUID
}

// Trigger returns the trigger for the event
func (e *Event) Trigger() chan bool {
	return e.trigger.Trigger()
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

	newTrigger, ok := e.GetSubject().(Trigger)
	if !ok {
		return fmt.Errorf("Internal error while parsing event")
	}
	e.trigger = newTrigger

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

// Fire exectures all the subscribed tasks for this event
func (e *Event) Fire() {
	log.Printf("Triggered event %s\n", e.Type())
	for _, task := range e.Observers {
		state := state.New()
		e.StoreOutput(state)
		if err := task.Run(state); err != nil {
			log.Println(err)
		}
	}
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
