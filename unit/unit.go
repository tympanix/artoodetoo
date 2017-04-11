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
		Desc:   a.Describe(),
		In:     describeInput(a.Input()),
		Out:    describeOutput(a.Output()),
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
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Desc   string    `json:"description"`
	In     []*Input  `json:"input"`
	Out    []*Output `json:"output"`
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

// GetOutputByName returns the output field as a reflected value
func (c *Unit) GetOutputByName(name string) (output *Output, err error) {
	for _, output := range c.Out {
		if output.Name == name {
			return output, nil
		}
	}
	err = fmt.Errorf("Output with name '%s' could not be resolved", name)
	return
}

// GetInputByName returns the input field as a reflected value
func (c *Unit) GetInputByName(name string) (input *Input, err error) {
	for _, input := range c.In {
		if input.Name == name {
			return input, nil
		}
	}
	err = fmt.Errorf("Input with name '%s' could not be resolved", name)
	return
}

// Validate makes sure that the unit is set up correclt for execution
func (c *Unit) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("Unit was not given a name")
	}
	if c.ID != unitName(c.action) {
		return fmt.Errorf("Unit with id '%s' is not valid", c.ID)
	}
	for _, input := range c.In {
		if err := input.Validate(); err != nil {
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

	for _, input := range c.In {
		f, err := getIOField(input.Name, c.Input())

		if err != nil {
			return fmt.Errorf("Field ”%s” not found in ”%v”", input.Name, c)
		}

		if !f.IsValid() || !f.CanSet() {
			return fmt.Errorf("Could not set field ”%v” for unit ”%v”", input.Name, c.Name)
		}

		if len(input.Recipe) == 0 {
			return fmt.Errorf("Missing recipe for field ”%s” of unit ”%s”", input.Name, c.Name)
		}

		ingredient := input.Recipe[0]

		value, err := ingredient.GetValue(state)

		if err != nil {
			return err
		}

		if !value.Type().AssignableTo(f.Type()) {
			return fmt.Errorf("Field <%s> of value <%v> can't be assigned <%v>", input.Name, f.Type(), value.Type())
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
func (c *Unit) AddVar(argument string, source string, variable string) error {
	input, err := c.GetInputByName(argument)
	if err != nil {
		return err
	}
	input.AddIngredient(Ingredient{
		Type:   IngredientVar,
		Source: source,
		Value:  variable,
	})
	return nil
}

// AddStatic is a shortcut method for adding an ingredient to the unit
// which is a static argument
func (c *Unit) AddStatic(argument string, value interface{}) error {
	input, err := c.GetInputByName(argument)
	if err != nil {
		return err
	}
	input.AddIngredient(Ingredient{
		Type:  IngredientStatic,
		Value: value,
	})
	return nil
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

func (c *Unit) bindInput() error {
	input := describeInput(c.Input())
	if len(input) != len(c.In) {
		return fmt.Errorf("Unexpected number of inputs for unit %s", c.ID)
	}
	if len(input) == 0 {
		c.In = make([]*Input, 0)
		return nil
	}
	for _, in := range input {
		inn, err := c.GetInputByName(in.Name)
		if err != nil {
			return err
		}
		if !in.Compatible(*inn) {
			return fmt.Errorf("Input %s has incorrect type %s", inn.Name, inn.Type)
		}
		inn.field = in.field
	}
	return nil
}

func (c *Unit) bindOutput() error {
	output := describeOutput(c.Output())
	if len(output) != len(c.Out) {
		return fmt.Errorf("Unexpected number of output for unit %s", c.ID)
	}
	if len(output) == 0 {
		c.Out = make([]*Output, 0)
		return nil
	}
	for _, out := range output {
		outt, err := c.GetOutputByName(out.Name)
		if err != nil {
			return err
		}
		if !out.Compatible(*outt) {
			return fmt.Errorf("Output %s has incorrect type %s", outt.Name, outt.Type)
		}
		outt.field = out.field
	}
	return nil
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

	if err := c.bindInput(); err != nil {
		return err
	}
	if err := c.bindOutput(); err != nil {
		return err
	}

	return nil
}
