package printer

import "strings"

/*
This is intended to contain helper functions for printing to the
terminal.
*/

// CenterText takes a text string and a width integer, and centers the text within the specified width by padding with spaces.
func CenterText(text string, width int) string {
	if len(text) >= width {
		return text
	}
	leftPad := (width - len(text)) / 2
	rightPad := width - len(text) - leftPad
	return strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
}
