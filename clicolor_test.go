package clicolor

import (
	"fmt"
	"testing"
)

func TestColorString(t *testing.T) {
	colorized := Colorize("Golang", "white", "blue")
	fmt.Println(colorized)
}
