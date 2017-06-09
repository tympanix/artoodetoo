package task_test

import (
	"bytes"
	"log"
	"os"

	"github.com/Tympanix/artoodetoo/event"
)

type DummyUnit struct {
	StringA string `io:"input"`
	StringB string `io:"input"`

	Result string `io:"output"`
}

// Describe describes what a Concatenation does
func (s *DummyUnit) Describe() string {
	return "Concatenates string A with string B"
}

// Execute function concatenates strings
func (s *DummyUnit) Execute() error {
	s.Result = s.StringA + s.StringB
	return nil
}

type DummyEvent struct {
	event.Base
	Value string `io:"input"`
}

func (d *DummyEvent) Listen(<-chan struct{}) error {
	return nil
}

func (DummyEvent) Describe() string {
	return "A dummy event"
}

var buf bytes.Buffer

func startCapture() {
	log.SetOutput(&buf)
}

func stopCapture() string {
	log.SetOutput(os.Stderr)
	return buf.String()
}
