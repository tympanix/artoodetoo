package example

import (
	"fmt"

	"github.com/Tympanix/automato/hub"
	"github.com/Tympanix/automato/task"
)

// EmailAction mimcs sending an email as an action
type EmailAction struct {
	task.Base
	Email struct {
		Receiver string
		Subject  string
		Message  string
	}
}

func init() {
	hub.Register(&EmailAction{})
}

// Execute sends the email
func (a *EmailAction) Execute() {
	fmt.Printf("New Mail:\nTo: <%s>\nSubject: %s\nMessage: %s", a.Email.Receiver, a.Email.Subject, a.Email.Message)
}

// Input defines the input type which is accepted by an email event
func (a *EmailAction) Input() interface{} {
	return a.Email
}
