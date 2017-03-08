package unit

// Base is a utility struct that inplements the action interface but does
// not provide any input or output
type Base struct{}

// Output retruns nil as the base action does not by default return any output
func (b *Base) Output() interface{} {
	return nil
}

// Input returns nil as the base action does not by default return ant input
func (b *Base) Input() interface{} {
	return nil
}
