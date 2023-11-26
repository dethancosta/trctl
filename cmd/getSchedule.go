/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getScheduleCmd represents the getSchedule command
var getScheduleCmd = &cobra.Command{
	Use:   "getSchedule",
	Short: "Get today's schedule",
	Long: `Send a request to the timeruler server to get today's schedule as a formatted string.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getSchedule called")

		// TODO use address from config, then fallback to default
		resp, err := http.Get("http://localhost:6576/get")
		if err != nil {
			fmt.Println("Error getting schedule: ", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			fmt.Println("No schedule set")
			return
		} else if resp.StatusCode != http.StatusOK {
			fmt.Println("Error getting schedule: ", resp.Status)
			os.Exit(1)
		}

		var schedule map[string]string

		err = json.NewDecoder(resp.Body).Decode(&schedule)

		if err != nil {
			fmt.Println("Error parsing schedule: ", err)
			os.Exit(1)
		}

		fmt.Println("--- Schedule ---\n" + schedule["Schedule"])
	},
}

func init() {
	rootCmd.AddCommand(getScheduleCmd)
}
