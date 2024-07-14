package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if _, err := tea.NewProgram(InitModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Some error occured")
		os.Exit(1)
	}
}
