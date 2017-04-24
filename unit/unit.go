package unit

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
		Subject: *subject.New(a, new(ActionResolver)),
		Desc:    a.Describe(),
		action:  a,
	}
}

// ActionResolver is a subject resolver which can look up actions
type ActionResolver struct{}

// ResolveSubject resolves an action
func (a *ActionResolver) ResolveSubject(t string) (action interface{}, err error) {
	action, ok := GetActionByID(t)
	if !ok {
		err = fmt.Errorf("Could not resolve action with type %v", t)
	}
	log.Printf("Resolved action %v", unitName(action))
	return
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
	type jsonUnit Unit
	c.SetResolver(new(ActionResolver))
	u := jsonUnit(*c)
	if err := json.Unmarshal(b, &u); err != nil {
		return err
	}

	*c = Unit(u)

	return nil
}
