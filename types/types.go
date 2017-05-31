package types

// Triggerable is a type which returns a trigger and can listen on events
type Triggerable interface {
	Listen(<-chan struct{}) error
	Trigger()
}

// Startable is a type which can be started and stopped
type Startable interface {
	Start()
	Stop()
}

// Observable is a type which you can subscribe and unsunscibe to
type Observable interface {
	Subscribe(Runnable) error
	Unsubscribe(Runnable) error
}

// Eventable is a type which can be subsribed to and can trigger events
type Eventable interface {
	Startable
	Observable
}

// Runnable is a type which can run certain tasks
type Runnable interface {
	Run(TupleSpace) error
}

// IO is a type which can offer output and accept input
type IO interface {
	Input() interface{}
	Output() interface{}
}

// TupleSpace is used to store tuples per the LINDA communication language
type TupleSpace interface {
	Close()
	Get(interface{}, ...interface{}) error
	Put(interface{}, ...interface{}) error
	Query(interface{}, ...interface{}) error
}
