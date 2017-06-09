package unit_test

// SendEmail mimcs sending an email as an action
type DummyEmail struct {
	Receiver string `io:"input"`
	Subject  string `io:"input"`
	Message  string `io:"input"`
}

// Describe describes what an email action does
func (a *DummyEmail) Describe() string {
	return "Dummy Email"
}

// Execute sends the email
func (a *DummyEmail) Execute() error {
	return nil
}
