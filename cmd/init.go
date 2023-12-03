package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	tr "github.com/dethancosta/tr-cli/utils"
	"github.com/muesli/go-app-paths"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a config file and schedule build file",
	Long: `Initialize a config file and schedule build file for using tr-cli with a timerule server.
$ tr-cli init --server <server> --username <username> --password <password> --build <buildfile>
Username, password and buildfile are optional.`,
	Run: func(cmd *cobra.Command, args []string) {
		scope := gap.NewScope(gap.User, "timeruler")
		configDir := tr.GetConfigDir()
		
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			fmt.Println("Error creating config directory: ", err.Error())
		}

		// TODO conditional error handling
		_, err = os.OpenFile(configDir + "/config.json", os.O_RDWR, 0644)
		if err == nil {
			fmt.Println("Config file already exists. Use config command instead")
			os.Exit(1)
		}

		if buildFile == "" {
			buildFile, err  = scope.DataPath("build.csv")
			if err != nil {
				buildFile = ""
			}
		}
		config := map[string]string{
			"server": serverUrl,
			"user": user,
			"password": password,
			"build": buildFile,
		}
		configBytes, err := json.Marshal(config)
		if err != nil {
			fmt.Println("Error marshalling config data: ", err.Error())
			os.Exit(1)
		}

		
		configFile, err := os.Create(configDir + "/config.json")
		if err != nil {
			fmt.Println("Error creating config file: ", err.Error())
			os.Exit(1)
		}
		defer configFile.Close()
		_, err = io.Copy(configFile, bytes.NewReader(configBytes))
		if err != nil {
			fmt.Println("Error writing config file: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("Config file created successfully")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
