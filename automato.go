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
	// Create a new event and give it a name for reference
	event := task.NewComponent(example.PersonEvent{})
	event.SetName("Person")

	// Create a new converter, set its name, and give it an ingredient
	converter := task.NewComponent(example.StringConverter{})
	converter.SetName("Strcon").
		AddStatic("Format", "Person %s would like to say hello").
		AddVar("Placeholder", "Person", "Name")

	// Create a new action, set its name, and give it an ingredient
	action := task.NewComponent(example.EmailAction{})
	action.SetName("email").
		AddVar("Message", "Strcon", "Formatted").
		AddStatic("Subject", "A new friend").
		AddStatic("Receiver", "johndoe@email.com")

	// Finally create the task consisting of the components above
	task := task.Task{
		Event:   event,
		Actions: []*task.Component{converter, action},
	}

	task.Run()

	fmt.Println(hub.Events)

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(hub.Components)

	fmt.Println("Task completed!")
}
