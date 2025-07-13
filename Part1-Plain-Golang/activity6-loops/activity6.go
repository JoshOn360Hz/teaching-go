// In this activity, we are learning about loops in Go.
// Loops let us run the same code again and again — for example, asking the user something multiple times.

// We will make a simple loop that keeps asking the user to type something,
// and repeats it back until they type "exit".

package main

// We're importing fmt for input/output and strings to clean up the input.
import (
	"fmt"
	"strings"
)

func main() {
	var input string

	// Now we're going to create a loop that runs forever until we tell it to stop.
	// In Go, we use "for { }" to make an infinite loop.

	// for example we can start with for {

	// ==== do the for statement below here ==== //

	// ==== do the for statement above here ==== //

	fmt.Print("Say something (or type 'exit' to quit): ")
	fmt.Scanln(&input)
	input = strings.TrimSpace(input)

	// Inside the loop, we will check if the user typed "exit".

	// we can do this using an if statement. For example if input == "exit" { break }

	// we use a == to check if two things are equal. and {break} to stop the loop.

	//  ==== write the logic to check for "exit" below here ==== //

	// ==== write the logic to check for "exit" above here ==== //

	fmt.Println("You said:", input)

}
