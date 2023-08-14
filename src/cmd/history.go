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
		argPath, err := cmd.Flags().GetString("path")
		argUser, err := cmd.Flags().GetString("user")
		argSort, err := cmd.Flags().GetString("sort")
		if err != nil {
			fmt.Println("Error while getting arg :", err)
		}

		rawCommits := utils.GetRawHistory(argPath)
		commits := utils.GetFormatedHistory(rawCommits, argUser)
		utils.TableHistory(commits, argSort)
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
