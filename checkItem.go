package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type checkItemModel struct {
	item InventoryItem
}

func (m checkItemModel) Update(msg tea.Msg) (*SubMenu, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {
		case "enter":
			gameInstace.Player.EquipInventoryItem(&m.item)
			return nil, animate()
		default:
			return nil, animate()
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return nil, nil
}

func (m checkItemModel) View() string {
	s := "Checking: \n\n"
	s += m.item.Name + "\n"
	s += "Type: " + m.item.Type + "\n"
	s += fmt.Sprintf("Power: %d \n", m.item.Power)
	s += fmt.Sprintf("Weight: %d \n\n", m.item.Weight)
	if m.item.IsEquipped {
		s += "Is Equipped. \n\n"
	}
	// The footer
	s += "\nPress q to quit.\n"

	// return out.String()
	return s
}

func checkItem(item InventoryItem) ShowableThing {
	return &checkItemModel{
		item: item,
	}
}
