package task

import (
	"errors"
	"fmt"
)

func (t *Task) reduceEdgeCount(m map[string]int, key string) error {
	for _, action := range t.Actions {
		for _, input := range action.In {
			for _, ingr := range input.Recipe {
				if ingr.IsVariable() && ingr.Source == key {
					n, ok := m[action.Name]
					if !ok {
						return fmt.Errorf("Invalid reference to %s", action.Name)
					}
					m[action.Name] = n - 1
				}
			}
		}
	}
	return nil
}

func (t *Task) detectCycles() error {
	m := make(map[string]int)

	if err := t.Event.Validate(); err != nil {
		return err
	}

	m[t.Event.Name] = 0

	for _, action := range t.Actions {
		m[action.Name] = action.NumVariables()
	}

CYCLE:
	for len(m) > 0 {
		for key, val := range m {
			if val <= 0 {
				delete(m, key)
				if err := t.reduceEdgeCount(m, key); err != nil {
					return err
				}
				continue CYCLE
			}
			return errors.New("Task has cycles")
		}
	}
	return nil
}
