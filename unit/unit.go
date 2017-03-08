package unit

import (
	"encoding/json"
	"fmt"
	"log"
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
		Action: a,
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
	Action Action
}

// Output returns the output from the unit or nil if the unit has no output
func (c *Unit) Output() interface{} {
	return c.Action.Output()
}

// Input return the input from the unit or nil if the unit has no input
func (c *Unit) Input() interface{} {
	return c.Action.Input()
}

// AddIngredient sets the recipe wanted by this unit as input
func (c *Unit) AddIngredient(i Ingredient) *Unit {
	c.Recipe = append(c.Recipe, i)
	return c
}

// AssignInput find all ingredients in the state given and assigns it as input
func (c *Unit) AssignInput(state state.State) *Unit {
	if c.Input() == nil {
		return c
	}

	t := reflect.ValueOf(c.Input())
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for _, ingredient := range c.Recipe {

		f := t.FieldByName(ingredient.Argument)

		if !f.IsValid() || !f.CanSet() {
			log.Fatalf("Could not set field %v for unit %v\n", f, c)
		}

		value, err := ingredient.GetValue(state)

		if err != nil {
			log.Fatal(err)
		}

		if !value.Type().AssignableTo(f.Type()) {
			log.Fatalf("Field <%s> of value <%v> can't be assigned <%v>\n", ingredient.Argument, f.Type(), value.Type())
		}

		log.Printf("Setting <%v> to \"%v\" as type <%v>\n", ingredient.Argument, value, value.Type())
		f.Set(value)
	}
	return c
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

// AddConstraint is a shortcut method for adding an ingredient to the unit
// which is a constrain property for another unit to finish before this one
func (c *Unit) AddConstraint(unitName string) *Unit {
	c.AddIngredient(Ingredient{
		Type:  IngredientFinish,
		Value: unitName,
	})
	return c
}

// Execute executes the unit by evaluating input and assigning output
func (c *Unit) Execute() {
	c.Action.Execute()
}

// SetName sets a new name for the unit
func (c *Unit) SetName(name string) *Unit {
	c.Name = name
	return c
}

func (c *Unit) String() string {
	return c.ID
}

// MarshalJSON returns a json representation of the unit. The json representation
// can be used by frontsends to inspect the units type, it's identification and
// the input/output is can handle.
func (c *Unit) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m[id] = c.ID

	if out := c.Output(); out != nil {
		m[output] = describeIO(out)
	}

	if in := c.Input(); in != nil {
		m[input] = describeIO(in)
	}

	return json.Marshal(&m)
}

type iodescription struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func describeIO(obj interface{}) *[]iodescription {
	var desc []iodescription

	s := reflect.ValueOf(obj)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		iodesc := iodescription{
			Name: typeOfT.Field(i).Name,
			Type: f.Type().String(),
		}
		desc = append(desc, iodesc)
	}

	return &desc
}

// UnmarshalJSON is used to transform json data into a units
func (c *Unit) UnmarshalJSON(b []byte) error {
	type unit Unit
	comp := unit{}
	if err := json.Unmarshal(b, &comp); err != nil {
		return err
	}

	action, ok := GetActionByID(comp.ID)
	if !ok {
		return fmt.Errorf("Can't find action by id %s", comp.ID)
	}

	comp.Action = action
	*c = Unit(comp)
	return nil
}
