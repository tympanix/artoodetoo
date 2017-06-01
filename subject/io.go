package subject

import (
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

func (io *IO) Key(name string) string {
	return fmt.Sprintf("%s:%s", name, io.Name)
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

// CopyRecipe assigns the recipe from raw json data
func (i *Input) CopyRecipe(input *Input) error {
	if i.Name != input.Name {
		return fmt.Errorf("Input name %s does not match %s", i.Name, input.Name)
	}

	if i.TypeStr != input.TypeStr {
		return fmt.Errorf("Input type %s does not match %s", input.TypeStr, i.TypeStr)
	}

	i.Recipe = input.Recipe

	return nil
}

// AddIngredient adds an ingredient to the input
func (i *Input) AddIngredient(ingr *Ingredient) {
	i.Recipe = append(i.Recipe, ingr)
}

// Validate makes sure that ingredients are given
func (i *Input) Validate() error {
	if len(i.Recipe) == 0 {
		return fmt.Errorf("Missing ingredient for input '%s'", i.Name)
	}
	for _, ingr := range i.Recipe {
		if err := ingr.Validate(); err != nil {
			return err
		}
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
