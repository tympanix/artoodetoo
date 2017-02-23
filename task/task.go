package task

import "io"

// Task is an object that processes data based on events, converters and actions
type Task struct {
	Event      Event
	Converters []Converter
	Action     Action
}

// Run starts the given task by connecting all components by pipes
// such that they cna comminucate with each other as a single workflow
func (t *Task) Run() {
	oldR, oldW := io.Pipe()
	go t.Event.Trigger(oldW)

	var newR *io.PipeReader
	var newW *io.PipeWriter
	for _, cnv := range t.Converters {
		newR, newW = io.Pipe()
		go cnv.Convert(oldR, newW)
		oldR = newR
	}
	t.Action.Execute(newR)
}
