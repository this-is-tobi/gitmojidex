/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	repoPath string
	user     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitmojidex",
	Short: "A tool to get git repository stats.",
	Long: `Gitmojidex is a CLI library that gives you git information.
It will parse git files to show statistics about a repository.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitmojidex.yaml)")
	rootCmd.PersistentFlags().StringVarP(&repoPath, "path", "p", "./", "Path of the git repository")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "User for search filter with a regex pattern")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
