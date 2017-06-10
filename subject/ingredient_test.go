package subject_test

import (
	"testing"

	"github.com/Tympanix/artoodetoo/subject"
	"github.com/stretchr/testify/assert"
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
