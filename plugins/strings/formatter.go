package string

import (
	"fmt"

	"github.com/Tympanix/artoodetoo/unit"
)

// Formatter formats a string using variables
type Formatter struct {
	Format      string      `io:"input"`
	Placeholder interface{} `io:"input"`

	Formatted string `io:"output"`
}

func init() {
	unit.Register(&Formatter{})
}

// Describe describes what a Formatter does
func (s *Formatter) Describe() string {
	return "Formats a string using a placeholder"
}

// Execute function converts the string using the input and parameters
func (s *Formatter) Execute() {
	s.Formatted = fmt.Sprintf(s.Format, s.Placeholder)
}
