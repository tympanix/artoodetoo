package subject_test

import (
	"encoding/json"
	"testing"

	"github.com/Tympanix/artoodetoo/assert"
	"github.com/Tympanix/artoodetoo/state"
	"github.com/Tympanix/artoodetoo/subject"
)

type TestSubject struct {
	Name    string `io:"input"`
	Allowed bool   `io:"output"`
}

func (t *TestSubject) ResolveSubject(typ string) (interface{}, error) {
	return new(TestSubject), nil
}

func TestSubjectJSON(t *testing.T) {

	test := new(TestSubject)
	sub := subject.New(test, nil)

	data, err := json.Marshal(sub)
	assert.NotError(t, err)

	copy := new(subject.Subject)
	copy.SetResolver(new(TestSubject))

	err = json.Unmarshal(data, copy)
	assert.NotError(t, err)
}

func TestNumVariables(t *testing.T) {
	test := new(TestSubject)
	sub := subject.New(test, nil)

	assert.Equal(t, sub.NumVariables(), 0)
	sub.AddVar("Name", "Source", "Variable")
	assert.Equal(t, sub.NumVariables(), 1)
}

func TestSubjectOutput(t *testing.T) {
	test := new(TestSubject)
	sub := subject.New(test, nil)

	test.Allowed = true

	out, err := sub.GetOutputByName("Allowed")
	assert.NotError(t, err)
	assert.Equal(t, out.Bool(), true)

	_, err = sub.GetOutputByName("Does not exist")
	assert.Error(t, err)
}

func TestSubjectNoName(t *testing.T) {
	test := new(TestSubject)
	sub := subject.New(test, nil)
	err := sub.Validate()
	assert.Error(t, err)
}

func TestSubjectNoInput(t *testing.T) {
	test := new(TestSubject)
	sub := subject.New(test, nil)
	sub.SetName("Some Name")
	err := sub.Validate()
	assert.Error(t, err)
}

func TestSubjectStaticInput(t *testing.T) {
	test := new(TestSubject)
	sub := subject.New(test, nil)
	sub.AddStatic("Name", "John Doe")
	s := state.New()
	err := sub.AssignInput(s)
	assert.NotError(t, err)
}

func TestSubjectNilInput(t *testing.T) {
	test := new(TestSubject)
	sub := subject.New(test, nil)
	sub.AddStatic("Name", nil)
	s := state.New()
	err := sub.AssignInput(s)
	assert.Error(t, err)
}

func TestSubjectVarInput(t *testing.T) {
	test := new(TestSubject)
	sub := subject.New(test, nil)

	ts := state.New()
	ts.Put("source:var", "John Doe")

	sub.AddVar("Name", "source", "var")
	err := sub.AssignInput(ts)
	assert.NotError(t, err)
	assert.Equal(t, test.Name, "John Doe")
}

func TestSubjectVarInputError(t *testing.T) {
	test := new(TestSubject)
	sub := subject.New(test, nil)

	ts := state.New()
	ts.Put("source:var", false)

	sub.AddVar("Name", "source", "var")
	err := sub.AssignInput(ts)
	assert.Error(t, err)
}

type ErrorTest struct {
	Error error `io:"output"`
}

func TestSubjectStoreNil(t *testing.T) {

	sub := subject.New(&ErrorTest{}, nil)
	ts := state.New()

	err := sub.StoreOutput(ts)
	assert.Error(t, err)
}

func TestSubjectStoreOutput(t *testing.T) {
	test := new(TestSubject)
	sub := subject.New(test, nil)
	sub.SetName("Test")
	test.Allowed = true

	ts := state.New()
	sub.StoreOutput(ts)
	var b bool
	ts.Get("Test:Allowed", &b)
	assert.True(t, b)
}

type fakeResolver struct {
	subject interface{}
}

func (r *fakeResolver) ResolveSubject(hej string) (interface{}, error) {
	return r.subject, nil
}

func TestSubjectRebuild(t *testing.T) {
	test := new(TestSubject)
	sub := subject.New(test, nil)
	sub.AddStatic("Name", "John Doe")

	err := sub.RebuildSubject(&fakeResolver{test})
	assert.NotError(t, err)
	assert.NotEqual(t, sub.GetSubject(), test)
}

func TestSubjectWrongRebuild(t *testing.T) {
	test := new(TestSubject)
	sub := subject.New(test, nil)
	sub.AddStatic("Name", "John Doe")

	type wrongType struct {
		Input string `io:"input"`
	}

	err := sub.RebuildSubject(&fakeResolver{wrongType{}})
	assert.Error(t, err)
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
