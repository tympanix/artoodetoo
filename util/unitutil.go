package util

import "github.com/Tympanix/automato/unit"

func AllMetas() []*unit.Meta {
	var metas []*unit.Meta
	for _, v := range unit.Metas {
		metas = append(metas, v)
	}
	return metas
}
