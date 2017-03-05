package task

// Components contains all available components in the aplication
var Components map[string]*Component

func init() {
	Components = make(map[string]*Component)
}

// Register is called to register a new component in the hub thus
// to make it public for use by the web app
func Register(action Action) {
	component := NewComponent(action)
	Components[component.ID] = component
}

// GetActionByID returns the underlying action for the component identified by id
func GetActionByID(id string) (action Action, ok bool) {
	component, ok := Components[id]
	if !ok {
		return
	}
	return component.Action, ok
}
