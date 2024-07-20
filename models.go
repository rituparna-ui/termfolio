package main

import (
	"github.com/charmbracelet/bubbles/progress"
)

type Window struct {
	Width  int
	Height int
}

type Screen string

const (
	EntryScreen Screen = "ENTRY_VIEW"
	HomeScreen  Screen = "HOME_VIEW"
)

type Pane string

const (
	LeftPane  Pane = "left"
	RightPane Pane = "right"
)

type EntryView struct {
	showLoader bool
	progress   progress.Model
}

type LeftPaneModel struct {
	items  []string
	cursor int
}

type HomeView struct {
	pane     Pane
	LeftPane LeftPaneModel
}

type Model struct {
	Window Window
	Screen Screen

	Visitors int

	EntryView EntryView
	HomeView  HomeView
}

func NewModel() Model {
	return Model{
		Screen: EntryScreen,
		HomeView: HomeView{
			pane: LeftPane,
		},
	}
}
