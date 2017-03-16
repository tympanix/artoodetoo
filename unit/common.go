package unit

// Units contains all available units in the aplication
var Units map[string]*Unit

func init() {
	Units = make(map[string]*Unit)
}

// Register is called to register a new unit in the hub thus
// to make it public for use by the web app
func Register(action Action) {
	unit := NewUnit(action)
	Units[unit.ID] = unit
}

// Unregister removes an action from the application
func Unregister(action Action) {
	unit := NewUnit(action)
	delete(Units, unit.ID)
}

// GetActionByID returns the underlying action for the unit identified by id
func GetActionByID(id string) (action Action, ok bool) {
	unit, ok := Units[id]
	if !ok {
		return
	}
	return *unit.Action(), ok
}

// GetUnitByID retrieves the unit object associated with the given name
func GetUnitByID(id string) (unit *Unit, ok bool) {
	unit, ok = Units[id]
	return
}
