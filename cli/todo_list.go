package cli

import (
	"log"
	"time"

	"github.com/kevinl75/macmahome/core"
	"github.com/spf13/cobra"
)

// noteListCmd represents the 'note list' command to list note(s).
var todoListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your daily notes.",
	Long: `List your notes. By default, the notes created today are print.
You can change this behavior thanks to the command flag.`,
	Run: func(cmd *cobra.Command, args []string) {

		inputDayFlag, err := cmd.Flags().GetString("day")
		if err != nil {
			log.Fatal(err)
		}
		inputAllFlag, err := cmd.Flags().GetBool("all")
		if err != nil {
			log.Fatal(err)
		}

		var todos core.TodoList = core.LoadTodosFromJson("todos.json")
		if inputAllFlag {
			todos.PrintAllTodosList()
		} else {
			todos.PrintTodosListForDay(inputDayFlag)
		}
	},
}

func init() {
	todoCmd.AddCommand(todoListCmd)
	todoListCmd.Flags().StringP("day", "d", time.Now().Format("2006-01-02"), "Print this day's todo(s). (format YYYY-MM-DD)")
	todoListCmd.Flags().BoolP("all", "a", false, "List all your todos.")
}
