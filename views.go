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

	leftContainer := GetHomeScreenLeftContainerStyle(m)
	rightContainer := GetHomeScreenRightContainerStyle(m)

	if m.HomeView.pane == LeftPane {
		leftContainer = leftContainer.BorderForeground(lipgloss.Color("#FF00FF"))
		rightContainer = rightContainer.BorderForeground(lipgloss.Color("#404040"))
	} else {
		leftContainer = leftContainer.BorderForeground(lipgloss.Color("#404040"))
		rightContainer = rightContainer.BorderForeground(lipgloss.Color("#FF00FF"))
	}

	leftContainerContents := ""
	for i, item := range m.HomeView.LeftPane.items {
		menuStyle := lipgloss.
			NewStyle()

		cursor := " "
		if i == m.HomeView.LeftPane.cursor {
			cursor = ">"
			menuStyle = menuStyle.
				Bold(true).
				Foreground(lipgloss.Color("#FF00FF")).
				Background(lipgloss.Color("#060606")).
				Width(m.Window.Width/4 - 2)
		}
		leftContainerContents += menuStyle.Render(fmt.Sprintf("\n%s %s\n", cursor, item))
	}

	rightContainerContents := GenerateTabView(m)

	row1 := lipgloss.NewStyle().Faint(true).Render("↑/k: Move Up  \t|\tTab/Enter: Switch Panes\t")
	row2 := lipgloss.NewStyle().Faint(true).Render("↓/j: Move down\t|\tEsc      : Back to Left Pane\t")

	helpText := lipgloss.JoinVertical(
		lipgloss.Left,
		row1,
		row2,
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			HomeScreenName.Render("Rituparna Warwatkar"),
			" ",
			HomeScreenHref.Render("https://rituparnawarwatkar.com"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			leftContainer.Render(leftContainerContents),
			rightContainer.Render(rightContainerContents),
		),
		helpText,
	)
}
