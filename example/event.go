package example

import (
	"fmt"

	"github.com/Tympanix/automato/hub"
	"github.com/Tympanix/automato/task"
)

// PersonEvent event to test the application
type PersonEvent struct {
	task.Event
	Output Person
}

// Person event to test the application
type Person struct {
	Name    string
	Age     int
	Heigth  float32
	Married bool
}

func init() {
	fmt.Println("Register person!")
	hub.Register(&PersonEvent{})
}

// Trigger creates a dummy event which output is a data collection of a person
func (e *PersonEvent) Trigger() error {
	e.Output.Name = "John Doe"
	e.Output.Age = 42
	e.Output.Heigth = 182
	e.Output.Married = true
	return nil
}
