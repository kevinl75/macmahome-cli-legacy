package cli

import (
	"strings"

	"github.com/kevinl75/macmahome/core"
	"github.com/spf13/cobra"
)

// sumCmd represents the sum command
var todoAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo.",
	Long:  `Add a new todo.`,
	Run: func(cmd *cobra.Command, args []string) {
		todoContent := strings.Join(args, " ")
		todo := core.NewTodo(todoContent)
		todo.PrintTodo()

		saveTodo(todo)
	},
}

func saveTodo(newTodo core.Todo) {
	var todos []core.Todo = core.LoadTodosFromJson("todos.json")
	todos = append(todos, newTodo)
	core.DumpTodosToJson(todos, "todos.json")
}

func init() {
	todoCmd.AddCommand(todoAddCmd)
}
