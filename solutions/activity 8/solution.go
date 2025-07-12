package main

import (
	"fmt"
)

func main() {
	var word string
	var times int

	fmt.Print("Enter a word: ")
	fmt.Scanln(&word)

	fmt.Print("How many times?: ")
	fmt.Scanln(&times)

	for i := 0; i < times; i++ {
		fmt.Println(word)
	}
}
