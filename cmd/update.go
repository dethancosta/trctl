package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	tr "github.com/dethancosta/trctl/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update one or more tasks",
	Long: `Send a request to the timeruler server to update one or more
given tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO test
		config := tr.GetConfig()
		serverUrl, ok := config["server"]
		if !ok {
			serverUrl = tr.DefaultServerUrl
		}

		var tasks []tr.Task
		var err error
		if filename != "" {
			tasks, err = tr.TasksFromCsv(filename)
			if err != nil {
				fmt.Println("Couldn't build tasks from file: ", err.Error())
				os.Exit(1)
			}
		} else {
			// TODO get tasks from stdin
			fmt.Println("Give up to 16 tasks ('done' + Enter to finish):")
			input := bufio.NewReader(os.Stdin)
			now := time.Now()
			for i := 1; i <= 16; i++ {
				fmt.Printf("--Task %d--\n", i)
				fmt.Printf("Description: ")
				desc, err := input.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading input: ", err.Error())
					os.Exit(1)
				} else if desc == "done\n" {
					break
				}
				desc = strings.TrimSuffix(desc, "\n")
				fmt.Printf("Start Time (hh:mm:ss): ")
				startStr, err := input.ReadString('\n')
				startStr = strings.TrimSuffix(startStr, "\n")
				if err != nil {
					fmt.Println("Error reading input: ", err.Error())
					os.Exit(1)
				}
				start, err := time.Parse(time.TimeOnly, startStr)
				start = time.Date(now.Year(), now.Month(), now.Day(), start.Hour(), start.Minute(), 0, 0, time.Local)
				if err != nil {
					fmt.Println("Couldn't parse time: ", err.Error())
					os.Exit(1)
				}
				fmt.Printf("End Time (hh:mm:ss): ")
				endStr, err := input.ReadString('\n')
				endStr = strings.TrimSuffix(endStr, "\n")
				if err != nil {
					fmt.Println("Error reading input: ", err.Error())
					os.Exit(1)
				}
				end, err := time.Parse(time.TimeOnly, endStr)
				if err != nil {
					fmt.Println("Couldn't parse time: ", err.Error())
					os.Exit(1)
				}
				end = time.Date(now.Year(), now.Month(), now.Day(), end.Hour(), end.Minute(), 0, 0, time.Local)
				fmt.Printf("Tag: ")
				tag, err := input.ReadString('\n')
				tag = strings.TrimSuffix(tag, "\n")
				if err != nil {
					fmt.Println("Error reading input: ", err.Error())
					os.Exit(1)
				}
				tasks = append(tasks, tr.Task{
					Description: desc,
					StartTime:   start,
					EndTime:     end,
					Tag:         tag,
				})
			}
		}

		req, err := json.Marshal(tasks)
		if err != nil {
			fmt.Println("Error marshalling request: ", err)
		}

		resp, err := http.Post(
			serverUrl+"/update",
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
			body, _ := io.ReadAll(resp.Body)
			fmt.Println("Body: ", string(body))
			return
		}
		fmt.Println("Schedule updated")
		fmt.Println("Tasks:")
		// TODO align formatting
		for _, task := range tasks {
			fmt.Printf("%s\t%s\t%s\t%s\n", task.Description, task.StartTime.Format(time.TimeOnly), task.EndTime.Format(time.TimeOnly), task.Tag)
		}
	},
}

var filename string

func init() {
	updateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		command.Flags().MarkHidden("server")
		command.Flags().MarkHidden("user")
		command.Flags().MarkHidden("password")
		command.Flags().MarkHidden("build")
		command.Parent().HelpFunc()(command, strings)
	})
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&filename, "file", "f", "", "File with new tasks")
}
