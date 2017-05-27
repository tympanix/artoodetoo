package mail

import (
	"log"

	"github.com/Tympanix/automato/unit"
)

// SendEmail mimcs sending an email as an action
type SendEmail struct {
	Receiver string `io:"input"`
	Subject  string `io:"input"`
	Message  string `io:"input"`
}

func init() {
	unit.Register(&SendEmail{})
}

// Describe describes what an email action does
func (a *SendEmail) Describe() string {
	return "Sends an E-mail to a chosen receiver with a chosen subject and message"
}

// Execute sends the email
func (a *SendEmail) Execute() {
	log.Printf("New Mail:\nTo: <%s>\nSubject: %s\nMessage: %s\n", a.Receiver, a.Subject, a.Message)
}
