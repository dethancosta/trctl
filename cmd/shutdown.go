/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	tr "github.com/dethancosta/tr-cli/utils"
	"github.com/spf13/cobra"
)

// shutdownCmd represents the shutdown command
var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Shutdown a local timeruler server",
	Long: `Shutdown a local timeruler server running as a daemon.`,
	Run: func(cmd *cobra.Command, args []string) {

		err = tr.RemovePid()
		if err != nil {
			fmt.Println("Couldn't remove pid: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Process killed\n")
	},
}

func init() {
	rootCmd.AddCommand(shutdownCmd)
}
