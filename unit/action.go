package unit

// Action interface
type Action interface {
	Execute() error
	Describe() string
}
