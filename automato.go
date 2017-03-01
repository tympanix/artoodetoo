package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Tympanix/automato/example"
	"github.com/Tympanix/automato/hub"
	"github.com/Tympanix/automato/task"
	"github.com/Tympanix/automato/task/converter"
)

func main() {
	task := task.Task{
		Event:      &example.PersonEvent{},
		Converters: []converter.Converter{},
		Action:     &example.EmailAction{},
	}

	task.Run()

	fmt.Println(hub.Events)

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(hub.Components)

	fmt.Println("Task completed!")
}
