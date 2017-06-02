package task

import (
	"errors"

	"github.com/Tympanix/artoodetoo/unit"
)

type dag struct {
	task  *Task
	edges map[*unit.Unit]int
}

func (c *dag) PutAction(action *unit.Unit) {
	c.edges[action] = action.NumVariables()
}

func (c *dag) Next() (*unit.Unit, error) {
	for unit, value := range c.edges {
		if value <= 0 {
			delete(c.edges, unit)
			return unit, nil
		}
	}
	return nil, errors.New("Task has cyclic dependency")
}

func (c *dag) ReduceEdges(name string) error {
	for key, _ := range c.edges {
		for _, in := range key.In {
			for _, r := range in.Recipe {
				if r.IsVariable() {
					if r.Source == name {
						c.edges[key] = c.edges[key] - 1
					}
				}
			}
		}
	}
	return nil
}

func (t *Task) detectCycles() error {
	dag := &dag{
		task:  t,
		edges: make(map[*unit.Unit]int),
	}

	for _, action := range t.Actions {
		dag.PutAction(action)
	}

	dag.ReduceEdges(t.Event.Name)

	for len(dag.edges) > 0 {
		next, err := dag.Next()
		if err != nil {
			return err
		}
		dag.ReduceEdges(next.Name)
	}

	return nil
}
