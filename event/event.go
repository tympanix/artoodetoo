package event

import (
	"errors"
	"reflect"

	"github.com/Tympanix/automato/task"
)

// Event is an interface for types that can listen on events
type Event interface {
	Listen() error
	ID() string
}

// Base is a struct used for subtyping to implement different events
// for the application
type Base struct {
	Observers []*task.Task `json:"-"`
	Event     string       `json:"event"`
}

// Subscribe adds a new task to this event's observers
func (e *Base) Subscribe(task *task.Task) {
	e.Observers = append(e.Observers, task)
}

// ID returns the unique id for the event
func (e *Base) ID() string {
	t := reflect.TypeOf(e)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.String()
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
