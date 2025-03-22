package main

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func InitModel() Model {
	return Model{
		Screen: BootView,
		BootViewModel: BootViewModel{
			progress: progress.New(progress.WithDefaultGradient()),
		},
	}
}
