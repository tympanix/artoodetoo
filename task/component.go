package task

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	id     = "id"
	output = "output"
	input  = "input"
)

// NewComponent creates a new component from events, actions and converters
func NewComponent(a Action) *Component {
	return &Component{
		ID:     componentName(a),
		Action: a,
	}
}

// ComponentName returns a string representation of the component
// specified by the package name, a dot, followed by the name of the struct
func componentName(component interface{}) string {
	t := reflect.TypeOf(component)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.String()
}

// Component wraps the elements of the application and extends it's functionality.
type Component struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Recipe []Ingredient `json:"recipe"`
	Action Action
}

// Output returns the output from the component or nil if the component has no output
func (c *Component) Output() interface{} {
	return c.Action.Output()
}

// Input return the input from the component or nil if the component has no input
func (c *Component) Input() interface{} {
	return c.Action.Input()
}

// AddIngredient sets the recipe wanted by this component as input
func (c *Component) AddIngredient(i Ingredient) *Component {
	c.Recipe = append(c.Recipe, i)
	return c
}

// AddVar is a shortcut method for adding an ingredient to the component
// which is a variable reference from another component
func (c *Component) AddVar(argument string, source string, variable string) *Component {
	c.AddIngredient(Ingredient{
		Type:     IngredientVar,
		Argument: argument,
		Source:   source,
		Value:    variable,
	})
	return c
}

// AddStatic is a shortcut method for adding an ingredient to the component
// which is a static argument
func (c *Component) AddStatic(argument string, value interface{}) *Component {
	c.AddIngredient(Ingredient{
		Type:     IngredientStatic,
		Argument: argument,
		Value:    value,
	})
	return c
}

// AddConstraint is a shortcut method for adding an ingredient to the component
// which is a constrain property for another component to finish before this one
func (c *Component) AddConstraint(componentName string) *Component {
	c.AddIngredient(Ingredient{
		Type:  IngredientFinish,
		Value: componentName,
	})
	return c
}

// Execute executes the component by evaluating input and assigning output
func (c *Component) Execute() {
	c.Action.Execute()
}

// SetName sets a new name for the component
func (c *Component) SetName(name string) *Component {
	c.Name = name
	return c
}

func (c *Component) String() string {
	return c.ID
}

// MarshalJSON returns a json representation of the component. The json representation
// can be used by frontsends to inspect the components type, it's identification and
// the input/output is can handle.
func (c *Component) MarshalJSON() ([]byte, error) {
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

func describeIO(obj interface{}) *map[string]string {
	desc := make(map[string]string)

	s := reflect.ValueOf(obj)

	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		desc[typeOfT.Field(i).Name] = f.Type().String()
	}

	return &desc
}

// UnmarshalJSON is used to transform json data into a components
func (c *Component) UnmarshalJSON(b []byte) error {
	type component Component
	comp := component{}
	if err := json.Unmarshal(b, &comp); err != nil {
		return err
	}

	action, ok := GetActionByID(comp.ID)
	if !ok {
		return fmt.Errorf("Can't find action by id %s", comp.ID)
	}

	comp.Action = action
	*c = Component(comp)
	return nil
}
