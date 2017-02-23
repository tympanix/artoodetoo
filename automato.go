package main

import (
	"fmt"

	"github.com/Tympanix/automato/example"
	"github.com/Tympanix/automato/task"
)

func main() {
	task := task.Task{
		Event:      example.DupEvent{Length: 8, String: "A"},
		Converters: []task.Converter{example.RotConverter{Rotate: 13}},
		Action:     example.ConsoleAction{},
	}

	task.Run()

	fmt.Println("Task completed!")
}
