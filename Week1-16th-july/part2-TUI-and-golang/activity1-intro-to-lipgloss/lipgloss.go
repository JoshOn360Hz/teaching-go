//  TASK: Add Color and Style to a Word using Lipgloss
//
// Your program should:
// 1. Ask the user for a word
// 2. Use Lipgloss to print that word in a bright color and bold style
//
// Most of the setup is done — you just need to create and apply the style
//
//  HINTS:
// - Use lipgloss.NewStyle()
// - Add a foreground color with .Foreground(lipgloss.Color("205"))
// - Make it bold with .Bold(true)
// - Render the word with .Render(word)
// - will look like this:
// style := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
// fmt.Println(style.Render(word))

package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scanln(&word)

	// Create a style and apply it to the word

	// your code goes here //

	// your code ends here //
}
