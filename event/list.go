package event

import "encoding/json"

// List is a list of events that supports JSON serialization
type List []Event

// UnmarshalJSON uses reflection to determine the actual type of each event
func (l *List) UnmarshalJSON(data []byte) error {
	var rawEvents []*json.RawMessage
	err := json.Unmarshal(data, &rawEvents)
	if err != nil {
		return err
	}

	var events []Event
	for _, raw := range rawEvents {
		event, err := UnmarshalJSON([]byte(*raw))
		if err != nil {
			return err
		}
		events = append(events, event)
	}
	*l = events
	return nil
}
