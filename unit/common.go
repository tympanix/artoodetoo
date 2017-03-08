package unit

// Metas contains all available units in the aplication
var Metas map[string]*Meta

func init() {
	Metas = make(map[string]*Meta)
}

// Register is called to register a new unit in the hub thus
// to make it public for use by the web app
func Register(action Action) {
	meta := NewMeta(action)
	Metas[meta.ID] = meta
}

// GetActionByID returns the underlying action for the unit identified by id
func GetActionByID(id string) (action Action, ok bool) {
	meta, ok := Metas[id]
	if !ok {
		return
	}
	return meta.Action(), ok
}
