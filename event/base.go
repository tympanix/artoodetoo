package event

import "github.com/Tympanix/artoodetoo/types"

// Base is a struct used for subtyping to implement different events
// for the application
type Base struct {
	types.Triggerable `json:"-"`
}

// Bind bind the trigger to the eventable
func (b *Base) Bind(event types.Triggerable) {
	b.Triggerable = event
}
