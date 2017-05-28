package numbers

import "github.com/Tympanix/automato/unit"

// Substraction event to test the application
type Substraction struct {
	NumberA int `io:"input"`
	NumberB int `io:"input"`

	Result int `io:"output"`
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
