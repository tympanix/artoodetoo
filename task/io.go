package task

// Outputter is a component which is able to supply other components with output data
type Outputter interface {
	Output() interface{}
}

// Inputter is a component which is able to receive input from other components
type Inputter interface {
	Input() interface{}
}
