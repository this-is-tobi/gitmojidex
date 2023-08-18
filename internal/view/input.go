package view

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func newInput(placeholder string, focus bool) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.CharLimit = 156
	ti.Width = 30
	ti.PromptStyle.Align(lipgloss.Left)
	ti.TextStyle.Align(lipgloss.Left)
	ti.Cursor.Style.Align(lipgloss.Left)
	if focus {
		ti.Focus()
	}
	return ti
}
