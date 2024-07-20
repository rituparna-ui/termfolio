package main

import (
	"math/rand/v2"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type EntryViewLoaderTickMsg time.Time

func EntryViewLoaderTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return EntryViewLoaderTickMsg(t)
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Window.Width = msg.Width
		m.Window.Height = msg.Height

	case progress.FrameMsg:
		progressModel, cmd := m.EntryView.progress.Update(msg)
		m.EntryView.progress = progressModel.(progress.Model)
		return m, cmd

	case EntryViewLoaderTickMsg:
		if m.EntryView.progress.Percent() >= 1.0 {
			m.Screen = "HOME_VIEW"
			m.HomeView.LeftPane.items = GetNavItems()
		}
		cmd := m.EntryView.progress.IncrPercent(rand.Float64() * 1)
		return m, tea.Batch(cmd, EntryViewLoaderTick())

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

		if m.Screen == EntryScreen {
			switch msg.String() {
			case "enter":
				m.EntryView.progress = progress.New(progress.WithDefaultGradient(), progress.WithWidth(m.Window.Width/2))
				m.EntryView.showLoader = true
				return m, EntryViewLoaderTick()
			}
		}

		if m.Screen == HomeScreen {
			switch msg.String() {
			case "tab":
				if m.HomeView.pane == LeftPane {
					m.HomeView.pane = RightPane
				} else {
					m.HomeView.pane = LeftPane
				}
			case "enter":
				if m.HomeView.pane == LeftPane {
					m.HomeView.pane = RightPane
				}
			case "esc":
				if m.HomeView.pane == RightPane {
					m.HomeView.pane = LeftPane
				}
			}
			if m.HomeView.pane == LeftPane {
				switch msg.String() {
				case "up", "k":
					m.HomeView.LeftPane.cursor = (m.HomeView.LeftPane.cursor - 1 + len(m.HomeView.LeftPane.items)) % len(m.HomeView.LeftPane.items)
				case "down", "j":
					m.HomeView.LeftPane.cursor = (m.HomeView.LeftPane.cursor + 1) % len(m.HomeView.LeftPane.items)
				}
			}
			if m.HomeView.pane == RightPane {
				switch msg.String() {
				case "y":
					if m.HomeView.LeftPane.cursor == 3 {
						return m, tea.Quit
					}
				case "n":
					if m.HomeView.LeftPane.cursor == 3 {
						m.HomeView.pane = LeftPane
					}
				}
			}
		}
	}
	return m, nil
}
