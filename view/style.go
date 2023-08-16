package view

import "github.com/charmbracelet/lipgloss"

var (
	// modelStyle = lipgloss.NewStyle().
	// 		Align(lipgloss.Center, lipgloss.Center).
	// 		BorderStyle(lipgloss.NormalBorder()).
	// 		BorderForeground(lipgloss.Color("240"))
	// focusedModelStyle = lipgloss.NewStyle().
	// 			Align(lipgloss.Center, lipgloss.Center).
	// 			BorderStyle(lipgloss.NormalBorder()).
	// 			BorderForeground(lipgloss.Color("201")).
	// 			Height(10)
	tableGitmojiStyle = lipgloss.NewStyle().
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240")).
				Height(10)
	tableCommitStyle = lipgloss.NewStyle().
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240")).
				Height(28)
	inputStyle = lipgloss.NewStyle().
			Align(lipgloss.Left, lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			Width(40)
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

func focus(s lipgloss.Style) lipgloss.Style {
	return s.BorderForeground(lipgloss.Color("201"))
}

func unfocus(s lipgloss.Style) lipgloss.Style {
	return s.BorderForeground(lipgloss.Color("240"))
}
