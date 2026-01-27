package printer

import (
	"fmt"
	"strings"
)

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

// PrintToConsole takes a message string and prints it to the console after clearing the terminal screen
// Everything to feed the print related output to this function to ensure the scren is cleared.
func PrintToConsole(message string) {
	// clear terminal
	fmt.Print("\033[H\033[2J")
	// print the game banner
	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║     Gentle Wanderings - A Cozy Map-Making Adventure       ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
	fmt.Println()

	fmt.Println(message)
}
