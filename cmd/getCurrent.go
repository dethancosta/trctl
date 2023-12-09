package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	tr "github.com/dethancosta/trctl/utils"
	"github.com/spf13/cobra"
)

var getCurrentCmd = &cobra.Command{
	Use:   "getCurrent",
	Short: "Get the current task",
	Long:  `Send a request to the server to get the current task.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverUrl, ok := tr.GetConfig()["server"]
		if !ok {
			serverUrl = tr.DefaultServerUrl
		}

		resp, err := http.Get(serverUrl + "/current")
		if err != nil {
			fmt.Println("Error sending request: ", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			fmt.Println("No current task")
			return
		} else if resp.StatusCode != http.StatusOK {
			fmt.Println("Error: ", resp.Status)
			return
		}

		var task map[string]struct {
			Description string `json:"Description"`
			Tag         string `json:"Tag"`
			Until       string `json:"Until"`
		}

		err = json.NewDecoder(resp.Body).Decode(&task)
		if err != nil {
			fmt.Println("Error parsing response: ", err)
			os.Exit(1)
		}
		t := task["Task"]

		fmt.Println("Current task:")
		fmt.Printf("'%s' - (%s) until %s\n", t.Description, t.Tag, t.Until)

	},
}

func init() {
	getCurrentCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		command.Flags().MarkHidden("server")
		command.Flags().MarkHidden("user")
		command.Flags().MarkHidden("password")
		command.Flags().MarkHidden("build")
		command.Parent().HelpFunc()(command, strings)
	})
	rootCmd.AddCommand(getCurrentCmd)
}
