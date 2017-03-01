package task

// Event interface describes an event
type Event interface {
	Trigger() error
	Output() interface{}
}
