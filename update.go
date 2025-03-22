package main

import (
	"math/rand/v2"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	case tickMsg:
		if m.BootViewModel.progress.Percent() >= 1 {
			m.BootViewModel.progress.SetPercent(1)
			m.Screen = HomeView
			return m, nil
		}
		randomPerc := rand.Float64()
		if randomPerc > 0.25 {
			randomPerc -= 0.25
		}
		cmd := m.BootViewModel.progress.IncrPercent(randomPerc)
		return m, tea.Batch(tickCmd(), cmd)

	case tea.WindowSizeMsg:
		m.Dim.Width = msg.Width
		m.Dim.Height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.Screen == BootView {
				m.progress.Width = m.Dim.Width / 2
				return m, tickCmd()
			}
		}
	}
	return m, nil
}
