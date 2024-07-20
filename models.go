package main

import "github.com/charmbracelet/bubbles/progress"

type Window struct {
	Width  int
	Height int
}

type Screen string

const (
	EntryScreen Screen = "ENTRY_VIEW"
	HomeScreen  Screen = "HOME_VIEW"
)

type EntryView struct {
	showLoader bool
	progress   progress.Model
}

type Model struct {
	Window Window
	Screen Screen

	EntryView EntryView
}

func NewModel() Model {
	return Model{
		Screen: EntryScreen,
	}
}
