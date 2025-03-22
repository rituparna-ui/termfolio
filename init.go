package main

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Init() tea.Cmd {
	return nil
}

func InitModel() Model {
	return Model{
		Screen: BootView,
	}
}
