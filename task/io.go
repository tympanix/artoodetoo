package task

// Outputter is a unit which is able to supply other units with output data
type Outputter interface {
	Output() interface{}
}

// Inputter is a unit which is able to receive input from other units
type Inputter interface {
	Input() interface{}
}
