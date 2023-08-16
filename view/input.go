package view

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

var (
	pathPlaceholder string
	userPlaceholder string
)

func newInput(placeholder string) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40
	ti.PromptStyle.Align(lipgloss.Left)
	ti.TextStyle.Align(lipgloss.Left)
	ti.Cursor.Style.Align(lipgloss.Left)
	return ti
}
