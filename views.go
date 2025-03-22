package main

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	switch m.Screen {
	case BootView:
		return GenerateBootView()
	default:
		return GenerateNotFoundView(&m)
	}
}

func GenerateBootView() string {
	return "Boot View"
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
