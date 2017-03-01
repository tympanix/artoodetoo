package event

// Event interface describes an event
type Event interface {
	Output() interface{}
}
