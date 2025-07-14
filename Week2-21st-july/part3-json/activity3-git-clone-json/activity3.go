package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	// 1. Open the repos.json file
	file, err := os.Open("repos.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	defer file.Close()

	// 2. Decode the list of repo URLs
	var urls []string
	if err := json.NewDecoder(file).Decode(&urls); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 3. Loop over each URL and clone it
	for _, url := range urls {
		fmt.Println("Cloning:", url)

		// Create a folder name based on the repo
		parts := len(url)
		name := url[parts-1:]
		if len(name) == 0 || name == "/" {
			name = "repo"
		}

		// Clone into the folder
		_, err := git.PlainClone(name, false, &git.CloneOptions{
			URL:      url,
			Progress: os.Stdout,
		})

		if err != nil {
			fmt.Println("Failed to clone:", url)
			fmt.Println("Reason:", err)
		} else {
			fmt.Println("Done:", name)
		}
	}
}
