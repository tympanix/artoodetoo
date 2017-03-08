package unit

// Action interface
type Action interface {
	Execute()
	Input() interface{}
	Output() interface{}
}
