package view

import "github.com/charmbracelet/lipgloss"

var (
	mainColor         = "111"
	secondaryColor    = "0"
	helpColor         = "241"
	tableGitmojiStyle = lipgloss.NewStyle().
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(secondaryColor))
	tableCommitStyle = lipgloss.NewStyle().
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(secondaryColor))
	inputStyle = lipgloss.NewStyle().
			Align(lipgloss.Left, lipgloss.Center).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(secondaryColor)).
			Width(34)
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(helpColor))
)

func focus(s lipgloss.Style, isFocus bool) lipgloss.Style {
	var color string
	if isFocus {
		color = mainColor
	} else {
		color = secondaryColor
	}
	return s.BorderForeground(lipgloss.Color(color))
}
