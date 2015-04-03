package clicolor

import (
	"fmt"
)

var foregroundColors = map[string]string{
	"black":        "0;30",
	"dark_gray":    "1;30",
	"blue":         "0;34",
	"light_blue":   "1;34",
	"green":        "0;32",
	"light_green":  "1;32",
	"cyan":         "0;36",
	"light_cyan":   "1;36",
	"red":          "0;31",
	"light_red":    "1;31",
	"purple":       "0;35",
	"light_purple": "1;35",
	"brown":        "0;33",
	"yellow":       "1;33",
	"light_gray":   "0;37",
	"white":        "1;37",
}

var backgroundColors = map[string]string{
	"black":      "40",
	"red":        "41",
	"green":      "42",
	"yellow":     "43",
	"blue":       "44",
	"magenta":    "45",
	"cyan":       "46",
	"light_gray": "47",
}

func hasForegroundColor(color string) bool {
	_, ok := foregroundColors[color]
	return ok
}

func hasBackgroundColor(color string) bool {
	_, ok := backgroundColors[color]
	return ok
}

func Colorize(str, foreground, background string) string {
	if !hasForegroundColor(foreground) {
		foreground = "white"
	}

	if !hasBackgroundColor(background) {
		background = "black"
	}

	return fmt.Sprintf("\033[%sm\033[%sm%s\033[0m", foregroundColors[foreground], backgroundColors[background], str)
}
