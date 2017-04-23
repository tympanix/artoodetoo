package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Tympanix/automato/generate"
	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/subject"
	"github.com/Tympanix/automato/types"
)

// Trigger is an interfaces which describes the implementations needed for an event
type Trigger interface {
	types.IO
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

// Listen starts listening for events
func (e *Event) Listen() error {
	go e.trigger.Listen()
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

// New takes an event and applies its type. The same event is returned.
func New(trigger Trigger) *Event {
	return &Event{
		Subject:   *subject.New(trigger),
		trigger:   trigger,
		Observers: make([]types.Runnable, 0),
		UUID:      generate.NewUUID(12),
		Desc:      trigger.Describe(),
	}
}

// UnmarshalJSON serialized an event fron json encoding
func (e *Event) UnmarshalJSON(data []byte) error {
	type event Event
	var newEvent event
	if err := json.Unmarshal(data, &newEvent); err != nil {
		return err
	}

	*e = Event(newEvent)

	eventTemplate, ok := Templates[e.Type()]

	if !ok {
		return fmt.Errorf("Event ”%s” is not a registered event type", e.Type())
	}

	t := reflect.ValueOf(eventTemplate.trigger)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	newEventInterface := reflect.New(t.Type()).Interface()
	newTrigger, ok := newEventInterface.(Trigger)

	if !ok {
		return fmt.Errorf("Internal error while parsing event")
	}

	if err := e.BindIO(newTrigger); err != nil {
		return err
	}

	e.trigger = newTrigger

	if err := e.AssignInput(state.New()); err != nil {
		return err
	}

	return nil
}

func eventType(unit interface{}) string {
	t := reflect.TypeOf(unit)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.String()
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
