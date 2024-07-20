package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func GenerateEntryView(m *Model) string {
	centerText := lipgloss.JoinVertical(
		lipgloss.Center,
		"Hey There!",
		"This is Ritu's Termfolio",
		"A Terminal based Portfolio over SSH !!",
	)

	centerText = EntryTextStyle.Render(centerText)
	blinkText := EntryBlinkStyle.Width(m.Window.Width).Render("Press Enter to continue...")

	centerDisplay := lipgloss.
		NewStyle().
		Width(m.Window.Width).
		Height(m.Window.Height).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)

	render := lipgloss.JoinVertical(lipgloss.Center, centerText, blinkText)

	if m.EntryView.showLoader {
		render = lipgloss.JoinVertical(lipgloss.Center, centerText, m.EntryView.progress.View())
	}

	return centerDisplay.Render(render)
}

func GenerateHomeView(m *Model) string {
	name := lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("#FF00FF")).
		Padding(0, 1)

	leftContainer := lipgloss.NewStyle().
		Width(m.Window.Width / 4).
		Height(m.Window.Height - 4).
		Border(lipgloss.RoundedBorder())

	rightContainer := lipgloss.NewStyle().
		Width(m.Window.Width*3/4 - 4).
		Height(m.Window.Height - 4).
		Border(lipgloss.RoundedBorder()).
		MarginRight(4)

	if m.HomeView.pane == LeftPane {
		leftContainer = leftContainer.BorderForeground(lipgloss.Color("#FF00FF"))
		rightContainer = rightContainer.BorderForeground(lipgloss.Color("#121212"))
	} else {
		leftContainer = leftContainer.BorderForeground(lipgloss.Color("#121212"))
		rightContainer = rightContainer.BorderForeground(lipgloss.Color("#FF00FF"))
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		name.Render("Rituparna Warwatkar"),
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			leftContainer.Render("Left Container"),
			rightContainer.Render(fmt.Sprintf("Right Container (%s)", m.HomeView.pane)),
		),
	)
}
