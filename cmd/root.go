/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/this-is-tobi/gitmojidex/internal/git"
	"github.com/this-is-tobi/gitmojidex/internal/view"
)

var (
	repoPath  string
	user      string
	emojiless bool
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
	rootCmd.PersistentFlags().StringVarP(&repoPath, "path", "p", "./", "Path to the git repository.")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "User for search filter with a regex pattern.")
	rootCmd.PersistentFlags().BoolVarP(&emojiless, "emojiless", "e", false, "Disable emoji rendering in the UI")
}

func main(cmd *cobra.Command, args []string) {
	git.FetchHistory(repoPath)
	// honor emojiless flag
	git.ShowEmoji = !emojiless
	view.Render(repoPath, user)
}
