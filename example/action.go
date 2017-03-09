package example

import (
	"log"

	"github.com/Tympanix/automato/unit"
)

// EmailAction mimcs sending an email as an action
type EmailAction struct {
	unit.Base
	Email struct {
		Receiver string
		Subject  string
		Message  string
	}
}

func init() {
	unit.Register(&EmailAction{})
}

// Describe describes what an email action does
func (a *EmailAction) Describe() string {
	return "An example action which mimics sending an email"
}

// Execute sends the email
func (a *EmailAction) Execute() {
	log.Printf("New Mail:\nTo: <%s>\nSubject: %s\nMessage: %s\n", a.Email.Receiver, a.Email.Subject, a.Email.Message)
}

// Input defines the input type which is accepted by an email event
func (a *EmailAction) Input() interface{} {
	return &a.Email
}
