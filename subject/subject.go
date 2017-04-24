package subject

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Tympanix/automato/state"
)

const (
	ioTag     = "io"
	inputTag  = "input"
	outputTag = "output"
)

// Subject is a type which can manipulate and analyze structs
type Subject struct {
	Resolver
	subject  interface{}
	Identity string    `json:"id"`
	Name     string    `json:"name"`
	In       []*Input  `json:"input"`
	Out      []*Output `json:"output"`
}

// Resolver is an interface for a type which can resolve the subject given
// its type as a string. The resolver has to know the type of object
// and resolve that object. This will be used with JSON serialization
// because only a string representation of the subject type is known at that point
type Resolver interface {
	ResolveSubject(string) (interface{}, error)
}

// New creates a new subject from a input/output type. The input and output
// of the type is analysed and can be manipulated through the subject
func New(io interface{}, resolver Resolver) *Subject {
	subject := &Subject{
		Resolver: resolver,
		subject:  io,
		Identity: structName(io),
		In:       make([]*Input, 0),
		Out:      make([]*Output, 0),
	}
	subject.analyseSubject()
	return subject
}

// SetResolver sets the resolver used to identify objects
func (s *Subject) SetResolver(resolver Resolver) {
	s.Resolver = resolver
}

// GetSubject returns the underlying subject
func (s *Subject) GetSubject() interface{} {
	return s.subject
}

func (s *Subject) addInput(input *Input) {
	s.In = append(s.In, input)
}

func (s *Subject) addOutput(output *Output) {
	s.Out = append(s.Out, output)
}

func (s *Subject) analyseSubject() {
	structValue := reflect.Indirect(reflect.ValueOf(s.subject))
	structType := structValue.Type()

	if structValue.Kind() != reflect.Struct {
		log.Fatal("Subject must be of type struct")
	}

	for i := 0; i < structType.NumField(); i++ {
		fieldType := structType.Field(i)
		fieldValue := structValue.Field(i)
		fieldTag, ok := fieldType.Tag.Lookup(ioTag)

		if !ok {
			continue
		}

		if !fieldValue.IsValid() || !fieldValue.CanSet() {
			log.Fatalf("Field %s of %s not assignable", fieldType.Name, s.Type())
		}

		if fieldTag == inputTag {
			s.addInput(NewInput(fieldType, fieldValue))
		} else if fieldTag == outputTag {
			s.addOutput(NewOutput(fieldType, fieldValue))
		} else {
			log.Fatalf("Unknown input/put %s for %s", fieldTag, s.Type())
		}
	}

	if len(s.In) == 0 && len(s.Out) == 0 {
		log.Fatalf("No input/output specified for %s", s.Type())
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

	for _, input := range s.In {
		if !input.IsValid() || !input.CanSet() {
			return fmt.Errorf("Could not set field ”%v” for ”%v”", input.Name, s.Name)
		}

		if len(input.Recipe) == 0 {
			return fmt.Errorf("Missing recipe for field ”%s” of ”%s”", input.Name, s.Type())
		}

		ingredient := input.Recipe[0]

		value, err := ingredient.GetValue(state)

		if err != nil {
			return err
		}

		if !value.Type().AssignableTo(input.Type()) {
			return fmt.Errorf("Field ”%s” of value ”%v” can't be assigned ”%v”", input.Name, input.Value.Type(), value.Type())
		}

		input.Value.Set(value)

	}

	return nil
}

// StoreOutput saves the computed output in the state
func (s *Subject) StoreOutput(state state.State) error {
	for _, output := range s.Out {
		if err := state.PutValue(s.Name, output.Name, output.Value); err != nil {
			return err
		}
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
	input.AddIngredient(&Ingredient{
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
	input.AddIngredient(&Ingredient{
		Type:  IngredientStatic,
		Value: value,
	})
	return nil
}

// UnmarshalJSON creates an subject from JSON
func (s *Subject) UnmarshalJSON(data []byte) error {
	log.Println("Unmarshal subject")
	if s.Resolver == nil {
		return errors.New("Subject could not be serialized from json because it has no resolver")
	}

	type subjectFromJSON struct {
		Identity string             `json:"id"`
		Name     string             `json:"name"`
		In       []*json.RawMessage `json:"input"`
		Out      []*json.RawMessage `json:"output"`
	}

	jsonSubject := new(subjectFromJSON)

	if err := json.Unmarshal(data, &jsonSubject); err != nil {
		return err
	}

	subject, err := s.ResolveSubject(jsonSubject.Identity)

	if err != nil {
		return err
	}

	*s = *New(subject, s.Resolver)

	if len(s.In) != len(jsonSubject.In) {
		return fmt.Errorf("Expected %d inputs got %d", len(s.In), len(jsonSubject.In))
	}

	if len(s.Out) != len(jsonSubject.Out) {
		return fmt.Errorf("Expected %d outputs got %d", len(s.Out), len(jsonSubject.Out))
	}

	for i, input := range s.In {
		if err := input.ParseRaw(jsonSubject.In[i]); err != nil {
			return err
		}
	}

	//TODO: Set recipe for inputs

	return nil
}
