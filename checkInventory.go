package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
)

type inventoryScreenModel struct {
	inventory Inventory
	cursor    int // which to-do list item our cursor is pointing at
	selected  int
	Pos       harmonica.Point
}

func initiateInventory(inventory Inventory) ShowableThing {
	return &inventoryScreenModel{
		selected:  -1,
		inventory: inventory,
	}
}

func (m *inventoryScreenModel) Update(msg tea.Msg) (*SubMenu, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.inventory.Items)-1 {
				m.cursor++
			}

		case "enter":
			m.selected = m.cursor

			return &SubMenu{
				Content: checkItem(m.inventory.Items[m.selected]),
			}, animate()

		default:
			return nil, animate()
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return nil, nil
}
func (m inventoryScreenModel) View() string {
	// The header
	s := "Inventory items\n\n"

	// Iterate over our choices
	for i, choice := range m.inventory.Items {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		equipped := ""
		if m.inventory.Items[i].IsEquipped {
			equipped += "-e-"
		}
		// Render the row
		s += fmt.Sprintf("%s %s %s\n", cursor, equipped, choice.Name)
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

func checkInventory(inventory Inventory) ShowableThing {
	m := initiateInventory(inventory)
	return m
}
