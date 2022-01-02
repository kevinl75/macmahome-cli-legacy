package core

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Note struct {
	Id           string `json:"id"`
	Content      string `json:"content"`
	CreationDate string `json:"create_date"`
}

func generate6DigitsRandString() string {
	var idArray []string = make([]string, 6)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		idArray[i] = strconv.Itoa(rand.Intn(10))
	}
	return strings.Join(idArray, "")
}

func NewNotes(content string) Note {
	return Note{
		generate6DigitsRandString(),
		content,
		time.Now().Format("2006-01-02"),
	}
}

func (n *Note) PrintNote() {
	fmt.Printf("\nNote id     : %s", n.Id)
	fmt.Printf("\nNote Date   : %s", n.CreationDate)
	fmt.Printf("\nNote Content: %s", n.Content)
	fmt.Printf("\n")
}
