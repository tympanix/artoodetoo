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

// Unregister removes an action from the application
func Unregister(action Action) {
	meta := NewMeta(action)
	delete(Metas, meta.ID)
}

// GetActionByID returns the underlying action for the unit identified by id
func GetActionByID(id string) (action Action, ok bool) {
	meta, ok := Metas[id]
	if !ok {
		return
	}
	return meta.Action(), ok
}

// GetMetaByID retrieves the meta object associated with the given name
func GetMetaByID(id string) (meta *Meta, ok bool) {
	meta, ok = Metas[id]
	return
}
