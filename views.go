package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func GenerateInitView(m Model) string {
	str := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF00FF")).
		Bold(true).
		AlignHorizontal(lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 4).
		Render("Hello, World!\nWelcome to Ritu's Termfolio.\nA Terminal Based Portfolio")

	continueStr := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#008080")).
		Faint(true).
		Blink(true).
		AlignHorizontal(lipgloss.Center).
		Render("Press Enter to continue...")

	outputString := lipgloss.NewStyle().
		Width(m.Window.Width).
		Height(m.Window.Height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(fmt.Sprintf("%s\n\n%s", str, continueStr))

	return outputString
}

func GenerateHomeView(m Model) string {
	return "Home"
}
