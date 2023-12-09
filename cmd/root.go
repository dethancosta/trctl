package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "trctl",
	Short: "CLI client for a timeruler server",
	Long: `trctl is a CLI client for a timeruler server. It is used to
get, update, and build schedules.`,
}

var (
	user      string
	password  string
	serverUrl string
	buildFile string

	err error
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&serverUrl, "server", "s", "", "Remote server address")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "Username for a remote server")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password for a remote server")
	rootCmd.MarkFlagsRequiredTogether("user", "password")

	rootCmd.PersistentFlags().StringVarP(&buildFile, "build", "b", "", "Path to the build file")
}
