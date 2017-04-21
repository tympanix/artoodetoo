package types

// Eventable is a type which can trigger runnables
type Eventable interface {
	Listen() error
	Subscribe(Runnable)
	Unsubscribe(Runnable)
}

// Runnable is a type which can run certain tasks
type Runnable interface {
	Run() error
}

// IO is a type which can offer output and accept input
type IO interface {
	Input() interface{}
	Output() interface{}
}
