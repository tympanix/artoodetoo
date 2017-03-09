package unit_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/Tympanix/automato/assert"
	"github.com/Tympanix/automato/example"
	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/unit"
)

func TestUnitConstructor(t *testing.T) {
	event := &example.PersonEvent{}
	unit := unit.NewUnit(event)

	expectedID := "example.PersonEvent"
	name := "MyFirstTask"

	assert.Equal(t, unit.ID, expectedID)
	assert.Equal(t, unit.SetName(name).Name, name)
	assert.Equal(t, unit.Input(), event.Input())
	assert.Equal(t, unit.Output(), event.Output())
	assert.Equal(t, len(unit.Recipe), 0)
	assert.Equal(t, unit.String(), expectedID)
}

func TestUnitAddIngredient(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)

	assert.Equal(t, len(u.Recipe), 0)

	ingredient := unit.Ingredient{
		Type:     unit.IngredientStatic,
		Argument: "Subject",
		Value:    "Important Email",
	}

	u.AddIngredient(ingredient)

	assert.Equal(t, len(u.Recipe), 1)
	assert.Equal(t, u.Recipe[0], ingredient)

}

func TestUnitAddStatic(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)

	u.AddStatic("Message", "My Message")
	ingredient := u.Recipe[0]

	assert.True(t, ingredient.IsStatic())
	assert.Equal(t, ingredient.Argument, "Message")
	assert.Equal(t, ingredient.Value, "My Message")
	assert.Equal(t, ingredient.Source, "")
	assert.Equal(t, ingredient.Type, unit.IngredientStatic)
}

func TestUnitAddVar(t *testing.T) {
	email := &example.EmailAction{}
	u := unit.NewUnit(email)

	u.AddVar("Message", "TestEvent", "TestVariable")
	ingredient := u.Recipe[0]

	assert.True(t, ingredient.IsVariable())
	assert.Equal(t, ingredient.Argument, "Message")
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

	assert.DeepEqual(t, u, obj)
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
	err := unit.AssignInput(state)

	assert.NotError(t, err)
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
