package state

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// Tuple is a slice of reflection values
type Tuple []reflect.Value

// State is a mapping of unit names and variable names to variable values.
// It is used to store the current state of variables when executing a task
// by adding new variables to the structure when computed and retrieving variables
// when they are needed for computing a new unit
type State struct {
	Tuples map[interface{}][]Tuple
	cond   *sync.Cond
}

// New returns a new state
func New() *State {
	return &State{
		Tuples: make(map[interface{}][]Tuple),
		cond:   sync.NewCond(new(sync.Mutex)),
	}
}

func (s *State) String() string {
	var out string
	for key, val := range s.Tuples {
		out = fmt.Sprintf("Key: %v\n", key)
		for _, tuple := range val {
			out = fmt.Sprintf("%s\t - %v\n", out, tuple)
		}
	}
	return out
}

// Get blocks until the tuple with the corrosponding key is available
func (s *State) Get(key interface{}, values ...interface{}) error {
	//log.Printf("Getting value %s %s\n", key, values)
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	template := createTuple(values)

	for {
		ok, err := s.get(key, template)
		if err != nil {
			return err
		}
		if !ok {
			s.cond.Wait()
		} else {
			break
		}
	}
	return nil
}

func (s *State) get(key interface{}, template Tuple) (ok bool, err error) {
	for i, tuple := range s.Tuples[key] {
		if ok, err = check(template, tuple); err != nil {
			return
		} else if ok {
			s.removeTuple(key, i)
			return
		}
	}
	return
}

func (s *State) removeTuple(key interface{}, i int) {
	s.Tuples[key] = append(s.Tuples[key][:i], s.Tuples[key][i+1:]...)
}

func check(template Tuple, tuple Tuple) (ok bool, err error) {
	if ok = len(template) == len(tuple); !ok {
		return
	}
	if ok = match(template, tuple); ok {
		return assign(template, tuple)
	}
	return
}

func match(template Tuple, tuple Tuple) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()

	for i := range template {
		temVal := reflect.Indirect(template[i])
		tupVal := tuple[i]
		if temVal.Kind() == reflect.Func {
			args := []reflect.Value{tupVal}
			if ok = temVal.Call(args)[0].Bool(); !ok {
				return
			}
		} else if !temVal.CanSet() {
			if ok = reflect.DeepEqual(temVal.Interface(), tupVal.Interface()); !ok {
				return
			}
		}
	}
	return true
}

func assign(template Tuple, tuple Tuple) (ok bool, err error) {
	for i := range template {
		temVal := reflect.Indirect(template[i])
		tupVal := tuple[i]
		if temVal.CanSet() {
			if err = assignValue(temVal, tupVal); err != nil {
				return
			}
		}
	}
	return true, nil
}

func assignValue(variable reflect.Value, value reflect.Value) (err error) {
	if !variable.CanSet() || !variable.IsValid() {
		return errors.New("Variable in tuple is not addressable")
	}

	value, err = tryConvert(value, variable.Type())

	if err != nil {
		return err
	}

	if !value.Type().AssignableTo(variable.Type()) {
		return fmt.Errorf("Type %v can't be assigned to %v", value.Type(), variable.Type())
	}

	variable.Set(value)

	return nil
}

// TryConvert tries to convert the value to another type and returns an error instead of panicking
func tryConvert(v reflect.Value, t reflect.Type) (c reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Could not convert type %s to type %s", v.Type(), t)
		}
	}()
	c = v.Convert(t)
	return
}

func createTuple(values []interface{}) Tuple {
	var tuple Tuple
	for _, value := range values {
		if refVal, ok := value.(reflect.Value); ok {
			tuple = append(tuple, refVal)
		} else {
			tuple = append(tuple, reflect.ValueOf(value))
		}
	}
	return tuple
}

// Put writes a new tuple to the tuple space
func (s *State) Put(key interface{}, values ...interface{}) error {
	//log.Printf("Putting value %s %s\n", key, values)
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	tuple := createTuple(values)

	s.Tuples[key] = append(s.Tuples[key], tuple)

	s.cond.Broadcast()

	return nil
}

// Query blocks until a tuple matching the corrosponding key is returned without removing it
func (s *State) Query(key interface{}, values ...interface{}) error {
	//log.Printf("Query value %s\n", key)
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	template := createTuple(values)

	for {
		ok, err := s.query(key, template)
		if err != nil {
			return err
		}
		if !ok {
			s.cond.Wait()
		} else {
			//log.Printf("Got value %s %s\n", key, template[0])
			break
		}
	}
	return nil
}

func (s *State) query(key interface{}, template Tuple) (ok bool, err error) {
	for _, tuple := range s.Tuples[key] {
		if ok, err = check(template, tuple); err != nil {
			return
		} else if ok {
			return
		}
	}
	return
}
