package numbers

import "github.com/Tympanix/automato/unit"

// Addition event to test the application
type Addition struct {
	NumberA float32 `io:"input"`
	NumberB float32 `io:"input"`

	Result float32 `io:"output"`
}

func init() {
	unit.Register(&Addition{})
}

// Describe describes what a person event does
func (e *Addition) Describe() string {
	return "Perform addition on two numbers"
}

// Execute performs addition
func (e *Addition) Execute() {
	e.Result = e.NumberA + e.NumberB
}
