package action

// Action interface
type Action interface {
	Execute()
	Input() interface{}
}
