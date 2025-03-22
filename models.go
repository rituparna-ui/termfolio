package main

type Screen int

type Dim struct {
	Width, Height int
}

const (
	BootView Screen = iota
)

type Model struct {
	Screen
	Dim
}
