package main

import (
	"fmt"
	"strings"
)

func main() {
	var choice string
	var a, b int

	for {
		fmt.Print("Choose operation: (A)dd, (M)ultiply, (D)ivide or type 'exit' to quit: ")
		fmt.Scanln(&choice)

		choice = strings.TrimSpace(choice)

		if choice == "exit" {
			break
		}

		fmt.Print("Enter the first number: ")
		fmt.Scanln(&a)

		fmt.Print("Enter the second number: ")
		fmt.Scanln(&b)

		if choice == "A" {
			fmt.Println("Result:", a+b)
		} else if choice == "M" {
			fmt.Println("Result:", a*b)
		} else if choice == "D" {
			fmt.Println("Result:", a/b)
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}
