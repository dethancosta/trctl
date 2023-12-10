/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	tr "github.com/dethancosta/trctl/utils"
	"github.com/spf13/cobra"
)

// getconfigCmd represents the getconfig command
var getconfigCmd = &cobra.Command{
	Use:   "getconfig",
	Short: "Print the contents of the timeruler config file",
	Long: `Print the contents of the config file. Format is a pretty-printed json object.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO delete this line
		fmt.Println("getconfig called")

		config := tr.GetConfig()
		if config == nil || len(config) == 0 {
			fmt.Println("No config file found.")
		} else {
			fmt.Println("{")
			for k,v := range config {
				fmt.Printf("\t\"%s\": \"%s\"\n", k, v)
			}
			fmt.Println("}")
		}
	},
}

func init() {
	rootCmd.AddCommand(getconfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getconfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getconfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
