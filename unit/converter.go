package unit

// Converter interface describes an object which can convert one data stream to another
type Converter interface {
	Convert()
	Input() interface{}
	Output() interface{}
}
