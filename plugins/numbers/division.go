package numbers

import (
	"errors"

	"github.com/Tympanix/artoodetoo/unit"
)

// Division event to test the application
type Division struct {
	NumberA float64 `io:"input"`
	NumberB float64 `io:"input"`

	Result float64 `io:"output"`
}

func init() {
	unit.Register(&Division{})
}

// Describe describes what a person event does
func (e *Division) Describe() string {
	return "Perform division on two numbers"
}

// Execute performs Multiplication
func (e *Division) Execute() error {
	if e.NumberB == 0.0 {
		return errors.New("Division by zero")
	}
	e.Result = e.NumberA / e.NumberB
	return nil
}
