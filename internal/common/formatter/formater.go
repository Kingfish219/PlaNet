package formatter

import "github.com/fatih/color"

func Success() *color.Color {
	return color.New(color.FgGreen)
}
