package style

import (
	"fmt"

	"github.com/Tympanix/artoodetoo/types"
)

// Make creates a new style from a styleable implementation
func Make(style types.Styleable) Style {
	return Style{
		Color: fmt.Sprintf("#%X", style.Color()),
		Icon:  style.Icon(),
	}
}

// Style is an object which has a color and an icon
type Style struct {
	Color string `json:"color"`
	Icon  string `json:"icon"`
}
