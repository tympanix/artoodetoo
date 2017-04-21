package subject

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/types"
)

// Subject is a type which can manipulate and analyze structs
type Subject struct {
	types.IO `json:"-"`
	Identity string    `json:"id"`
	Name     string    `json:"name"`
	In       []*Input  `json:"input"`
	Out      []*Output `json:"output"`
}

// New creates a new subject from a input/output type. The input and output
// of the type is analysed and can be manipulated through the subject
func New(io types.IO) *Subject {
	return &Subject{
		IO:       io,
		Identity: structName(io),
		In:       describeInput(io.Input()),
		Out:      describeOutput(io.Output()),
	}
}

func (s *Subject) String() string {
	return s.Type()
}

// Type returns the underlying struct name
func (s *Subject) Type() string {
	return s.Identity
}

// SetName sets the name of the subject
func (s *Subject) SetName(name string) {
	s.Name = name
}

func structName(unit interface{}) string {
	t := reflect.TypeOf(unit)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.String()
}

func getIOField(name string, obj interface{}) (value reflect.Value, err error) {
	if obj == nil {
		err = errors.New("Resolving field form nil object")
		return
	}
	t := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	f := t.FieldByName(name)
	if !f.IsValid() || !f.CanSet() {
		err = errors.New("Field is not valid")
		return
	}
	value = f
	return
}

// GetOutputByName returns the output field as a reflected value
func (s *Subject) GetOutputByName(name string) (output *Output, err error) {
	for _, output := range s.Out {
		if output.Name == name {
			return output, nil
		}
	}
	err = fmt.Errorf("Output with name '%s' could not be resolved", name)
	return
}

// GetInputByName returns the input field as a reflected value
func (s *Subject) GetInputByName(name string) (input *Input, err error) {
	for _, input := range s.In {
		if input.Name == name {
			return input, nil
		}
	}
	err = fmt.Errorf("Input with name '%s' could not be resolved", name)
	return
}

// Validate makes sure that the unit is set up correctly for execution
func (s *Subject) Validate() error {
	for _, input := range s.In {
		if err := input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// AssignInput finds all ingredients in the state given and assigns it as input
func (s *Subject) AssignInput(state state.State) error {
	if s.Input() == nil {
		return nil
	}

	t := reflect.ValueOf(s.Input())
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for _, input := range s.In {
		f, err := getIOField(input.Name, s.Input())

		if err != nil {
			return fmt.Errorf("Field ”%s” not found in ”%v”", input.Name, s)
		}

		if !f.IsValid() || !f.CanSet() {
			return fmt.Errorf("Could not set field ”%v” for unit ”%v”", input.Name, s.Name)
		}

		if len(input.Recipe) == 0 {
			return fmt.Errorf("Missing recipe for field ”%s” of unit ”%s”", input.Name, s.Name)
		}

		ingredient := input.Recipe[0]

		value, err := ingredient.GetValue(state)

		if err != nil {
			return err
		}

		if !value.Type().AssignableTo(f.Type()) {
			return fmt.Errorf("Field <%s> of value <%v> can't be assigned <%v>", input.Name, f.Type(), value.Type())
		}

		f.Set(value)

	}

	return nil
}

// StoreOutput saves the computed output in the state
func (s *Subject) StoreOutput(state state.State) error {
	if err := state.StoreStruct(s.Name, s.Output()); err != nil {
		return err
	}
	return nil
}

// AddVar is a shortcut method for adding an ingredient to the unit
// which is a variable reference from another unit
func (s *Subject) AddVar(argument string, source string, variable string) error {
	input, err := s.GetInputByName(argument)
	if err != nil {
		return err
	}
	input.AddIngredient(Ingredient{
		Type:   IngredientVar,
		Source: source,
		Value:  variable,
	})
	return nil
}

// AddStatic is a shortcut method for adding an ingredient to the unit
// which is a static argument
func (s *Subject) AddStatic(argument string, value interface{}) error {
	input, err := s.GetInputByName(argument)
	if err != nil {
		return err
	}
	input.AddIngredient(Ingredient{
		Type:  IngredientStatic,
		Value: value,
	})
	return nil
}

func (s *Subject) bindInput() error {
	input := describeInput(s.Input())
	if len(input) != len(s.In) {
		return fmt.Errorf("Unexpected number of inputs for unit %s", s.Type())
	}
	if len(input) == 0 {
		s.In = make([]*Input, 0)
		return nil
	}
	for _, in := range input {
		inn, err := s.GetInputByName(in.Name)
		if err != nil {
			return err
		}
		if !in.Compatible(*inn) {
			return fmt.Errorf("Input %s has incorrect type %s", inn.Name, inn.Type)
		}
		inn.field = in.field
	}
	return nil
}

func (s *Subject) bindOutput() error {
	output := describeOutput(s.Output())
	if len(output) != len(s.Out) {
		return fmt.Errorf("Unexpected number of output for unit %s", s.Type())
	}
	if len(output) == 0 {
		s.Out = make([]*Output, 0)
		return nil
	}
	for _, out := range output {
		outt, err := s.GetOutputByName(out.Name)
		if err != nil {
			return err
		}
		if !out.Compatible(*outt) {
			return fmt.Errorf("Output %s has incorrect type %s", outt.Name, outt.Type)
		}
		outt.field = out.field
	}
	return nil
}

// BindIO rebinds the IO to the subject after json unmarshal has been run
// This is neccessarry to correclt bootstrap a new subject from json
func (s *Subject) BindIO(io types.IO) error {
	if io == nil {
		return errors.New("Cannot bind nil value as a subject")
	}

	s.IO = io

	if err := s.bindInput(); err != nil {
		return err
	}
	if err := s.bindOutput(); err != nil {
		return err
	}
	return nil
}
