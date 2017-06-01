package subject

import "fmt"

const (
	// IngredientVar defines an ingredient type which has a input variable
	IngredientVar = iota

	// IngredientStatic defines an ingredient type which has a static value
	// such as string, int, bool ect.
	IngredientStatic = iota
)

// Ingredient describes a variable or static value. If the source is a variable
// it will be a string representation of which unit the ingredient links to.
// The frontend will use ingredients to define input for units is json format
type Ingredient struct {
	Type   int         `json:"type"`
	Source string      `json:"source"`
	Value  interface{} `json:"value"`
}

// IsStatic returns whether or not the ingredient has static content
func (i *Ingredient) IsStatic() bool {
	return i.Type == IngredientStatic
}

// IsVariable returns whether or not the ingredient is a variable representation
func (i *Ingredient) IsVariable() bool {
	return i.Type == IngredientVar
}

func (i *Ingredient) Validate() error {
	// if i.Value == nil {
	// 	return errors.New("Ingredient does not allow nil values")
	// }
	return nil
}

func (i *Ingredient) Key() string {
	if !i.IsVariable() {
		panic("Ingredient is not a variable")
	}
	return fmt.Sprintf("%s:%s", i.Source, i.Value)
}
