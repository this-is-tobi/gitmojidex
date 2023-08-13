/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/this-is-tobi/gitmojidex/utils"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Get repository history.",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Println("Error while getting 'path' flag", err)
		}

		user, err := cmd.Flags().GetString("user")
		if err != nil {
			fmt.Println("Error while getting 'user' flag", err)
		}

		rawCommits := utils.GetRawHistory(path)
		commits := utils.GetFormatedHistory(rawCommits, user)
		utils.TableHistory(commits)
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
