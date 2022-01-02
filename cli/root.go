/* Main command declaration.*/
package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command macmahome when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "macmahome",
	Short: "Your personnal assistant to organise Notes and Tasks",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello CLI")
	},
}

// Execute adds all child commands to the root command and sets flags
// appropriately.Then it executes the wanted command accordingly.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("An error occured while setting the command.")
	}
}

func init() {}
