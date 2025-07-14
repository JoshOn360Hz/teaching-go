//  TASK: Add a Progress Bar to a Git Repo Cloner using Bubble Tea + Bubbles
//
// In this exercise, you’ll complete a Go program that reads a list of GitHub repositories
// from a JSON file and clones them. Before cloning, it displays a progress bar to simulate
// a loading state using Bubble Tea and the bubbles/progress component.
//
// Your job is to complete the code in the marked sections.
//
//  What you need to do:
//
// 1. Inside the `model` struct:
//    - Add `progress progress.Model` to hold the loading bar
//    - Add `percent float64` to track completion (from 0.0 to 1.0)
//    - Add `loading bool` to control state
//
// 2. Inside `initialModel()`:
//    - Create a new progress bar using `progress.New(...)`
//    - Use `progress.WithDefaultGradient()` and `progress.WithWidth(40)` to style it
//    - Return a model with progress, percent set to 0, and loading set to true
//
// 3. In `cloneRepos()`:
//    - Complete the path in `os.Open("")` so it loads your `repos.json` file
//
// 4. In `repoNameFromURL()`:
//    - Add an if-statement: if the split result is empty, return `"repo"`
//    - Otherwise return the last part (e.g. `bubbletea` from the URL)
//
// 5. In repos.json:
//    Add a few GitHub repository URLs to clone, like in the last activity
//
//  When finished:
// - Run the program with `go run .`
// - You should see a smooth progress bar
// - Once it completes, your repos will be cloned silently in the background
//
// This is meant to be diffucult so dont worry it it takes a while to complete, if you get really stuck the answer is in the solution.go file

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
)

// firstly we define our model, which will hold the progress bar and the percentage of completion
// you need to define progress as progress.Model, and percent as float64 and loading as bool

type model struct {
	// your code goes below here //

	// your code goes above here //
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func initialModel() model {
	// setup and return a new progress bar model, we need to set up the loading bar to do this type progress.New(
	// then configure it with options like WithDefaultGradient() and WithWidth(40)

	// your code goes below here //

	// your code goes above here //

	// now we return the model with progress : bar , set the percent to 0 and loading to true
	return model{
		// your code goes below here //

		// your code goes above here //
	}
}

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return t
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case time.Time:
		if m.percent >= 1.0 {
			cloneRepos()
			return m, tea.Quit
		}

		m.percent += 0.05
		if m.percent > 1.0 {
			m.percent = 1.0
		}

		return m, tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
			return t
		})
	}

	return m, nil
}

func (m model) View() string {
	if m.percent >= 1.0 {
		return "\n All repositories cloned!\n"
	}
	return fmt.Sprintf("\nCloning repositories...\n%s", m.progress.ViewAs(m.percent))
}

func cloneRepos() {
	f, err := os.Open("") // compleate this with the path to your repos.json file
	if err != nil {
		return
	}
	defer f.Close()

	var urls []string
	if err := json.NewDecoder(f).Decode(&urls); err != nil {
		return
	}

	for _, url := range urls {
		name := repoNameFromURL(url)
		_, _ = git.PlainClone(name, false, &git.CloneOptions{
			URL:      url,
			Progress: io.Discard,
		})
	}
}

func repoNameFromURL(url string) string {
	parts := strings.Split(strings.TrimSuffix(url, ".git"), "/")

	// now we check if parts is empty, if it is we return "repo" as the default name we do this with an if statement
	//  using len(parts) == 0 and then return "repo"

	// your code goes below here //

	// your code goes above here //

	return parts[len(parts)-1]
}
