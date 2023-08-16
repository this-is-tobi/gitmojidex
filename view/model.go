package view

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	state        sessionState
	gitmojiTable table.Model
	commitTable  table.Model
	pathInput    textinput.Model
	userInput    textinput.Model
}

func newModel(path textinput.Model, user textinput.Model, gitmojis table.Model, commits table.Model) model {
	m := model{state: pathView}
	m.pathInput = path
	m.userInput = user
	m.gitmojiTable = gitmojis // table.New()
	m.commitTable = commits   // table.New()
	return m
}
