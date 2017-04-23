package event

import (
	"fmt"
	"reflect"
)

type eventResolver struct{}

func (e *eventResolver) ResolveSubject(typ string) (interface{}, error) {
	eventTemplate, ok := Templates[typ]

	if !ok {
		return nil, fmt.Errorf("Event ”%s” is not a registered event type", typ)
	}

	t := reflect.ValueOf(eventTemplate.trigger)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	newEventInterface := reflect.New(t.Type()).Interface()
	newTrigger, ok := newEventInterface.(Trigger)

	if !ok {
		return nil, fmt.Errorf("Internal error while parsing event")
	}

	return newTrigger, nil
}
