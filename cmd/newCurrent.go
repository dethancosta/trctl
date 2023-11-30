/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	tr "github.com/dethancosta/tr-cli/utils"
	"github.com/spf13/cobra"
)

// newCurrentCmd represents the newCurrent command
var newCurrentCmd = &cobra.Command{
	Use:   "newCurrent",
	Short: "Update the current task",
	Long: `Send a request to the timeruler server to update the 
current task until a given time.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := tr.GetConfig()
		serverUrl, ok := config["serverUrl"]
		if !ok {
			serverUrl = tr.DefaultServerUrl
		}
		if desc == "" {
			fmt.Println("Error: description (-d) is required")
			return
		} else if until == "" {
			fmt.Println("Error: until (-u) is required")
			return
		}
		task := struct {
			Desc  string `json:"Description"`
			Until string `json:"Until"`
			Tag   string `json:"Tag"`
		}{
			Desc:  desc,
			Until: until,
			Tag:   tag,
		}
		req, err := json.Marshal(task)
		if err != nil {
			fmt.Println("Error marshalling request: ", err)
			return
		}

		resp, err := http.Post(
			serverUrl+"/change_current",
			"application/json",
			bytes.NewBuffer(req),
		)
		if err != nil {
			fmt.Println("Error sending request: ", err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Println("Error: ", resp.Status)
			return
		}
		fmt.Println("Task updated")
	},
}

var (
	desc  string
	until string
	tag   string
)

func init() {
	rootCmd.AddCommand(newCurrentCmd)

	newCurrentCmd.Flags().StringVarP(&desc, "desc", "d", "", "Task description")
	newCurrentCmd.Flags().StringVarP(&until, "until", "u", "", "Task end time")
	newCurrentCmd.Flags().StringVarP(&tag, "tag", "t", "", "Task tag")
}
