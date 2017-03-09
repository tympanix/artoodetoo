package unit

import (
	"encoding/json"
	"errors"
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

// GetOutputByName returns the output field as a reflected value
func (c *Unit) GetOutputByName(name string) (value reflect.Value, err error) {
	value, err = getIOField(name, c.Output())
	if err != nil {
		err = fmt.Errorf("Output with name '%s' could not be resolved", name)
		return
	}
	return
}

// GetInputByName returns the input field as a reflected value
func (c *Unit) GetInputByName(name string) (value reflect.Value, err error) {
	value, err = getIOField(name, c.Input())
	if err != nil {
		err = fmt.Errorf("Input with name '%s' could not be resolved", name)
		return
	}
	return
}

// GetIngredientByArgument returns the ingredient from the recipe which matches
// the argument. If the ingredient is not found an error is returned as 2nd return value
func (c *Unit) GetIngredientByArgument(argument string) (ingr Ingredient, err error) {
	for _, ingredient := range c.Recipe {
		if ingredient.Argument == argument {
			ingr = ingredient
			return
		}
	}
	err = fmt.Errorf("Ingredient for argument '%s' missing", argument)
	return
}

// Validate makes sure that the unit is set up correclt for execution
func (c *Unit) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("Unit was not given a name")
	}
	meta, ok := GetMetaByID(c.ID)
	if !ok {
		return fmt.Errorf("Uknown unit name '%v'", c.ID)
	}
	for _, input := range meta.Input {
		if _, err := c.GetIngredientByArgument(input.Name); err != nil {
			return err
		}
	}
	return nil
}

func getIOField(name string, obj interface{}) (value reflect.Value, err error) {
	if obj == nil {
		err = errors.New("Resolving field form nil object")
		return
	}
	t := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	f := t.FieldByName(name)
	if !f.IsValid() || !f.CanSet() {
		err = errors.New("Field is not valid")
		return
	}
	value = f
	return
}

// AssignInput finds all ingredients in the state given and assigns it as input
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
func (c *Unit) StoreOutput(state state.State) error {
	if err := state.StoreStruct(c.Name, c.Output()); err != nil {
		return err
	}
	return nil
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
