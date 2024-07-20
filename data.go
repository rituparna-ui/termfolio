package main

import (
	"github.com/charmbracelet/lipgloss"
)

func GenerateExitView(m *Model) string {
	exitStyle := lipgloss.
		NewStyle().
		Width(m.Window.Width*3/4 - 8).
		Height(m.Window.Height - 5).
		AlignVertical(lipgloss.Center).
		AlignHorizontal(lipgloss.Center)

	thankYou := "Thank you for visiting my Termfolio !!!\n"
	thankYou = lipgloss.NewStyle().Blink(true).Foreground(lipgloss.Color("62")).Render(thankYou)

	openSourceBold := lipgloss.NewStyle().Bold(true).Render("Open Source")
	github := "It's " + openSourceBold + "! check it out at\n"
	github += "https://github.com/rituparna-ui/termfolio"

	githubRender := lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#FF00FF")).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Render(github)

	confirmation := lipgloss.
		NewStyle().
		Width(m.Window.Width*3/4 - 8).
		Foreground(lipgloss.Color("#FF0000")).
		Render("Are you sure you want to exit ? (y/n)")

	return exitStyle.Render(
		lipgloss.JoinVertical(lipgloss.Center, thankYou, githubRender, "\n", confirmation),
	)
}

func GetNavItems() []string {
	return []string{"About", "Projects", "Contact", "Exit"}
}

func GenerateTabView(m *Model) string {
	switch cursor := m.HomeView.LeftPane.cursor; cursor {
	case 0:
		return "About"
	case 1:
		return "Projects"
	case 2:
		return "Contact"
	case 3:
		return GenerateExitView(m)
	}
	return "404 | Not Found"
}
