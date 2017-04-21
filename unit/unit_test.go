package unit_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Tympanix/automato/assert"
	"github.com/Tympanix/automato/example"
	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/unit"
)

func TestUnitConstructor(t *testing.T) {
	event := &example.PersonEvent{}
	unit := unit.NewUnit(event)

	expectedType := "example.PersonEvent"
	name := "MyFirstTask"

	unit.SetName(name)

	assert.Equal(t, unit.Type(), expectedType)
	assert.Equal(t, unit.Name, name)
	assert.Equal(t, unit.Input(), event.Input())
	assert.Equal(t, unit.Output(), event.Output())
	assert.Equal(t, unit.String(), expectedType)
}

func TestUnitAddIngredient(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)

	for _, in := range u.In {
		assert.Equal(t, len(in.Recipe), 0)
	}

	u.AddStatic("Subject", "Important Email")

	subject, _ := u.GetInputByName("Subject")
	assert.Equal(t, len(subject.Recipe), 1)

}

func TestUnitAddStatic(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)

	u.AddStatic("Message", "My Message")
	message, _ := u.GetInputByName("Message")
	ingredient := message.Recipe[0]

	assert.True(t, ingredient.IsStatic())
	assert.Equal(t, ingredient.Value, "My Message")
	assert.Equal(t, ingredient.Source, "")
	assert.Equal(t, ingredient.Type, unit.IngredientStatic)
}

func TestUnitAddVar(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)

	u.AddVar("Message", "TestEvent", "TestVariable")
	message, _ := u.GetInputByName("Message")

	ingredient := message.Recipe[0]

	assert.True(t, ingredient.IsVariable())
	assert.Equal(t, ingredient.Source, "TestEvent")
	assert.Equal(t, ingredient.Value, "TestVariable")

}

func TestUnitMarshal(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(u)

	var obj *unit.Unit
	dec := json.NewDecoder(&buf)
	err := dec.Decode(&obj)

	assert.NotError(t, err)

	fmt.Printf("%v", u)
	//assert.DeepEqual(t, u, obj)
	assert.DeepEqual(t, u.Action(), obj.Action())
}

func TestUnitMarshalError(t *testing.T) {
	var buf bytes.Buffer

	// An id should be a string
	buf.WriteString("{\"id\":42}")

	var obj *unit.Unit
	dec := json.NewDecoder(&buf)
	err := dec.Decode(&obj)

	assert.Error(t, err)
}

func TestUnitMarshalAction(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(u)

	// Remove email action such that it can't be found
	unit.Unregister(email)

	var obj *unit.Unit
	dec := json.NewDecoder(&buf)
	err := dec.Decode(&obj)

	assert.Error(t, err)
}

func TestUnitExecute(t *testing.T) {
	event := &example.PersonEvent{}
	unit := unit.NewUnit(event)
	unit.SetName("MyUnit")

	state := state.New()
	unit.Execute()
	unit.StoreOutput(state)

	name, _ := state.GetValue(unit.Name, "Name")
	assert.Equal(t, name.Elem().String(), "John Doe")

	age, _ := state.GetValue(unit.Name, "Age")
	assert.Equal(t, age.Elem().Int(), int64(42))

	height, _ := state.GetValue(unit.Name, "Height")
	assert.Equal(t, height.Elem().Float(), float64(182.0))

	married, _ := state.GetValue(unit.Name, "Married")
	assert.Equal(t, married.Elem().Bool(), true)
}

func TestUnitInput(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)
	u.SetName("MyEmail")

	state := state.New()

	source := "MyEvent"
	state.PutValue(source, "Subject", "This is my subject")
	state.PutValue(source, "Receiver", "e@mail.com")

	u.AddStatic("Message", "This is my message")
	u.AddVar("Subject", source, "Subject")
	u.AddVar("Receiver", source, "Receiver")

	err := u.AssignInput(state)
	assert.NotError(t, err)

	assert.Equal(t, email.Email.Message, "This is my message")
	assert.Equal(t, email.Email.Subject, "This is my subject")
	assert.Equal(t, email.Email.Receiver, "e@mail.com")
}

func TestUnitNoInput(t *testing.T) {
	event := &example.PersonEvent{}
	unit := unit.NewUnit(event)

	assert.Equal(t, event.Input(), nil)

	state := state.New()
	noerr := unit.AssignInput(state)

	assert.NotError(t, noerr)

	_, err := unit.GetInputByName("BlahBlah")
	assert.Error(t, err)
}

func TestUnitInputNotValid(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)
	u.SetName("MyEmail")

	u.AddStatic("DoesNotExists", "...")

	state := state.New()
	err := u.AssignInput(state)
	assert.Error(t, err)
}

func TestUnitInputNotInState(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)
	u.SetName("MyEmail")

	u.AddStatic("Message", "My message")
	u.AddStatic("Subject", "My subject")
	u.AddVar("Receiver", "DoesNotExist", "...") // Oops!

	state := state.New()

	err := u.AssignInput(state)
	assert.Error(t, err)
}

func TestUnitInputNotAssignable(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)
	u.SetName("MyEmail")

	u.AddStatic("Message", "My message")
	u.AddStatic("Subject", "My subject")
	u.AddStatic("Receiver", 42) // Oops!

	state := state.New()
	err := u.AssignInput(state)
	assert.Error(t, err)
}

func TestUnitStoreOutputError(t *testing.T) {
	event := &example.PersonEvent{}
	unit := unit.NewUnit(event)
	unit.SetName("MyUnit")

	state := state.New()
	state.PutValue("MyUnit", "Name", "Occupied") // Oops!

	err := unit.StoreOutput(state)
	assert.ErrorContains(t, err, "duplicate")
}

func TestUnitValidateNoName(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)
	err := u.Validate()
	assert.ErrorContains(t, err, "not", "given", "name")
}

func TestUnitValidateNoMeta(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)
	u.SetName("MyEmail")
	u.Identity = "BlahBlah" // Oops!

	err := u.Validate()
	assert.ErrorContains(t, err, "Unit", "not", "valid")
}

func TestUnitValidateMissingIngredient(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)
	unit.Register(email)
	u.SetName("MyEmail")

	u.AddStatic("Subject", "...")
	u.AddStatic("Receiver", "...")

	err := u.Validate()
	assert.ErrorContains(t, err, "Missing", "Message")
}

func TestUnitValidateSucess(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)
	unit.Register(email)
	u.SetName("MyEmail")

	u.AddStatic("Subject", "...")
	u.AddStatic("Receiver", "...")
	u.AddStatic("Message", "...")

	err := u.Validate()
	assert.NotError(t, err)
}
