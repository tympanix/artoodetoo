package subject

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// IO is a type which can describe input/output fields
type IO struct {
	reflect.Value `json:"-"`
	Name          string `json:"name"`
	TypeStr       string `json:"type"`
}

// NewIO return a new IO type
func NewIO(field reflect.StructField, value reflect.Value) IO {
	return IO{
		Value:   value,
		Name:    field.Name,
		TypeStr: field.Type.String(),
	}
}

// Compatible returns whether or not this output is compatible with another.
// Two outputs are compatible if they have the same name and the same type
func (io *IO) Compatible(other Output) bool {
	return io.Name == other.Name && io.TypeStr == other.TypeStr
}

// Input describes the type of input and the ingredients used
type Input struct {
	IO
	Recipe []*Ingredient `json:"recipe"`
}

// NewInput returns a new input object from a struct field and value
func NewInput(field reflect.StructField, value reflect.Value) *Input {
	return &Input{
		IO:     NewIO(field, value),
		Recipe: make([]*Ingredient, 0),
	}
}

// ParseRaw assigns the recipe from raw json data
func (i *Input) ParseRaw(raw *json.RawMessage) error {
	type inputFromJSON struct {
		Name    string        `json:"name"`
		TypeStr string        `json:"type"`
		Recipe  []*Ingredient `json:"recipe"`
	}

	jsonInput := new(inputFromJSON)

	if err := json.Unmarshal(*raw, &jsonInput); err != nil {
		return err
	}

	if i.Name != jsonInput.Name {
		return fmt.Errorf("Input name %s does not match %s", i.Name, jsonInput.Name)
	}

	i.Recipe = jsonInput.Recipe

	return nil
}

// AddIngredient adds an ingredient to the input
func (i *Input) AddIngredient(ingr *Ingredient) {
	i.Recipe = append(i.Recipe, ingr)
}

// Validate makes sure that ingredients are given
func (i Input) Validate() error {
	if len(i.Recipe) == 0 {
		return fmt.Errorf("Missing ingredient for input '%s'", i.Name)
	}
	return nil
}

// Output describes the name and type of an output
type Output struct {
	IO
}

// NewOutput returns a new output object from a struct field and value
func NewOutput(field reflect.StructField, value reflect.Value) *Output {
	return &Output{
		NewIO(field, value),
	}
}
