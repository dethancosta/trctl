/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getScheduleCmd represents the getSchedule command
var getScheduleCmd = &cobra.Command{
	Use:   "getSchedule",
	Short: "Get today's schedule",
	Long: `Send a request to the timeruler server to get today's schedule as a formatted string.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getSchedule called")
	},
}

func init() {
	rootCmd.AddCommand(getScheduleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getScheduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getScheduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
