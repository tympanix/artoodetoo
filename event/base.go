package event

import "github.com/Tympanix/automato/types"

// Base is a struct used for subtyping to implement different events
// for the application
type Base struct {
	types.Triggerable
}

// Bind bind the trigger to the eventable
func (b *Base) Bind(event types.Eventable) {
	b.Triggerable = event
}
