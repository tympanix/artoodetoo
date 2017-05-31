package meta

import (
	"log"

	"github.com/Tympanix/automato/unit"
)

// Debug is used to debug program code
type Debug struct {
	Log interface{} `io:"input"`
}

func init() {
	unit.Register(new(Debug))
}

// Describe debugging
func (d *Debug) Describe() string {
	return "Debugger to print statements to the console"
}

// Execute debugging
func (d *Debug) Execute() {
	log.Println(d.Log)
}
