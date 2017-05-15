package event

import "fmt"

type eventResolver struct{}

func (e *eventResolver) ResolveSubject(typ string) (interface{}, error) {
	eventTemplate, ok := Templates[typ]

	if !ok {
		return nil, fmt.Errorf("Event ”%s” is not a registered event type", typ)
	}

	return eventTemplate.trigger, nil
}
