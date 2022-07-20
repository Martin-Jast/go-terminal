package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	fps       = 60
	maxHeight = 100
)

type frameMsg time.Time

func animate() tea.Cmd {
	return tea.Tick(time.Second/fps, func(t time.Time) tea.Msg {
		return frameMsg(t)
	})
}

func wait(d time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(d)
		return nil
	}
}

type ShowableThing interface {
	View() string
	Update(msg tea.Msg) (*SubMenu, tea.Cmd)
}
type SubMenu struct {
	OpenedSubMenu *SubMenu
	ParentSubMenu *SubMenu
	Content       ShowableThing
}

func (s SubMenu) View() string {
	if s.OpenedSubMenu != nil {
		return s.OpenedSubMenu.View()
	}
	return s.Content.View()
}

func (s *SubMenu) OnOpenedSubMenuExit() {
	s.OpenedSubMenu = nil
}

func (s *SubMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			if s.ParentSubMenu != nil {
				s.ParentSubMenu.OnOpenedSubMenuExit()
				return s.ParentSubMenu, animate()
			}
			return s, tea.Quit
		case "ctrl+c":
			return s, tea.Quit
		}
	}
	sbm, cmd := s.Content.Update(msg)
	if sbm == nil {
		sbm = s
	} else {
		sbm.ParentSubMenu = s
		s.OpenedSubMenu = sbm
	}
	return sbm, cmd
}

func (s SubMenu) Init() tea.Cmd {
	return tea.Sequentially(wait(time.Second/2), animate())
}
