package unit

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Tympanix/automato/state"
)

const (
	id     = "id"
	output = "output"
	input  = "input"
)

// NewUnit creates a new unit from events, actions and converters
func NewUnit(a Action) *Unit {
	return &Unit{
		ID:     unitName(a),
		Recipe: make([]Ingredient, 0),
		action: a,
	}
}

// UnitName returns a string representation of the unit
// specified by the package name, a dot, followed by the name of the struct
func unitName(unit interface{}) string {
	t := reflect.TypeOf(unit)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.String()
}

// Unit wraps the elements of the application and extends it's functionality.
type Unit struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Recipe []Ingredient `json:"recipe"`
	action Action
}

// Output returns the output from the unit or nil if the unit has no output
func (c *Unit) Output() interface{} {
	return c.action.Output()
}

// Input return the input from the unit or nil if the unit has no input
func (c *Unit) Input() interface{} {
	return c.action.Input()
}

// AddIngredient sets the recipe wanted by this unit as input
func (c *Unit) AddIngredient(i Ingredient) *Unit {
	c.Recipe = append(c.Recipe, i)
	return c
}

// AssignInput find all ingredients in the state given and assigns it as input
func (c *Unit) AssignInput(state state.State) error {
	if c.Input() == nil {
		return nil
	}

	t := reflect.ValueOf(c.Input())
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for _, ingredient := range c.Recipe {

		f := t.FieldByName(ingredient.Argument)

		if !f.IsValid() || !f.CanSet() {
			return fmt.Errorf("Could not set field %v for unit %v", f, c)
		}

		value, err := ingredient.GetValue(state)

		if err != nil {
			return err
		}

		if !value.Type().AssignableTo(f.Type()) {
			return fmt.Errorf("Field <%s> of value <%v> can't be assigned <%v>", ingredient.Argument, f.Type(), value.Type())
		}

		f.Set(value)
	}
	return nil
}

// StoreOutput saves the computed output in the state
func (c *Unit) StoreOutput(state state.State) *Unit {
	state.StoreStruct(c.Name, c.Output())
	return c
}

// AddVar is a shortcut method for adding an ingredient to the unit
// which is a variable reference from another unit
func (c *Unit) AddVar(argument string, source string, variable string) *Unit {
	c.AddIngredient(Ingredient{
		Type:     IngredientVar,
		Argument: argument,
		Source:   source,
		Value:    variable,
	})
	return c
}

// AddStatic is a shortcut method for adding an ingredient to the unit
// which is a static argument
func (c *Unit) AddStatic(argument string, value interface{}) *Unit {
	c.AddIngredient(Ingredient{
		Type:     IngredientStatic,
		Argument: argument,
		Value:    value,
	})
	return c
}

// Execute executes the unit by evaluating input and assigning output
func (c *Unit) Execute() {
	c.action.Execute()
}

// SetName sets a new name for the unit
func (c *Unit) SetName(name string) *Unit {
	c.Name = name
	return c
}

func (c *Unit) String() string {
	return c.ID
}

// Action returns the underlying action represented by the unit
func (c *Unit) Action() *Action {
	return &c.action
}

// UnmarshalJSON is used to transform json data into a units
func (c *Unit) UnmarshalJSON(b []byte) error {
	type unit Unit
	u := unit{}
	if err := json.Unmarshal(b, &u); err != nil {
		return err
	}

	action, ok := GetActionByID(u.ID)
	if !ok {
		return fmt.Errorf("Can't find action by id %s", u.ID)
	}

	u.action = action
	*c = Unit(u)
	return nil
}
