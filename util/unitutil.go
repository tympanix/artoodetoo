package util

import "github.com/Tympanix/automato/unit"

// AllMetas returns a slice of all available units as meta objects
func AllMetas() []*unit.Meta {
	metas := make([]*unit.Meta, len(unit.Metas))
	idx := 0
	for _, v := range unit.Metas {
		metas[idx] = v
		idx++
	}
	return metas
}
