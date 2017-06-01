package data

import "io"

// Stream is an object which produces readers from which to obtain data
type Stream interface {
	Mimetype() string
	NextReader() io.Reader
}
