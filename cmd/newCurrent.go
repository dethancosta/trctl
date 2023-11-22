/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCurrentCmd represents the newCurrent command
var newCurrentCmd = &cobra.Command{
	Use:   "newCurrent",
	Short: "Update the current task",
	Long: `Send a request to the timeruler server to update the 
current task until a given time.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("newCurrent called")
	},
}

func init() {
	rootCmd.AddCommand(newCurrentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCurrentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCurrentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
