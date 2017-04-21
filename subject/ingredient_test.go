package subject_test

import (
	"testing"

	"github.com/Tympanix/automato/assert"
	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/subject"
)

func TestIngredientVariableNameNotString(t *testing.T) {
	ingredient := subject.Ingredient{
		Type:   subject.IngredientVar,
		Source: "mySourceEvent",
		Value:  42, /* Oops */
	}

	state := state.New()
	_, err := ingredient.GetValue(state)
	assert.Error(t, err)
}

func TestIngredientWrongType(t *testing.T) {
	ingredient := subject.Ingredient{
		Type:   42,
		Source: "...",
		Value:  "...",
	}

	state := state.New()
	_, err := ingredient.GetValue(state)
	assert.Error(t, err)
}
