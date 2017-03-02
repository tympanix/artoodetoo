package task

import (
	"log"
	"reflect"
)

// State is a mapping of component names and variable names to variable values.
// It is used to store the current state of variables when executing an actions
// by adding new variables to the structure when computed and retrieving variables
// when they are needed for computing a new component
type State map[string]map[string]reflect.Value

// AddOutput takes a component and adds its output to the state
func (s State) AddOutput(c *Component) {
	state, ok := s[c.Name()]

	if !ok {
		state = make(map[string]reflect.Value)
		s[c.Name()] = state
	}

	output := c.Output()

	if output == nil {
		return
	}

	t := reflect.ValueOf(output)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	typeOfT := t.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		state[typeOfT.Field(i).Name] = f.Addr()
	}
}

// GetInput reads the recipe from the component and assigns all variables from
// the state as input to the component
func (s State) GetInput(c *Component) {
	input := c.Input()
	if input == nil {
		return
	}
	for _, ingredient := range c.ingredients {
		t := reflect.ValueOf(input)

		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		f := t.FieldByName(ingredient.Argument)

		if !f.IsValid() || !f.CanSet() {
			log.Fatalf("Could not set field %v for component %v\n", f, c)
		}

		value, err := ingredient.GetValue(s)

		if err != nil {
			log.Fatal(err)
		}

		if !value.Type().AssignableTo(f.Type()) {
			log.Fatalf("Field <%s> of value <%v> can't be assigned <%v>\n", ingredient.Argument, f.Type(), value.Type())
		}

		log.Printf("Setting <%v> to \"%v\" as type <%v>\n", ingredient.Argument, value, value.Type())
		f.Set(value)
	}
}
