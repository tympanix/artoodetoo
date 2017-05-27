package unit

// Action interface
type Action interface {
	Execute()
	Describe() string
}
