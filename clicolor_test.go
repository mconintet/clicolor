package clicolor

import (
	"fmt"
	"testing"
)

func TestColorString(t *testing.T) {
	colorized, _ := Colorize("Golang", "white", "blue")
	fmt.Println(colorized)
}
