package task

import (
	"fmt"

	"github.com/Tympanix/automato/task/action"
	"github.com/Tympanix/automato/task/event"
)

// Task is an object that processes data based on events, converters and actions
type Task struct {
	Event      event.Event
	Converters []Converter
	Action     action.Action
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
