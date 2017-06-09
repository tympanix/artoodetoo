package subject_test

import (
	"testing"

	"github.com/Tympanix/artoodetoo/assert"
	"github.com/Tympanix/artoodetoo/subject"
)

func TestIngredientStatic(t *testing.T) {
	ingredient := subject.Ingredient{
		Type:   subject.IngredientStatic,
		Source: "mySourceEvent",
		Value:  42, /* Oops */
	}

	assert.True(t, !ingredient.IsVariable())
	assert.True(t, ingredient.IsStatic())
}
