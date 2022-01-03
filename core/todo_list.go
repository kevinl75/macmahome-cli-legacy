package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"time"
)

type TodoList []Todo

func (n TodoList) Len() int {
	return len(n)
}

func (n TodoList) Less(i, j int) bool {
	creationDateI, err := time.Parse("2006-01-02", n[i].CreationDate)
	if err != nil {
		log.Fatal(err)
	}
	creationDateJ, err := time.Parse("2006-01-02", n[j].CreationDate)
	if err != nil {
		log.Fatal(err)
	}
	return creationDateI.After(creationDateJ)
}

func (n TodoList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (todos TodoList) PrintAllTodosList() {
	if len(todos) == 0 {
		fmt.Printf("No todos to print...\n")
	}

	firstDate := (todos)[0].CreationDate
	for i := 0; i < len(todos); i++ {
		if firstDate != (todos)[i].CreationDate {
			firstDate = (todos)[i].CreationDate
			fmt.Printf("\n")
		}
		fmt.Printf("Todos for day %s:\n", firstDate)
		(todos)[i].PrintTodo()
	}
}

func (todos TodoList) PrintTodosListForDay(day string) {
	fmt.Printf("Todos for day %s:\n\n", day)
	for i := 0; i < len(todos); i++ {
		if todos[i].CreationDate == day {
			todos[i].PrintTodo()
		}
	}
}

func LoadTodosFromJson(jsonFilePath string) []Todo {

	content, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var todos []Todo
	err = json.Unmarshal(content, &todos)
	if err != nil {
		log.Fatal(err)
	}
	return todos
}

func DumpTodosToJson(todos []Todo, jsonFilePath string) {

	sort.Sort(TodoList(todos))

	buffer, err := json.Marshal(todos)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(jsonFilePath, buffer, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
