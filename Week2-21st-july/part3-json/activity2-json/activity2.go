// in this activity, you will create a json file and read it in go, demo.json has a structure you can use, once you have
// created your own json file, change the path below to activity.json to print it out

// to start try running this file with the command go run . and then edit the activity.json and go file

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"` // change these to match your json file e.g name to user_name and`json:"userID"`
	Age  int    `json:"age"`  // change these to match your json file e.g age to user_age
}

func main() {
	f, _ := os.Open("demo.json") // change this to demo.json after you have made your own
	defer f.Close()
	var people []Person
	_ = json.NewDecoder(f).Decode(&people)

	for _, p := range people {
		// change these to match your json file e.g Name to UserName and Age to UserAge like p.name to p.userID
		fmt.Printf("Name: %s | Age: %d\n", p.Name, p.Age)
	}
}
