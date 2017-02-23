package task

import "io"

// Converter interface describes an object which can convert one data stream to another
type Converter interface {
	Convert(*io.PipeReader, *io.PipeWriter)
}
