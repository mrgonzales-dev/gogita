package keys

import tea "github.com/charmbracelet/bubbletea"

func IsQuit(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyCtrlC || msg.String() == "q"
}

func IsUp(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyUp || msg.Type == tea.KeyRunes && string(msg.Runes) == "k"
}

func IsDown(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyDown || msg.Type == tea.KeyRunes && string(msg.Runes) == "j"
}

func IsEnter(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyEnter
}

// refresh using f5, re-renders the tui
func IsRefresh(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyF5
}
