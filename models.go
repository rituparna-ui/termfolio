package main

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
)

type Screen int
type tickMsg time.Time

type Dim struct {
	Width, Height int
}

type BootViewModel struct {
	progress progress.Model
}

const (
	BootView Screen = iota
	HomeView
)

type Model struct {
	Screen
	Dim

	BootViewModel
}
