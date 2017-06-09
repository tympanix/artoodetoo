package subject_test

import (
	"encoding/json"
	"testing"

	"github.com/Tympanix/artoodetoo/assert"
	"github.com/Tympanix/artoodetoo/subject"
)

type Test struct {
	Name    string `io:"input"`
	Allowed bool   `io:"output"`
}

func (t *Test) ResolveSubject(typ string) (interface{}, error) {
	return new(Test), nil
}

func TestSubjectJSON(t *testing.T) {

	test := new(Test)
	sub := subject.New(test, nil)

	data, err := json.Marshal(sub)
	assert.NotError(t, err)

	copy := new(subject.Subject)
	copy.SetResolver(new(Test))

	err = json.Unmarshal(data, copy)
	assert.NotError(t, err)
}

//
// type test struct {
// 	input struct {
// 		Name    string
// 		Age     int
// 		Married bool
// 		Height  float32
// 	}
// 	output struct {
// 		Success bool
// 	}
// }
//
// func (test *test) Input() interface{} {
// 	return &test.input
// }
//
// func (test *test) Output() interface{} {
// 	return &test.output
// }
//
// func TestSubjectIO(t *testing.T) {
//
// 	test := &test{}
// 	subject := subject.New(test)
// 	assert.Equal(t, len(subject.In), 4)
// 	assert.Equal(t, len(subject.Out), 1)
//
// }
