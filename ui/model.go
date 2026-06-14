package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gogita/ui/keys"
	"gogita/ui/styles"
	"os/exec"
	"strings"
	"time"
)

type branchMsg string
type branchesMsg []string
type errMsg string
type commitMsg []string
type tickMsg time.Time

func getRecentCommits() tea.Cmd {
	return func() tea.Msg {
		out, err := exec.Command("git", "log", "--oneline", "-15").Output()
		if err != nil {
			return errMsg(fmt.Sprintf("Error getting recent commits: %s", err))
		}
		return commitMsg(strings.Split(string(out), "\n"))
	}
}

func getBranchName() tea.Cmd {
	return func() tea.Msg {
		out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		if err != nil {
			return errMsg(fmt.Sprintf("Error getting branch name: %s", err))
		}
		return branchMsg(strings.TrimSpace(string(out)))
	}
}

func fetchAllBranches() tea.Cmd {
	return func() tea.Msg {
		out, err := exec.Command("git", "branch").Output()
		if err != nil {
			return errMsg(fmt.Sprintf("Error getting branch name: %s", err))
		}
		return branchesMsg(strings.Split(strings.TrimSuffix(string(out), "\n"), "\n")) //removes the trailing
	}
}

type Model struct {
	width       int
	height      int
	branchName  string
	commits     []string
	allBranches []string
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		getBranchName(),
		getRecentCommits(),
		fetchAllBranches())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case branchMsg:
		m.branchName = string(msg)
	case branchesMsg:
		m.allBranches = []string(msg)
	case errMsg:
		m.branchName = fmt.Sprintf("Error: %s", string(msg))
	case commitMsg:
		m.commits = []string(msg)
	case tea.KeyMsg:
		switch {
		case keys.IsQuit(msg):
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			return m, cmd
		}
	}
	return m, nil
}

func renderCommits(m Model, innerWidth int) []string {
	var commitLines []string
	for i, commit := range m.commits {
		if i%2 == 0 {
			commitLines = append(commitLines, styles.CommitEven.Width(innerWidth).Render(commit))
		} else {
			commitLines = append(commitLines, styles.CommitOdd.Width(innerWidth).Render(commit))
		}
	}
	return commitLines
}

func (m Model) View() string {

	mainWidth := m.width * 3 / 4
	sideWidth := m.width / 4

	//render the branches
	var branches []string
	maxLen := 0
	for _, branch := range m.allBranches {
		if len(branch) > maxLen {
			maxLen = len(branch)
		}
	}
	for _, branch := range m.allBranches {
		branches = append(branches, styles.Button.Width(maxLen+8).Render(branch))
	}

	mainContent := lipgloss.JoinVertical(
		lipgloss.Center,
		//render the branches no empty
		lipgloss.JoinVertical(
			lipgloss.Center,
			branches...,
		),
	)

	// MAIN PANE holds the branch name
	mainPane := styles.MainPanel.Width(mainWidth).Height(m.height).Render(
		lipgloss.Place(
			mainWidth,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			mainContent,
		),
	)

	innerWidth := sideWidth - 2

	//fetch the commits
	var commitLines = renderCommits(m, innerWidth)

	// Side Pane, holds the commits
	sidePane := styles.SidePanel.
		Width(sideWidth - 2).
		Height(m.height).
		Render(strings.Join(commitLines, "\n"))

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		sidePane,
		mainPane,
	)
}
