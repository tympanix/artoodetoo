package unit

import (
	"fmt"
	"reflect"
)

// Input describes the type of input and the ingredients used
type Input struct {
	Name   string       `json:"name"`
	Type   string       `json:"type"`
	Recipe []Ingredient `json:"recipe"`
	field  reflect.Value
}

// Compatible returns whether or not this input is compatible with another.
// Two inputs are compatible if they have the same name and the same type
func (i Input) Compatible(o Input) bool {
	return i.Name == o.Name && i.Type == o.Type
}

// AddIngredient adds an ingredient to the input
func (i *Input) AddIngredient(ingr Ingredient) {
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
	Name  string `json:"name"`
	Type  string `json:"type"`
	field reflect.Value
}

// Compatible returns whether or not this output is compatible with another.
// Two outputs are compatible if they have the same name and the same type
func (o Output) Compatible(oo Output) bool {
	return o.Name == oo.Name && o.Type == oo.Type
}

func describeInput(obj interface{}) []*Input {
	var input []*Input

	if obj == nil {
		return make([]*Input, 0)
	}

	s := reflect.ValueOf(obj)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		in := Input{
			Name:  typeOfT.Field(i).Name,
			Type:  f.Type().String(),
			field: f,
		}
		input = append(input, &in)
	}

	return input
}

func describeOutput(obj interface{}) []*Output {
	var output []*Output

	if obj == nil {
		return make([]*Output, 0)
	}

	s := reflect.ValueOf(obj)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		out := Output{
			Name:  typeOfT.Field(i).Name,
			Type:  f.Type().String(),
			field: f,
		}
		output = append(output, &out)
	}

	return output
}
