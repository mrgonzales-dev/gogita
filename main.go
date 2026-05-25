package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"gogita/ui"
)

func main() {
	p := tea.NewProgram(ui.Model{})
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error running program:", err)
		os.Exit(1)
	}
}
