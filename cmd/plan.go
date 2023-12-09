package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Plan tomorrow's tasks",
	Long: `Send a request to the server to plan tomorrow's tasks.
A build file is sent that contains the tasks that are to be planned.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plan called")
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
