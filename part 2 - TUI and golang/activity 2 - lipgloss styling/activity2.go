// ACTIVITY 2: Lipgloss Style Playground
//
// In this activity, we’ll explore different ways to style terminal text using Lipgloss.
// The goal is to experiment with colors, bold, underline, italic, and background styles.
//
//  Try changing the styles, combining them, or creating your own, you can also change the word displayed by editing the `word` variable.
//
//  HINTS:
// - Foreground sets text color
// - Background sets background color
// - Bold, Italic, Underline can be chained
// - Use .Render("your text") to apply the style
//
//  Color Chart: https://github.com/charmbracelet/lipgloss#colors

package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	word := "hello, world!!"

	// Basic pink bold style
	pinkBold := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true)

	// Green with underline
	greenUnderline := lipgloss.NewStyle().
		Foreground(lipgloss.Color("34")).
		Underline(true)

	// Cyan italic on black background
	cyanOnBlack := lipgloss.NewStyle().
		Foreground(lipgloss.Color("69")).
		Background(lipgloss.Color("0")).
		Italic(true)

	// Yellow bold with underline
	yellowCombo := lipgloss.NewStyle().
		Foreground(lipgloss.Color("226")).
		Bold(true).
		Underline(true)

	// create your own custom style
	custom := lipgloss.NewStyle().
		Foreground(lipgloss.Color("")). // put a color here
		Background(lipgloss.Color("")). // put a color here
		Bold().                         // true or false
		Italic().                       // true or false
		Underline()                     //true or false

	// Print examples
	fmt.Println(pinkBold.Render("Pink Bold: " + word))
	fmt.Println(greenUnderline.Render("Green Underline: " + word))
	fmt.Println(cyanOnBlack.Render("Cyan Italic on Black: " + word))
	fmt.Println(yellowCombo.Render("Yellow Bold + Underline: " + word))
	fmt.Println(custom.Render("Custom Style: " + word))
}
