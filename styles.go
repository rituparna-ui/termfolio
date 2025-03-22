package main

import "github.com/charmbracelet/lipgloss"

func CenterTextStyle(m *Model) lipgloss.Style {
	return lipgloss.
		NewStyle().
		Width(m.Dim.Width).
		Height(m.Dim.Height)
}
