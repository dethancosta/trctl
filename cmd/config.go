/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/muesli/go-app-paths"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Update configuration variables",
	Long: `Update variables in the configuration file:
	config -u <username> -p <password> -s <server>`,
	Run: func(cmd *cobra.Command, args []string) {
		scope := gap.NewScope(gap.User, "timeruler")
		configPath, err := scope.ConfigPath("config.json")
		if err != nil {
			fmt.Println("Error getting config path: ", err.Error())
			os.Exit(1)
		}
		configFile, err := os.OpenFile(configPath, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error opening config file: ", err.Error())
			os.Exit(1)
		}
		configBytes, err := io.ReadAll(configFile)
		configFile.Close()

		var config map[string]string
		err = json.Unmarshal(configBytes, &config)
		if err != nil {
			fmt.Println("Error reading config file: ", err.Error())
			os.Exit(1)
		}

		if _, ok := config["server"]; !ok {
			config["server"] = serverUrl
		} else if serverUrl != "" {
			config["server"] = serverUrl
		}
		if _, ok := config["user"]; !ok {
			config["user"] = user
		} else if user != "" {
			config["user"] = user
		}
		if _, ok := config["password"]; !ok {
			config["password"] = password
		} else if password != "" {
			config["password"] = password
		}
		if _, ok := config["build"]; !ok {
			config["build"] = buildFile
		} else if buildFile != "" {
			config["build"] = buildFile
		}

		configBytes, err = json.Marshal(config)
		if err != nil {
			fmt.Println("Error marshalling config: ", err.Error())
			os.Exit(1)
		}

		configFile, err = os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening config file: ", err.Error())
			os.Exit(1)
		}
		defer configFile.Close()
		_, err = configFile.Write(configBytes)
		if err != nil {
			fmt.Println("Error writing to config file: ", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
