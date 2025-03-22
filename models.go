package main

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
)

type Screen int
type SelectedPane int
type TabIndex int
type tickMsg time.Time

type Dim struct {
	Width, Height int
}

type BootViewModel struct {
	progress progress.Model
}

type HomeViewModel struct {
	SelectedPane
	TabIndex
}

const (
	BootView Screen = iota
	HomeView
)

const (
	LeftPane SelectedPane = iota
	RightPane
)

type Model struct {
	Screen
	Dim

	BootViewModel
	HomeViewModel
}
