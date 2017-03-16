package util

import "github.com/Tympanix/automato/unit"

// AllUnits returns a slice of all available units as meta objects
func AllUnits() []*unit.Unit {
	metas := make([]*unit.Unit, len(unit.Units))
	idx := 0
	for _, v := range unit.Units {
		metas[idx] = v
		idx++
	}
	return metas
}
