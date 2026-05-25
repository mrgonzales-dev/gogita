package keys

import tea "github.com/charmbracelet/bubbletea"

type Map struct {
	Quit      []string
	Increment []string
}

var Default = Map{
	Quit:      []string{"q", "ctrl+c"},
	Increment: []string{"up"},
}

func IsQuit(msg tea.KeyMsg) bool {
	s := msg.String()
	for _, k := range Default.Quit {
		if s == k {
			return true
		}
	}
	return false
}

func IsIncrement(msg tea.KeyMsg) bool {
	return msg.String() == Default.Increment[0]
}
