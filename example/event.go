package example

import "github.com/Tympanix/automato/task"

// PersonEvent event to test the application
type PersonEvent struct {
	task.Base
	Person struct {
		Name    string
		Age     int
		Heigth  float32
		Married bool
	}
}

func init() {
	task.Register(&PersonEvent{})
}

// Execute creates a dummy event which output is a data collection of a person
func (e *PersonEvent) Execute() {
	e.Person.Name = "John Doe"
	e.Person.Age = 42
	e.Person.Heigth = 182
	e.Person.Married = true
}

// Output returs a person object
func (e *PersonEvent) Output() interface{} {
	return &e.Person
}
