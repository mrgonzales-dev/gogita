package keys

import tea "github.com/charmbracelet/bubbletea"

func IsQuit(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyCtrlC || msg.String() == "q"
}
