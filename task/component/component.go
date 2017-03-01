package component

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Tympanix/automato/task"
)

// New creates a new component from events, actions and converters
func New(c interface{}) *Component {
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
	id        string
	component interface{}
}

// ID returns the id of the component
func (c *Component) ID() string {
	return c.id
}

// Output returns the output from the component or nil if the component has no output
func (c *Component) Output() interface{} {
	if out, ok := c.component.(task.Outputter); ok {
		return out.Output()
	}
	return nil
}

// Input return the input from the component or nil if the component has no input
func (c *Component) Input() interface{} {
	if in, ok := c.component.(task.Inputter); ok {
		return in.Input()
	}
	return nil
}

// MarshalJSON returns a json representation of the component. The json representation
// can be used by frontsends to inspect the components type, it's identification and
// the input/output is can handle.
func (c *Component) MarshalJSON() ([]byte, error) {
	fmt.Println("Calling marshal")
	m := make(map[string]interface{})
	m["id"] = c.ID()

	if out := c.Output(); out != nil {
		m["output"] = describeIO(out)
	}

	if in := c.Input(); in != nil {
		m["input"] = describeIO(in)
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
