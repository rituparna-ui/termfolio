package main

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
)

type Window struct {
	Width  int
	Height int
}

type InitViewLoaderTickMsg time.Time

type InitView struct {
	showLoader bool
	progress   progress.Model
}

type Screen string

type HomeView struct {
	Tabs       []string
	TabContent []string
	ActiveTab  int
}

type Model struct {
	Window Window
	Screen Screen

	InitView InitView
	HomeView HomeView
}

func InitModel() Model {
	return Model{
		Screen: "INIT_VIEW",
		HomeView: HomeView{
			Tabs:       []string{"Rituparna W", "Projects", "Resume", "Contact Me"},
			TabContent: []string{"Details", "Projects", "Resume", "Contact Me"},
			ActiveTab:  0,
		},
	}
}
