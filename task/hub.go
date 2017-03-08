package task

// Units contains all available components in the aplication
var Units map[string]*Unit

func init() {
	Units = make(map[string]*Unit)
}

// Register is called to register a new component in the hub thus
// to make it public for use by the web app
func Register(action Action) {
	component := NewUnit(action)
	Units[component.ID] = component
}

// GetActionByID returns the underlying action for the component identified by id
func GetActionByID(id string) (action Action, ok bool) {
	component, ok := Units[id]
	if !ok {
		return
	}
	return component.Action, ok
}
