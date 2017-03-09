package unit_test

import (
	"testing"

	"github.com/Tympanix/automato/assert"
	"github.com/Tympanix/automato/example"
	"github.com/Tympanix/automato/unit"
)

func TestMetaConstructor(t *testing.T) {
	event := &example.PersonEvent{}
	meta := unit.NewMeta(event)

	expectID := "example.PersonEvent"

	assert.Equal(t, meta.ID, expectID)
	assert.Equal(t, meta.Desc, event.Describe())
	assert.Equal(t, len(meta.Output), 4)
	assert.Equal(t, len(meta.Input), 0)
	assert.Equal(t, meta.Action(), event)
}
