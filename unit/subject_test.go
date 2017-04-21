package unit_test

import (
	"testing"

	"github.com/Tympanix/automato/assert"
	"github.com/Tympanix/automato/unit"
)

type test struct {
	input struct {
		Name    string
		Age     int
		Married bool
		Height  float32
	}
	output struct {
		Success bool
	}
}

func (test *test) Input() interface{} {
	return &test.input
}

func (test *test) Output() interface{} {
	return &test.output
}

func TestSubjectIO(t *testing.T) {

	test := &test{}
	subject := unit.NewSubject(test)
	assert.Equal(t, len(subject.In), 4)
	assert.Equal(t, len(subject.Out), 1)

}
