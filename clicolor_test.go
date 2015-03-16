package clicolor

import (
	"fmt"
	"testing"
)

func TestColorString(t *testing.T) {
	colorized, _ := ColorizeStr("Golang", "white", "blue")
	fmt.Println(colorized)
}
