package task

import "io"

// Action interface can operate on input data and perform an action whith that data
type Action interface {
	Execute(*io.PipeReader)
}
