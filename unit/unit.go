package unit

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/Tympanix/automato/subject"
)

const (
	id     = "id"
	output = "output"
	input  = "input"
)

// NewUnit creates a new unit from events, actions and converters
func NewUnit(a Action) *Unit {
	return &Unit{
		Subject: *subject.New(a),
		Desc:    a.Describe(),
		action:  a,
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
	subject.Subject
	Desc   string `json:"description"`
	action Action
}

// Validate makes sure that the unit is set up correctly for execution
func (c *Unit) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("Unit was not given a name")
	}
	if c.Type() != unitName(c.action) {
		return fmt.Errorf("Unit with id '%s' is not valid", c.Type())
	}
	return c.Subject.Validate()
}

// Execute executes the unit by evaluating input and assigning output
func (c *Unit) Execute() {
	c.action.Execute()
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

	*c = Unit(u)

	action, ok := GetActionByID(u.Type())
	if !ok {
		return fmt.Errorf("Can't find unit by id %s", u.Type())
	}

	if err := c.BindIO(action); err != nil {
		return err
	}

	c.action = action

	return nil
}
