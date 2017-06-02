package numbers

import "github.com/Tympanix/artoodetoo/unit"

// Multiplication event to test the application
type Multiplication struct {
	NumberA float64 `io:"input"`
	NumberB float64 `io:"input"`

	Result float64 `io:"output"`
}

func init() {
	unit.Register(&Multiplication{})
}

// Describe describes what a person event does
func (e *Multiplication) Describe() string {
	return "Perform Multiplication on two numbers"
}

// Execute performs Multiplication
func (e *Multiplication) Execute() {
	e.Result = e.NumberA * e.NumberB
}
