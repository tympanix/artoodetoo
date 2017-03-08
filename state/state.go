package state

import "reflect"

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
func (s State) PutValue(domain string, key string, value reflect.Value) {
	state, ok := s[domain]

	if !ok {
		state = make(map[string]reflect.Value)
		s[domain] = state
	}

	state[key] = value
}

// StoreStruct takes a struct and stores key values in the strcut
func (s State) StoreStruct(domain string, output interface{}) error {
	state, ok := s[domain]

	if !ok {
		state = make(map[string]reflect.Value)
		s[domain] = state
	}

	if output == nil {
		return nil
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
	return nil
}
