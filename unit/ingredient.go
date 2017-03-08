package unit

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Tympanix/automato/state"
)

const (
	// IngredientVar defines an ingredient type which has a input variable
	IngredientVar = iota

	// IngredientStatic defines an ingredient type which has a static value
	// such as string, int, bool ect.
	IngredientStatic = iota

	// IngredientFinish defines an ingredient type which is the finish event
	// of another unit
	IngredientFinish = iota
)

// Ingredient describes a variable or static value. If the source is a variable
// it will be a string representation of which unit the ingredient links to.
// The frontend will use ingredients to define input for units is json format
type Ingredient struct {
	Type     int
	Argument string
	Source   string
	Value    interface{}
}

// IsStatic returns whether or not the ingredient has static content
func (i *Ingredient) IsStatic() bool {
	return i.Type == IngredientStatic
}

// IsVariable returns whether or not the ingredient is a variable representation
func (i *Ingredient) IsVariable() bool {
	return i.Type == IngredientVar
}

// IsFinish retirns whether or not the ingredient is a finish incident
func (i *Ingredient) IsFinish() bool {
	return i.Type == IngredientFinish
}

// GetValue retrieves the value of the ingredient using the state provided. If the
// value is a static value, the value will be returned without using the state
func (i *Ingredient) GetValue(s state.State) (value reflect.Value, err error) {
	if i.IsStatic() {
		value = reflect.ValueOf(i.Value)
		return
	}
	if i.IsVariable() {
		variable, ok := i.Value.(string)
		if !ok {
			err = errors.New("Variable name is not a string")
			return
		}
		value, ok = s.GetValue(i.Source, variable)
		if !ok {
			err = fmt.Errorf("Value with domain:%s and key:%s was not found in state", i.Source, variable)
			return
		}
		value = reflect.Indirect(value)
		return
	}
	err = fmt.Errorf("Can't resolve ingredient type %d", i.Type)
	return
}
