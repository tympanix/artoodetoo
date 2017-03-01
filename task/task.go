package task

import (
	"fmt"

	"github.com/Tympanix/automato/task/component"
)

const (
	variable = iota
	static   = iota
	finish   = iota
)

// Task is an object that processes data based on events, converters and actions
type Task struct {
	Event   component.Component
	Actions []component.Component
}

// Ingredient describes a variable or static value. If the source is a variable
// it will be a string representation of which component the ingredient links to.
// The frontend will use ingredients to define input for components is json format
type Ingredient struct {
	Type  int
	Value interface{}
}

// Run starts the given task by connecting all components by pipes
// such that they cna comminucate with each other as a single workflow
func (t *Task) Run() {
	fmt.Println("Runnings task")

	// m := reflect.ValueOf(&FacebookPhoto{}).Elem()
	// typeOfT := m.Type()
	// for i := 0; i < m.NumField(); i++ {
	// 	fmt.Println(typeOfT.Field(i))
	// }
}
