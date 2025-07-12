// In this activity, we are expanding our calculator. This time, we’ll ask the user whether they want to
// (M)ultiply or (D)ivide two numbers, then perform the operation based on their choice.

// ive already written the input and prompt code for you, your task is to write the if/else logic.

package main

import "fmt"

func main() {
	var a, b int
	var choice string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&b)

	fmt.Print("Would you like to (M)ultiply or (D)ivide? ")
	fmt.Scanln(&choice)

	// in go an if statement can be written like this:

	// if condition {
	//     // do something
	// } else {

	// in this case we can do like this :

	// if choice == "M" {

	//     // print (logic for multiplication)

	// } else if choice == "D" {

	//     // print and (logic for division)

	// Write an if/else that checks the value of 'choice' and performs the correct operation.

	// write your logic here

}
