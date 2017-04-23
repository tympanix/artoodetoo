package event

// Base is a struct used for subtyping to implement different events
// for the application
type Base struct {
	trigger chan bool
}

// Trigger returns the trigger channel
func (b *Base) Trigger() chan bool {
	return b.trigger
}

// Fire fires of the event
func (b *Base) Fire() {
	b.trigger <- true
}
