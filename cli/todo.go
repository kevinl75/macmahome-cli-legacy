/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cli

import (
	"github.com/spf13/cobra"
)

// noteCmd represents the sum command
var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Add, list or remove your todos.",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(todoCmd)
}
