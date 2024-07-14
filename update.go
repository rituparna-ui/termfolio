package main

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Window.Width = msg.Width
		m.Window.Height = msg.Height

	case tea.KeyMsg:
		msgStr := msg.String()
		if msgStr == "ctrl+c" {
			return m, tea.Quit
		}
		if msgStr == "enter" && m.Screen == "INIT_VIEW" {
			m.Screen = "HOME_VIEW"
		}
	}
	return m, nil
}
