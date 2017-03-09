package unit

import "reflect"

// Meta is a structure which holds metadata about a unit, such that the
// the composition of the unit can be exmplained to the frontend
type Meta struct {
	ID     string `json:"id"`
	Desc   string `json:"description"`
	Output []IO   `json:"output"`
	Input  []IO   `json:"input"`
	action Action
}

// IO describes an input/output for a unit
type IO struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// NewMeta return a new meta object with information abot the action
func NewMeta(a Action) *Meta {
	return &Meta{
		ID:     unitName(a),
		Desc:   a.Describe(),
		Output: describeIO(a.Output()),
		Input:  describeIO(a.Input()),
		action: a,
	}
}

func describeIO(obj interface{}) []IO {
	var desc []IO

	if obj == nil {
		return make([]IO, 0)
	}

	s := reflect.ValueOf(obj)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		iodesc := IO{
			Name: typeOfT.Field(i).Name,
			Type: f.Type().String(),
		}
		desc = append(desc, iodesc)
	}

	return desc
}

// Action returns the actual action the meta object desribes
func (m *Meta) Action() Action {
	return m.action
}
