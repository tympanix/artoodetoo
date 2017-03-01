package example

import (
	"fmt"

	"github.com/Tympanix/automato/hub"
)

// EmailAction mimcs sending an email as an action
type EmailAction struct {
	Email Email
}

// Email action that mimics sending an email
type Email struct {
	Receiver string
	Subject  string
	Message  string
}

func init() {
	fmt.Println("Register email")
	hub.Register(&EmailAction{})
}

// Execute sends the email
func (a *EmailAction) Execute() {
	fmt.Printf("New Mail:\nTo: <%s>\nSubject: %s\nMessage: %s", a.Email.Receiver, a.Email.Subject, a.Email.Message)
}

func (a *EmailAction) Input() interface{} {
	return a.Email
}
