package unit

// Action interface
type Action interface {
	Execute()
	Describe() string
	Input() interface{}
	Output() interface{}
}
