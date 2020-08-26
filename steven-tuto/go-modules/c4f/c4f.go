package c4f

import (
	"github.com/fatih/color"
)

//Print string with color
func Print(str string) {
	if len(str) == 0 {
		return
	}
	color.Red(str)
}
