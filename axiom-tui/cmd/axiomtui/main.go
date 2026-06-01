package main

import (
	"fmt"
	"os"

	"axiomtui/internal/tui/app"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model, err := app.NewAppModel()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create model: %v\n", err)
		os.Exit(1)
	}
	program := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "axiomtui: %v\n", err)
		os.Exit(1)
	}
}
