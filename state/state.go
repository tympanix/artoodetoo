package state

import (
	"fmt"
	"reflect"
)

// State is a mapping of unit names and variable names to variable values.
// It is used to store the current state of variables when executing a task
// by adding new variables to the structure when computed and retrieving variables
// when they are needed for computing a new unit
type State map[string]map[string]reflect.Value

// New returns a new state
func New() State {
	return make(State)
}

// GetValue returns the value stores for the given domain and key
func (s State) GetValue(domain string, key string) (value reflect.Value, ok bool) {
	value, ok = s[domain][key]
	return
}

// PutValue puts a new value into the state specified by the domain and key
func (s State) PutValue(domain string, key string, value interface{}) error {
	_, ok := s.GetValue(domain, key)

	if ok {
		return fmt.Errorf("State already contains key ”%v” for domain ”%v”", key, domain)
	}

	state, ok := s[domain]

	if !ok {
		state = make(map[string]reflect.Value)
		s[domain] = state
	}

	state[key] = reflect.ValueOf(value)
	return nil
}
