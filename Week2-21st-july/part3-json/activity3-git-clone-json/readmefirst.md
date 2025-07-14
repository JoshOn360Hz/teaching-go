# Activity: Clone Git Repositories from JSON

In this activity, you'll build a Go program that reads a list of GitHub repositories from a JSON file and clones them into your local folder. You'll practice using JSON, working with files, and cloning real Git repos using Go.

---

## Instructions

### Step 1: Create a `repos.json` file

In the same folder as your Go file, create a new file named:

- repos.json

Then add the following content:

```json
[
  "https://github.com/charmbracelet/bubbletea",
  "https://github.com/charmbracelet/lipgloss"
]

This is a JSON array. Each GitHub URL is a string inside double quotes. The list must be wrapped in square brackets [], and items separated by commas.

Step 2: Add Your Own Repository

Find a GitHub project you like and add it to the list in repos.json. For example:

```json
[
  "https://github.com/charmbracelet/bubbletea",
  "https://github.com/your-username/your-repo"
]
```

Once your repos.json file is ready, run the Go script:
```bash
go run .
```

This will:
	•	Read your JSON file
	•	Loop through each GitHub URL
	•	Clone each repo into its own folder

Notes
	•	The folder name is automatically based on the repo URL.
	•	If a folder already exists, it may cause an error when cloning again.
	•	Make sure your URLs are valid and public GitHub repositories.


This is similar to how the finished script we are building will work.