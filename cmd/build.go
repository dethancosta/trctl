/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	// "net/http"

	tr "github.com/dethancosta/tr-cli/utils"
	"github.com/spf13/cobra"
)

var buildPath string

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

		var ok bool
		config := tr.GetConfig()
		serverUrl, ok := config["server"]
		if !ok {
			serverUrl = tr.DefaultServerUrl
		}
		if buildPath == "" {
			if buildPath, ok = config["build"]; !ok {
				fmt.Println("No build file specified in config file or as argument.")
			os.Exit(1)
			}
		}

		f, err := os.Open(buildPath)
		if err != nil {
			fmt.Println("Error opening build file:", err)
			os.Exit(1)
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fileWriter, err := writer.CreateFormFile("buildFile", buildPath)
		if err != nil {
			fmt.Println("Error reading build file:", err)
			os.Exit(1)
		}

		_, err = io.Copy(fileWriter, f)
		if err != nil {
			fmt.Println("Error reading build file:", err)
			os.Exit(1)
		}
		writer.Close()

		req, err := http.NewRequest("POST", serverUrl + "/build", body)
		if err != nil {
			fmt.Println("Error building request: ", err)
			os.Exit(1)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err := (&http.Client{}).Do(req)

		if resp.StatusCode != 200 {
			fmt.Printf("Error sending build file: Response Code '%s': %s\n", resp.Status, err)
			os.Exit(1)
		}
		fmt.Println("Build file sent successfully.")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// flags
	buildCmd.Flags().StringVarP(&buildPath, "file", "f", "", "Path to build file")
}
