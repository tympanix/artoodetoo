package task

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

const (
	id     = "id"
	output = "output"
	input  = "input"
)

// NewComponent creates a new component from events, actions and converters
func NewComponent(c interface{}) *Component {
	switch t := c.(type) {
	case Event:
	case Converter:
	case Action:
		break
	default:
		log.Fatalf("Cannot create component from type %T", t)
	}

	return &Component{
		id:        componentName(c),
		component: c,
	}
}

// ComponentName returns a string representation of the component
// specified by the package name, a dot, followed by the name of the struct
func componentName(component interface{}) string {
	t := reflect.TypeOf(component)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return fmt.Sprintf("%s", t.String())
}

// Component wraps the elements of the application and extends it's functionality.
type Component struct {
	id          string
	name        string
	ingredients []Ingredient
	component   interface{}
}

// Ingredient describes a variable or static value. If the source is a variable
// it will be a string representation of which component the ingredient links to.
// The frontend will use ingredients to define input for components is json format
type Ingredient struct {
	Type  int
	Value interface{}
}

// ID returns the id of the component
func (c *Component) ID() string {
	return c.id
}

// Output returns the output from the component or nil if the component has no output
func (c *Component) Output() interface{} {
	if out, ok := c.component.(Outputter); ok {
		return out.Output()
	}
	return nil
}

// Input return the input from the component or nil if the component has no input
func (c *Component) Input() interface{} {
	if in, ok := c.component.(Inputter); ok {
		return in.Input()
	}
	return nil
}

// IsEvent returns whether the component is an event or not
func (c *Component) IsEvent() bool {
	_, ok := c.component.(Event)
	return ok
}

// IsAction returns whether the component is an action or not
func (c *Component) IsAction() bool {
	_, ok := c.component.(Action)
	return ok
}

// IsConverter returns whether the component is a converter or not
func (c *Component) IsConverter() bool {
	_, ok := c.component.(Converter)
	return ok
}

// MarshalJSON returns a json representation of the component. The json representation
// can be used by frontsends to inspect the components type, it's identification and
// the input/output is can handle.
func (c *Component) MarshalJSON() ([]byte, error) {
	fmt.Println("Calling marshal")
	m := make(map[string]interface{})
	m[id] = c.ID()

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
