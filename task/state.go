package task

import (
	"fmt"
	"reflect"
)

// State is a mapping of component names and variable names to variable values.
// It is used to store the current state of variables when executing an actions
// by adding new variables to the structure when computed and retrieving variables
// when they are needed for computing a new component
type State map[string]map[string]reflect.Value

// AddOutput takes a component and adds its output to the state
func (s State) AddOutput(c *Component) {
	state, ok := s[c.ID()]

	if !ok {
		state = make(map[string]reflect.Value)
		s[c.ID()] = state
	}

	fmt.Println("Adding output to state")

	output := c.Output()

	if output == nil {
		return
	}

	t := reflect.ValueOf(c.Output())

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	typeOfT := t.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f)
		state[typeOfT.Field(i).Name] = f.Addr()
	}
}
