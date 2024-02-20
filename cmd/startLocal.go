package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var startlocalCmd = &cobra.Command{
	Use:   "startlcl",
	Short: "Start a local timeruler server (requires `tr-server`)",
	Long: `Start a timeruler server on your local machine. Uses port 6576,
and requires that you have timeruler installed on your PATH`,
	Run: func(cmd *cobra.Command, args []string) {
		exc := exec.Command("tr-server", "-sa", "true", "-n", ntfyId, "&")
		err := exc.Start()

		if err != nil {
			fmt.Println("Startup failed: ", err.Error())
		}
	},
}

var ntfyId string

func init() {
	rootCmd.AddCommand(startlocalCmd)

	startlocalCmd.Flags().StringVarP(&ntfyId, "ntfyId", "n", "", "ntfy.sh url path")
}
