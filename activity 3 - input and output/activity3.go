// in this activuty we are building on the past 2 , in this we are creating a program that asks for the user's name and then greets them.

package main

// this time we are importing more packages , buffio to read input from the user, fmt to format the output, os to access the operating system, and strings to manipulate strings.

import (
	"fmt"
	"strings"
)

// main funtions
func main() {
	fmt.Println("What's your name?")

	// we are going to set up a buffered reader to read input from the user. This allows us to read a line of text until the user presses Enter.

	// we can do it like this:
	// reader := bufio.NewReader(os.Stdin)
	// this creates a new reader that reads from standard input

	// set up a buffered reader below here: //

	// finish the buffered reader above here //

	// then we can read the input using and assign it to the value name like this:
	// name, _ := reader.ReadString('\n')
	// This reads until a newline character is encountered.

	// assign the input to name below here: //

	// finish the input assignment above here //

	// we can also trim any leading or trailing whitespace from the input using strings.TrimSpace
	name = strings.TrimSpace(name)

	// we are going to print a greeting message using fmt.Println

	fmt.Println("Hello,", name)
}
