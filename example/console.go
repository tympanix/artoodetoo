package example

import (
	"fmt"
	"io"
	"os"
)

// ConsoleAction prints the output the stdout
type ConsoleAction struct{}

// Execute starts the printing to stdout
func (a ConsoleAction) Execute(r *io.PipeReader) {
	io.Copy(os.Stdout, r)
	fmt.Println()
}
