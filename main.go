package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
)

var gameInstace *game

type game struct {
	cursor  int // which to-do list item our cursor is pointing at
	options []MenuOption
	Pos     harmonica.Point
	Player  *Player
}

func initialModel() game {
	p := InitPlayer()
	inventory := MenuOption{
		Description: "Check Inventory",
		asOption:    checkInventory(p.Inv),
	}
	m := game{
		options: []MenuOption{},
		Player:  p,
	}
	m.options = append(m.options, inventory)
	return m
}

func (m *game) Update(msg tea.Msg) (*SubMenu, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		if m.Pos.Y > maxHeight {
			return nil, tea.Quit
		}

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			return &SubMenu{
				Content: m.options[m.cursor].asOption,
			}, animate()
		default:
			return nil, animate()
		}
	}

	// Return the updated game to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return nil, nil
}
func (m game) View() string {
	// The header
	s := "What should i do?\n\n"

	// Iterate over our choices
	for i, choice := range m.options {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s  %s\n", cursor, choice.Description)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	var out strings.Builder

	for y := 0; y < int(m.Pos.Y); y++ {
		out.WriteString("\n")
	}

	for x := 0; x < int(m.Pos.X); x++ {
		out.WriteString(" ")
	}
	out.WriteString(fmt.Sprintf("(%.2f, %.2f)", m.Pos.X, m.Pos.Y))

	// return out.String()
	return s
}

func main() {
	m := initialModel()
	mainMenu := &SubMenu{
		Content: &m,
	}
	gameInstace = &m
	p := tea.NewProgram(mainMenu)
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
