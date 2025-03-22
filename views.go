package main

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	switch m.Screen {
	case BootView:
		return GenerateBootView(&m)
	case HomeView:
		return GenerateHomeView(&m)
	default:
		return GenerateNotFoundView(&m)
	}
}

func bootViewCenterTextBox() string {
	line1 := "Hey There!"
	line2 := "This is @rituu's Termfolio"
	line3 := "A Terminal based Portfolio over SSH !!!"

	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Padding(1, 2).
		Foreground(lipgloss.Color("#ea76cb")).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#8839ef")).
		Render(
			lipgloss.JoinVertical(lipgloss.Center,
				line1,
				line2,
				line3,
			),
		)
}

func GenerateBootView(m *Model) string {
	centerTextBox := bootViewCenterTextBox()
	booting := lipgloss.
		NewStyle().
		Blink(true).
		Foreground(lipgloss.Color("#df8e1d")).
		Render("Press Enter to continue ...")

	if m.BootViewModel.progress.Percent() > 0 {
		bootingUp := lipgloss.
			NewStyle().
			Foreground(lipgloss.Color("#8839ef")).
			Render("Booting up...")
		loadingBar := m.BootViewModel.progress.View()

		booting = lipgloss.JoinVertical(lipgloss.Center,
			bootingUp,
			loadingBar,
		)
	}

	return CenterTextStyle(m).Render(
		lipgloss.JoinVertical(lipgloss.Center,
			centerTextBox,
			booting,
		),
	)
}

func GenerateNotFoundView(m *Model) string {
	text := lipgloss.NewStyle().
		Blink(true).
		Foreground(lipgloss.Color("#FF0000")).
		Render("404, Not Found !!!")

	textbox := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF0000")).
		Render(text)

	return CenterTextStyle(m).
		Render(textbox)
}

func homeViewLeftPane(m *Model) string {
	title := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Background(lipgloss.Color("#ea76cb")).
		Foreground(lipgloss.Color("#000000")).
		Padding(0, 1).
		Render("Rituparna Warwatkar")

	leftPane := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Height(m.Dim.Height - 3).
		Width(m.Dim.Width / 4).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#8839ef")).
		Render("Left Pane")

	return lipgloss.JoinVertical(lipgloss.Left,
		title,
		leftPane,
	)
}

func homeViewRightPane(m *Model) string {
	website := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Background(lipgloss.Color("#ea76cb")).
		Foreground(lipgloss.Color("#000000")).
		Padding(0, 1).
		Render("https://rituparnawarwatkar.com")

	rightPane := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Height(m.Dim.Height - 3).
		Width(m.Dim.Width*3/4 - 4).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#8839ef")).
		Render("Right Pane")

	return lipgloss.JoinVertical(lipgloss.Left,
		website,
		rightPane,
	)
}

func GenerateHomeView(m *Model) string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		homeViewLeftPane(m),
		homeViewRightPane(m),
	)
}
