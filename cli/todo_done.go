package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/kevinl75/macmahome/core"
	"github.com/spf13/cobra"
)

// sumCmd represents the sum command
var todoDoneCmd = &cobra.Command{
	Use:   "done",
	Short: "Set a todo to done.",
	Long:  `Set a todo to done.`,
	Run: func(cmd *cobra.Command, args []string) {
		todoId, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Fatal(err)
		}
		if todoId == "" {
			fmt.Printf("You have to set an id.\n")
			os.Exit(1)
		}

		setTodoToDone(todoId)
	},
}

func setTodoToDone(todoId string) {
	var todos []core.Todo = core.LoadTodosFromJson("todos.json")

	for i := 0; i < len(todos); i++ {
		if todoId == todos[i].Id {
			todos[i].SetDone()
		}
	}

	core.DumpTodosToJson(todos, "todos.json")
}

func init() {
	todoCmd.AddCommand(todoDoneCmd)
	todoDoneCmd.Flags().StringP("id", "i", "", "Id of the todo to set to done.")
}
