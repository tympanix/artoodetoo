package task

// Action interface
type Action interface {
	Execute()
	Input() interface{}
}
