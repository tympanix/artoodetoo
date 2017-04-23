package types

import "github.com/Tympanix/automato/state"

// Triggerable is a type which returns a trigger and can listen on events
type Triggerable interface {
	Listen() error
	Trigger() chan bool
}

// Observable is a type which you can subscribe and unsunscibe to
type Observable interface {
	Subscribe(Runnable)
	Unsubscribe(Runnable) error
}

// Eventable is a type which can be subsribed to and can trigger events
type Eventable interface {
	Triggerable
	Observable
}

// Runnable is a type which can run certain tasks
type Runnable interface {
	Run(state.State) error
}

// IO is a type which can offer output and accept input
type IO interface {
	Input() interface{}
	Output() interface{}
}
