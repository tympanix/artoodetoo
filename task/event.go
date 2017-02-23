package task

import "io"

// Event interface describes an object which can be triggered by some incident.
type Event interface {
	Trigger(*io.PipeWriter)
}
