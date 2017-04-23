package event

import (
	"encoding/json"
	"fmt"
)

// Proxy is a eventable object with custom json mashal/unmarshal methods
type Proxy struct {
	*Event
}

// MarshalJSON return the UUID of the event
func (p *Proxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.UUID)
}

// UnmarshalJSON returns the event specified by it's id
func (p *Proxy) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err != nil {
		return err
	}
	event, found := Events[id]
	if !found {
		return fmt.Errorf("Could not find event with id %s", id)
	}
	*p = Proxy{event}
	return nil
}
