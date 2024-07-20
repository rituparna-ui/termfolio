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
	thankYou = lipgloss.NewStyle().Blink(true).Foreground(lipgloss.Color("63")).Render(thankYou)

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

var contactInfo = []string{
	"Portfolio", "https://rituparnawarwatkar.com/",
	"LinkedIn", "https://www.linkedin.com/in/rituparna-warwatkar/",
	"GitHub", "https://github.com/rituparna-ui/",
	"WhatsApp", "+91 77985 16764",
	"Email", "rwarwatkar@gmail.com",
	"Instagram", "https://www.instagram.com/rituparna_warwatkar/",
}

func GenerateContactView(m *Model) string {
	contactStyle := lipgloss.
		NewStyle().
		Width(m.Window.Width*3/4 - 8).
		Height(m.Window.Height - 5)

	contact := "Contact Me"
	contact = lipgloss.
		NewStyle().
		Bold(true).
		Background(lipgloss.Color("63")).
		PaddingLeft(1).
		PaddingRight(1).
		Render(contact)

	infoList := ""

	for i := 0; i < len(contactInfo); i += 2 {
		info := contactInfo[i]
		link := contactInfo[i+1]

		infoLink := lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.NewStyle().Bold(true).Render(info),
			lipgloss.NewStyle().Faint(true).Render(link),
		)

		infoList += infoLink + "\n\n"
	}

	return contactStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			contact,
			"\n\n",
			infoList,
		),
	)
}

func GetNavItems() []string {
	return []string{"About", "Projects", "Contact Me", "Exit"}
}

func GenerateTabView(m *Model) string {
	switch cursor := m.HomeView.LeftPane.cursor; cursor {
	case 0:
		return "About"
	case 1:
		return "Projects"
	case 2:
		return GenerateContactView(m)
	case 3:
		return GenerateExitView(m)
	}
	return "404 | Not Found"
}
