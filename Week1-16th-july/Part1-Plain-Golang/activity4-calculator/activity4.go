// In this activity we are creating a simple calculator that asks the user for two numbers, adds them together, and then prints the result.

package main

// We are importing the fmt package to handle formatted input and output.

import "fmt"

// main function
func main() {
	// First, we are going to declare two integer variables to store the numbers the user inputs.
	var a, b int

	// We then ask the user for the first number using fmt.Print
	fmt.Print("Enter the first number: ")

	// We read the user's input using fmt.Scanln, which waits for them to type a number and press Enter.
	// The input gets stored in the variable 'a'. when we do it it should look like this:
	// fmt.Scanln(&a)

	// add the first Scanln line below here: //

	// finish first input above here //

	// Now we ask for the second number.
	fmt.Print("Enter the second number: ")

	// We read the second number the same way and store it in 'b'.
	// add the second Scanln line below here: //

	// finish second input above here //

	// Finally, we add the two numbers together and print the result.
	// We are using fmt.Println again to display the result.

	fmt.Println("The sum is:", a+b)
}
