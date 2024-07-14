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

type Model struct {
	Window Window
	Screen Screen

	InitView InitView
}

func InitModel() Model {
	return Model{
		Screen: "INIT_VIEW",
	}
}
