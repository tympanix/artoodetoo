package task

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/unit"
)

// Task is an object that processes data based on events, converters and actions
type Task struct {
	Name    string       `json:"name"`
	Event   *unit.Unit   `json:"event"`
	Actions []*unit.Unit `json:"actions"`
}

// Describe prints our information about the action to the console
func (t *Task) Describe() {
	log.Printf("Task: %v\n", t.Name)
	log.Printf("Event: %v\n", t.Event)
	log.Printf("Actions:\n")

	for _, a := range t.Actions {
		log.Printf(" %s %v\n", "-", a)
	}
}

// Validate checks that all recipes for every task is fulfilled
func (t *Task) Validate() error {
	for _, action := range t.Actions {
		if err := action.Validate(); err != nil {
			return err
		}
		for _, ingr := range action.Recipe {
			argument, err := action.GetInputByName(ingr.Argument)
			if err != nil {
				return fmt.Errorf("Ingredient for unknown attribute '%s'", ingr.Argument)
			}
			value, err := t.getIngredientValue(ingr)
			if err != nil {
				return err
			}
			if !value.Type().AssignableTo(argument.Type()) {
				return fmt.Errorf("Ingredient '%s' of type '%s' incompatible with type '%s'",
					ingr.Argument, argument.Type(), value.Type())
			}
		}
	}
	return nil
}

func (t *Task) getIngredientValue(ingr unit.Ingredient) (value reflect.Value, err error) {
	if ingr.IsVariable() {
		var source *unit.Unit
		source, err = t.GetUnitByName(ingr.Source)
		if err != nil {
			err = fmt.Errorf("Ingredient has unknown source '%s'", ingr.Source)
			return
		}
		target, ok := ingr.Value.(string)
		if !ok {
			err = fmt.Errorf("Ingredient invalid value for '%s' must be string", ingr.Argument)
			return
		}
		value, err = source.GetOutputByName(target)
		if err != nil {
			err = fmt.Errorf("Ingredient referencing unknown variable '%s' from '%s'", ingr.Value, ingr.Source)
			return
		}
		return
	}
	if ingr.IsStatic() {
		return reflect.ValueOf(ingr.Value), nil
	}
	err = errors.New("Unknown ingredient")
	return
}

// GetUnitByName retrieves a unit in the actions list and returns it
func (t *Task) GetUnitByName(name string) (unit *unit.Unit, err error) {
	if t.Event.Name == name {
		return t.Event, nil
	}
	for _, u := range t.Actions {
		if u.Name == name {
			return u, nil
		}
	}
	err = fmt.Errorf("Task does not have unit with name '%s'", name)
	return
}

// Run starts the task
func (t *Task) Run() error {
	state := state.New()

	t.Event.Execute()
	if err := t.Event.StoreOutput(state); err != nil {
		return err
	}

	for _, action := range t.Actions {
		if err := action.AssignInput(state); err != nil {
			return err
		}
		action.Execute()
		if err := action.StoreOutput(state); err != nil {
			return err
		}
	}
	return nil
}
