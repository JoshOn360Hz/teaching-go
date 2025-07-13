//  TASK: Style Your Name using Lipgloss
//
// This activity will:
// 1. Ask the user for their name
// 2. Apply a custom style using Lipgloss
// 3. Print their name with the style
//
// Most of the code is done, you just need to define your own style
//
//  HINTS:
// - Use lipgloss.NewStyle()
// - Try combinations like:
//     .Foreground(lipgloss.Color("205"))
//     .Background(lipgloss.Color("0"))
//     .Bold(true)
//     .Italic(true)
// - Use style.Render(name) to style the text
//
// Color Chart: https://github.com/charmbracelet/lipgloss#colors
//

package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	var name string

	fmt.Print("What's your name? ")
	fmt.Scanln(&name) // read user input and store it in the variable 'name'

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("0")).
		Bold(true).
		Italic(true)

	// Print your styled name below here // if you need help check the last activity //

	// print the styled name above here //
	fmt.Println(style.Render(name))

}
