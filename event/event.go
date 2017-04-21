package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Tympanix/automato/generate"
	"github.com/Tympanix/automato/subject"
	"github.com/Tympanix/automato/types"
)

// Trigger is an interfaces which describes the implementations needed for an event
type Trigger interface {
	types.IO
	types.Eventable
	Describe() string
}

// Event is a type which is used to trigger tasks
type Event struct {
	subject.Subject
	event types.Eventable
	UUID  string `json:"uuid"`
	Desc  string `json:"description"`
}

// Listen starts listening for events
func (e *Event) Listen() error {
	return e.event.Listen()
}

// ID returns the unique id of the event
func (e *Event) ID() string {
	return e.UUID
}

// Base is a struct used for subtyping to implement different events
// for the application
type Base struct {
	Identity  string           `json:"id"`
	Observers []types.Runnable `json:"-"`
	Event     string           `json:"event"`
}

// New takes an event and applies its type. The same event is returned.
func New(event Trigger) *Event {
	return &Event{
		Subject: *subject.New(event),
		event:   event,
		UUID:    generate.NewUUID(12),
		Desc:    event.Describe(),
	}
}

// UnmarshalJSON takes a byte array and uses type assertion to determine the type
// of event that was passed
func UnmarshalJSON(data []byte) (*Event, error) {
	m := make(map[string]interface{})
	json.Unmarshal(data, &m)

	eventString, ok := m["id"].(string)

	if !ok {
		return nil, errors.New("Could not parse event, no event type set")
	}

	event, ok := Templates[eventString]

	if !ok {
		return nil, fmt.Errorf("Event ”%s” is not a registered event type", eventString)
	}

	t := reflect.ValueOf(event)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	eventType := reflect.TypeOf(t.Interface())
	newEventInterface := reflect.New(eventType).Interface()

	newEvent, ok := newEventInterface.(Event)

	if !ok {
		return nil, fmt.Errorf("Internal error while parsing event")
	}

	err := json.Unmarshal(data, &newEvent)

	return &newEvent, err
}

func eventType(unit interface{}) string {
	t := reflect.TypeOf(unit)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.String()
}

func (e *Base) setEvent(event string) {
	e.Event = event
}

// Type returns the type of the event as a string representation
func (e *Base) Type() string {
	return e.Event
}

// Subscribe adds a new task to this event's observers
func (e *Base) Subscribe(task types.Runnable) {
	e.Observers = append(e.Observers, task)
}

// ID returns the unique id for the event
func (e *Base) ID() string {
	return e.Identity
}

func (e *Base) findObserverIndex(task types.Runnable) (int, error) {
	for i, t := range e.Observers {
		if task == t {
			return i, nil
		}
	}
	return -1, errors.New("Observer not found")
}

func (e *Base) removeObserver(index int) {
	e.Observers = append(e.Observers[:index], e.Observers[index+1:]...)
}

// Trigger exectures all the subscribed tasks for this event
func (e *Base) Trigger() {
	log.Printf("Triggered event %s\n", e.Event)
	for _, task := range e.Observers {
		task.Run()
	}
}

// Unsubscribe removes a task from this events observables
func (e *Base) Unsubscribe(task types.Runnable) error {
	i, err := e.findObserverIndex(task)
	if err != nil {
		return err
	}
	e.removeObserver(i)
	return nil
}
