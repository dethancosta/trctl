/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	tr "github.com/dethancosta/tr-cli/utils"
	"github.com/spf13/cobra"
)

// getScheduleCmd represents the getSchedule command
var getScheduleCmd = &cobra.Command{
	Use:   "getSchedule",
	Short: "Get today's schedule",
	Long:  `Send a request to the timeruler server to get today's schedule as a formatted string.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverUrl, ok := tr.GetConfig()["server"]
		if !ok {
			serverUrl = tr.DefaultServerUrl
		}
		resp, err := http.Get(serverUrl + "/get")
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

		title := "--- Schedule ---"
		fmt.Println(title + "\n" + schedule["Schedule"] + strings.Repeat("-", len(title)))
	},
}

func init() {
	rootCmd.AddCommand(getScheduleCmd)
}
