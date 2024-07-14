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

	loader := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#008080")).
		Faint(true).
		Blink(true).
		AlignHorizontal(lipgloss.Center).
		Render("Press Enter to continue...")

	if m.InitView.showLoader {
		// loader = lipgloss.NewStyle().
		// 	Foreground(lipgloss.Color("#FF00FF")).
		// 	Bold(true).
		// 	AlignHorizontal(lipgloss.Center).
		// 	Padding(1, 4).
		// 	Render(fmt.Sprintf("Loading... %0.1f", m.InitView.progress.Percent()))
		loader = m.InitView.progress.View()
	}

	outputString := lipgloss.NewStyle().
		Width(m.Window.Width).
		Height(m.Window.Height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(fmt.Sprintf("%s\n\n%s", str, loader))

	return outputString
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

func GenerateHomeView(m Model) string {
	renderedTabs := []string{}
	inactiveTabBorder := tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder := tabBorderWithBottom("┘", " ", "└")

	for i, tab := range m.HomeView.Tabs {
		isActive := i == m.HomeView.ActiveTab

		var tabStyle lipgloss.Style

		if isActive {
			tabStyle = lipgloss.NewStyle().
				Border(activeTabBorder).
				Bold(true).
				Underline(true).
				Foreground(lipgloss.Color("#FF00FF"))
		} else {
			tabStyle = lipgloss.NewStyle().
				Border(inactiveTabBorder)
		}
		paddedText := lipgloss.NewStyle().
			Width(15).
			AlignHorizontal(lipgloss.Center).
			Render(tab)
		renderedTabs = append(renderedTabs, tabStyle.Padding(1, 2, 0, 2).Render(paddedText))
	}
	joined := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	return lipgloss.NewStyle().
		Width(m.Window.Width).
		AlignHorizontal(lipgloss.Center).
		Render(joined)
}
