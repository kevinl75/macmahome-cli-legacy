package core

import (
	"fmt"
	"time"
)

type Todo struct {
	Id           string `json:"id"`
	Content      string `json:"content"`
	CreationDate string `json:"creation_date"`
	IsDone       bool   `json:"is_done"`
}

func NewTodo(content string) Todo {
	return Todo{
		generate6DigitsRandString(),
		content,
		time.Now().Format("2006-01-02"),
		false,
	}
}

func (todo Todo) PrintTodo() {
	if todo.IsDone {
		fmt.Printf("x | %s | %s\n", todo.Id, todo.Content)
	} else {
		fmt.Printf("o | %s | %s\n", todo.Id, todo.Content)
	}
}

func (todo *Todo) SetDone() {
	*&todo.IsDone = !*&todo.IsDone
}
