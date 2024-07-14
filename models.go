package main

type Window struct {
	Width  int
	Height int
}

type Screen string

type Model struct {
	Window Window
	Screen Screen
}

func InitModel() Model {
	return Model{
		Screen: "INIT_VIEW",
	}
}
