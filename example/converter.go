package example

import (
	"fmt"

	"github.com/Tympanix/automato/hub"
)

// StringConverter formats a string using variables
type StringConverter struct {
	input struct {
		String       string
		Placeholders []interface{}
	}
	out struct {
		String string
	}
}

func init() {
	hub.Register(&StringConverter{})
}

// Input is the input of the converter
func (s *StringConverter) Input() interface{} {
	return &s.input
}

// Output is the output of the converter
func (s *StringConverter) Output() interface{} {
	return &s.out
}

// Convert function converts the string using the input and parameters
func (s *StringConverter) Convert() {
	s.out.String = fmt.Sprintf(s.input.String, s.input.Placeholders...)
}
