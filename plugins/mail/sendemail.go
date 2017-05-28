package mail

import (
	"log"
	"net/smtp"

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
	auth := smtp.PlainAuth(
		"iAutomaton",
		"iautomaton1@gmail.com",
		"iautomaton",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{a.Receiver}
	msg := []byte(
		"Subject: " + a.Subject + "\r\n" +
			"\r\n" +
			a.Message + "\r\n")

	if err := smtp.SendMail("smtp.gmail.com:25", auth, "hmm", to, msg); err != nil {
		log.Fatal(err)
	}

}
