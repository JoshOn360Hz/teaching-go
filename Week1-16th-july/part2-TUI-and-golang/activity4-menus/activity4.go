//  TASK: Build a Bubble Tea-powered menu
//
// Your program should:
// 1. Display a list with at least 3 menu items
// 2. Let the user navigate with ↑ ↓ and select with Enter
// 3. Print a message when an item is selected
//
//  HINTS:
// - Fill in the list of items
// - Add your own logic for what happens when each item is selected
// - Press `q` or `ctrl+c` to quit
// - we are using the default menu delegate, so dont worry about styling for now
//

package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// We’re using a simple item type for the menu
type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

type model struct {
	menu list.Model
}

func main() {
	//  Add your menu items here
	items := []list.Item{
		// Example:
		item("Start"),
		//  Add your menu items below here //

		// add your menu items above here //
	}

	const width = 20
	delegate := list.NewDefaultDelegate()
	menu := list.New(items, delegate, width, 10)
	menu.Title = "Menu example"

	m := model{menu: menu}

	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			choice := m.menu.SelectedItem().(item)

			// example:
			// If the user selects "Start", we can print a message and quit

			if choice == "Start" {
				fmt.Println("You chose Start")
				return m, tea.Quit
			}

			// handle your custom choice options below here //

			// handle your custom choice options above here //

		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.menu, cmd = m.menu.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return lipgloss.NewStyle().Padding(1).Render(m.menu.View())
}
