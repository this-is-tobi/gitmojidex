package view

import "github.com/charmbracelet/lipgloss"

var (
	tableGitmojiStyle = lipgloss.NewStyle().
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240"))
	tableCommitStyle = lipgloss.NewStyle().
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240"))
	inputStyle = lipgloss.NewStyle().
			Align(lipgloss.Left, lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			Width(34)
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

func focus(s lipgloss.Style, isFocus bool) lipgloss.Style {
	if isFocus {
		return s.BorderForeground(lipgloss.Color("201"))
	} else {
		return s.BorderForeground(lipgloss.Color("240"))
	}
}
