/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cli

import (
	"github.com/spf13/cobra"
)

// noteCmd represents the sum command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Add, list or remove your notes.",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
