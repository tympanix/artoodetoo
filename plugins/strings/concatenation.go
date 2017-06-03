package string

import "github.com/Tympanix/artoodetoo/unit"

// Concatenation formats a string using variables
type Concatenation struct {
	StringA string `io:"input"`
	StringB string `io:"input"`

	Result string `io:"output"`
}

func init() {
	unit.Register(&Concatenation{})
}

// Describe describes what a Concatenation does
func (s *Concatenation) Describe() string {
	return "Concatenates string A with string B"
}

// Execute function concatenates strings
func (s *Concatenation) Execute() error {
	s.Result = s.StringA + s.StringB
	return nil
}
