package main

import (
	"math/rand/v2"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Window.Width = msg.Width
		m.Window.Height = msg.Height

	case InitViewLoaderTickMsg:
		if m.InitView.progress.Percent() >= 1.0 {
			m.Screen = "HOME_VIEW"
		}
		cmd := m.InitView.progress.IncrPercent(rand.Float64() * 0.25)
		// cmd := m.InitView.progress.IncrPercent(0.5)
		return m, tea.Batch(cmd, InitViewLoaderTick())

	case progress.FrameMsg:
		progressModel, cmd := m.InitView.progress.Update(msg)
		m.InitView.progress = progressModel.(progress.Model)
		return m, cmd

	case tea.KeyMsg:
		msgStr := msg.String()
		if msgStr == "ctrl+c" {
			return m, tea.Quit
		}

		if m.Screen == "INIT_VIEW" && msgStr == "enter" {
			m.InitView.showLoader = true
			m.InitView.progress = progress.New(progress.WithDefaultGradient(), progress.WithWidth(m.Window.Width/2))
			return m, InitViewLoaderTick()
		}

		if m.Screen == "HOME_VIEW" {
			if msgStr == "right" || msgStr == "tab" {
				m.HomeView.ActiveTab = (m.HomeView.ActiveTab + 1) % len(m.HomeView.Tabs)
			}
			if msgStr == "left" || msgStr == "shift+tab" {
				m.HomeView.ActiveTab = (m.HomeView.ActiveTab - 1 + len(m.HomeView.Tabs)) % len(m.HomeView.Tabs)
			}
		}
	}
	return m, nil
}

func InitViewLoaderTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return InitViewLoaderTickMsg(t)
	})
}
