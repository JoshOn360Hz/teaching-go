package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
)

type model struct {
	progress progress.Model
	percent  float64
	loading  bool
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func initialModel() model {
	bar := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
	)
	return model{
		progress: bar,
		percent:  0,
		loading:  true,
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
	f, err := os.Open("repos.json")
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
	if len(parts) == 0 {
		return "repo"
	}
	return parts[len(parts)-1]
}
