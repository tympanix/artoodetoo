package example

import "github.com/Tympanix/automato/unit"

// PersonEvent event to test the application
type PersonEvent struct {
	unit.Base
	Person struct {
		Name    string
		Age     int
		Height  float32
		Married bool
	}
}

func init() {
	unit.Register(&PersonEvent{})
}

// Describe describes what a person event does
func (e *PersonEvent) Describe() string {
	return "An example event which, when triggered, outputs a sample person object"
}

// Execute creates a dummy event which output is a data collection of a person
func (e *PersonEvent) Execute() {
	e.Person.Name = "John Doe"
	e.Person.Age = 42
	e.Person.Height = 182
	e.Person.Married = true
}

// Output returs a person object
func (e *PersonEvent) Output() interface{} {
	return &e.Person
}
