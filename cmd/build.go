/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	// "net/http"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Send a build file to the server to start today's schedule",
	Long: `Send a request to the server to build today's schudule
with the build file specified in the config file. The build file
is a csv file with the following format, where each line is a task:
<description>,<start time>,<end time>,<tag>
The start and end times are in the format HH:MM:SS. The tag is optional,
but the last comma is required. The description cannot be empty.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
		// resp, err := http.Post()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
