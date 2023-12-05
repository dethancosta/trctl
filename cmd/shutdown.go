/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"syscall"

	tr "github.com/dethancosta/tr-cli/utils"
	"github.com/spf13/cobra"
)

// shutdownCmd represents the shutdown command
var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Shutdown a local timeruler server",
	Long: `Shutdown a local timeruler server running as a daemon.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO delete this line
		fmt.Println("shutdown called")

		pidStr, ok := tr.GetConfig()["pid"]
		if !ok {
			fmt.Println("No pid found. Are you sure you're running a local server?")
			os.Exit(1)
		}
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			fmt.Println("Couldn't get pid from config file: ", err.Error())
			os.Exit(1)
		}
		err = syscall.Kill(pid, syscall.SIGINT)
		fmt.Printf("Process killed at pid %d\n", pid)
	},
}

func init() {
	rootCmd.AddCommand(shutdownCmd)
}
