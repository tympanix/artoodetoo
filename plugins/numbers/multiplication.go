package numbers

import "github.com/Tympanix/automato/unit"

// Multiplication event to test the application
type Multiplication struct {
	NumberA int `io:"input"`
	NumberB int `io:"input"`

	Result int `io:"output"`
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
