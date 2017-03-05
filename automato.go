package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tympanix/automato/api"
	"github.com/Tympanix/automato/example"
	"github.com/Tympanix/automato/task"
)

const (
	apiRoot = "/api"
)

func main() {
	// Create a new event and give it a name for reference
	event := task.NewComponent(&example.PersonEvent{})
	event.SetName("Person")

	// Create a new converter, set its name, and give it an ingredient
	converter := task.NewComponent(&example.StringConverter{})
	converter.SetName("Strcon").
		AddStatic("Format", "Person %s would like to say hello").
		AddVar("Placeholder", "Person", "Name")

	// Create a new action, set its name, and give it an ingredient
	action := task.NewComponent(&example.EmailAction{})
	action.SetName("email").
		AddVar("Message", "Strcon", "Formatted").
		AddStatic("Subject", "A new friend").
		AddStatic("Receiver", "johndoe@email.com")

	// Finally create the task consisting of the components above
	task := task.Task{
		Name:    "My First Task",
		Event:   event,
		Actions: []*task.Component{converter, action},
	}

	task.Run()
	fmt.Println("Task completed!")

	// Set up api handler
	http.Handle(apiRoot+"/", http.StripPrefix(apiRoot, api.API))

	// Set up file server for static files
	fs := http.FileServer(http.Dir("web/dist"))
	http.Handle("/", fs)

	// Serve the web server
	log.Fatal(http.ListenAndServe("0.0.0.0:2800", nil))
}
