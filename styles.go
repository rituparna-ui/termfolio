package main

import "github.com/charmbracelet/lipgloss"

var (
	EntryTextStyle = lipgloss.
			NewStyle().
			Bold(true).
			AlignHorizontal(lipgloss.Center).
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("#FF00FF")).
			BorderForeground(lipgloss.Color("62"))

	EntryBlinkStyle = lipgloss.
			NewStyle().
			Blink(true).
			Foreground(lipgloss.Color("#008080"))
)

var (
	HomeScreenName = lipgloss.
			NewStyle().
			Background(lipgloss.Color("#008080")).
			PaddingLeft(1).
			PaddingRight(1).
			Bold(true)

	HomeScreenHref = lipgloss.
			NewStyle().
			Background(lipgloss.Color("#008080")).
			PaddingLeft(1).
			PaddingRight(1).
			Bold(true).
			Underline(true)
)

func GetHomeScreenLeftContainerStyle(m *Model) lipgloss.Style {
	return lipgloss.
		NewStyle().
		Width(m.Window.Width / 4).
		Height(m.Window.Height - 5).
		AlignVertical(lipgloss.Center).
		AlignHorizontal(lipgloss.Center).
		Border(lipgloss.RoundedBorder())
}

func GetHomeScreenRightContainerStyle(m *Model) lipgloss.Style {
	return lipgloss.
		NewStyle().
		Width(m.Window.Width*3/4 - 4).
		Height(m.Window.Height - 5).
		Border(lipgloss.RoundedBorder()).
		Padding(1).
		MarginRight(4)
}
