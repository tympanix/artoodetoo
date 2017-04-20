package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/Tympanix/automato/task"
)

// Event is an interface for types that can listen on events
type Event interface {
	Listen() error
	ID() string
	Type() string
	setEvent(event string)
}

// UnmarshalJSON takes a byte array and uses type assertion to determine the type
// of event that was passed
func UnmarshalJSON(data []byte) (Event, error) {
	m := make(map[string]interface{})
	json.Unmarshal(data, &m)

	eventString, ok := m["event"].(string)

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

	return newEvent, err
}

// Base is a struct used for subtyping to implement different events
// for the application
type Base struct {
	id        string
	Observers []*task.Task `json:"-"`
	Event     string       `json:"event"`
}

func New(event Event) Event {
	event.setEvent(eventType(event))
	return event
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

func (e *Base) Type() string {
	return e.Event
}

// Subscribe adds a new task to this event's observers
func (e *Base) Subscribe(task *task.Task) {
	e.Observers = append(e.Observers, task)
}

// ID returns the unique id for the event
func (e *Base) ID() string {
	return e.id
}

func (e *Base) findObserverIndex(task *task.Task) (int, error) {
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
	for _, task := range e.Observers {
		task.Run()
	}
}

// Unsubscribe removes a task from this events observables
func (e *Base) Unsubscribe(task *task.Task) error {
	i, err := e.findObserverIndex(task)
	if err != nil {
		return err
	}
	e.removeObserver(i)
	return nil
}
