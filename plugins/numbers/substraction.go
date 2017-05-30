package numbers

import "github.com/Tympanix/automato/unit"

// Substraction event to test the application
type Substraction struct {
	NumberA float64 `io:"input"`
	NumberB float64 `io:"input"`

	Result float64 `io:"output"`
}

func init() {
	unit.Register(&Substraction{})
}

// Describe describes what a person event does
func (e *Substraction) Describe() string {
	return "Perform Substraction on two numbers, A - B"
}

// Execute performs Substraction
func (e *Substraction) Execute() {
	e.Result = e.NumberA - e.NumberB
}
