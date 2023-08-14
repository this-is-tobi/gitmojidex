/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/this-is-tobi/gitmojidex/utils"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List gitmojis used in the repository.",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		user, err := cmd.Flags().GetString("user")
		// sort, err := cmd.Flags().GetString("sort")
		if err != nil {
			fmt.Println("Error while getting 'path' flag", err)
		}

		rawCommits := utils.GetRawHistory(path)
		commits := utils.GetFormatedEmoji(rawCommits, user)
		utils.TableEmoji(commits)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
