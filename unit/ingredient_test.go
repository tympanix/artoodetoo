package unit_test

import (
	"testing"

	"github.com/Tympanix/automato/assert"
	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/unit"
)

func TestIngredientVariableNameNotString(t *testing.T) {
	ingredient := unit.Ingredient{
		Type:     unit.IngredientVar,
		Argument: "myArgument",
		Source:   "mySourceEvent",
		Value:    42, /* Oops */
	}

	state := state.New()
	_, err := ingredient.GetValue(state)
	assert.Error(t, err)
}

func TestIngredientWrongType(t *testing.T) {
	ingredient := unit.Ingredient{
		Type:     42,
		Argument: "...",
		Source:   "...",
		Value:    "...",
	}

	state := state.New()
	_, err := ingredient.GetValue(state)
	assert.Error(t, err)
}
