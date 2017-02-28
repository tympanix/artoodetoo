package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Tympanix/automato/example"
	"github.com/Tympanix/automato/hub"
	"github.com/Tympanix/automato/task"
)

func main() {
	task := task.Task{
		Event:      &example.PersonEvent{},
		Converters: []task.Converter{},
		Action:     &example.EmailAction{},
	}

	task.Run()

	fmt.Println(hub.Events)

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(hub.Events)

	fmt.Println("Task completed!")
}
