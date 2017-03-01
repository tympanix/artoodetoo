package task

const (
	// IngredientVar defines an ingredient type which has a input variable
	IngredientVar = iota

	// IngredientStatic defines an ingredient type which has a static value
	// such as string, int, bool ect.
	IngredientStatic = iota

	// IngredientFinish defines an ingredient type which is the finish event
	// of another component
	IngredientFinish = iota
)

// Ingredient describes a variable or static value. If the source is a variable
// it will be a string representation of which component the ingredient links to.
// The frontend will use ingredients to define input for components is json format
type Ingredient struct {
	Type     int
	Argument string
	Source   string
	Value    interface{}
}

// IsStatic returns whether or not the ingredient has static content
func (i *Ingredient) IsStatic() bool {
	return i.Type == IngredientVar
}

// IsVariable returns whether or not the ingredient is a variable representation
func (i *Ingredient) IsVariable() bool {
	return i.Type == IngredientVar
}

// IsFinish retirns whether or not the ingredient is a finish incident
func (i *Ingredient) IsFinish() bool {
	return i.Type == IngredientFinish
}
