package example

import "github.com/Tympanix/automato/unit"

// PersonEvent event to test the application
type PersonEvent struct {
	unit.Base
	Name    string  `io:"output"`
	Age     int     `io:"output"`
	Height  float32 `io:"output"`
	Married bool    `io:"output"`
}

func init() {
	unit.Register(&PersonEvent{})
}

// Describe describes what a person event does
func (e *PersonEvent) Describe() string {
	return "An example event which, when triggered, outputs a sample person object"
}

// Execute creates a dummy event which output is a data collection of a person
func (e *PersonEvent) Execute() {
	e.Name = "John Doe"
	e.Age = 42
	e.Height = 182
	e.Married = true
}
