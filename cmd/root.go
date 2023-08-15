/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/this-is-tobi/gitmojidex/table"
)

var rootCmd = &cobra.Command{
	Use:   "gitmojidex",
	Short: "A tool to get git repository stats.",
	Long: `Gitmojidex is an interactive CLI tool that show you git informations.
It will parse git files to show statistics about a repository.`,
	Run: main,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func main(cmd *cobra.Command, args []string) {
	renderHistory()
	renderEmoji()
}
func renderHistory() {

	cols, rows := table.GetTableData("history")
	m := table.CreateTable(cols, rows)
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func renderEmoji() {
	cols, rows := table.GetTableData("emoji")
	m := table.CreateTable(cols, rows)
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Sprintln("Error running program:", err)
		os.Exit(1)
	}
}
