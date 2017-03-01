package action

// Factory is a factory function returning instantiations of wrappers around actions
func Factory(a Action) *Wrapper {
	return &Wrapper{
		id:     "myaction",
		action: a,
	}
}

// Wrapper wraps an action and extends it's functionality
type Wrapper struct {
	id     string
	action Action
}
