package example

import (
	"fmt"

	"github.com/Tympanix/automato/unit"
)

// StringConverter formats a string using variables
type StringConverter struct {
	Format      string      `io:"input"`
	Placeholder interface{} `io:"input"`

	Formatted string `io:"output"`
}

func init() {
	unit.Register(&StringConverter{})
}

// Describe describes what a stringconverter does
func (s *StringConverter) Describe() string {
	return "An example converter which formats a string using a placeholder"
}

// Execute function converts the string using the input and parameters
func (s *StringConverter) Execute() {
	s.Formatted = fmt.Sprintf(s.Format, s.Placeholder)
}
