package string

import (
	"fmt"

	"github.com/Tympanix/automato/unit"
)

// StringFormatter formats a string using variables
type StringFormatter struct {
	Format      string      `io:"input"`
	Placeholder interface{} `io:"input"`

	Formatted string `io:"output"`
}

func init() {
	unit.Register(&StringFormatter{})
}

// Describe describes what a StringFormatter does
func (s *StringFormatter) Describe() string {
	return "Formats a string using a placeholder"
}

// Execute function converts the string using the input and parameters
func (s *StringFormatter) Execute() {
	s.Formatted = fmt.Sprintf(s.Format, s.Placeholder)
}
