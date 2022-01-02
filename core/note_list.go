package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"time"
)

type NoteList []Note

func (n NoteList) Len() int {
	return len(n)
}

func (n NoteList) Less(i, j int) bool {
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

func (n NoteList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (nl *NoteList) RemoveNotesByDay(date string) {
	var newNoteList NoteList = make([]Note, 0)
	for i := 0; i < len(*nl); i++ {
		if (*nl)[i].CreationDate != date {
			newNoteList = append(newNoteList, (*nl)[i])
		}
	}
	*nl = newNoteList
}

func (n *NoteList) RemoveNotesById(id string) bool {
	for i := 0; i < len(*n); i++ {
		if (*n)[i].Id == id {
			(*n) = append((*n)[:i], (*n)[i+1:]...)
			return true
		}
	}
	return false
}

func LoadNotesFromJson(jsonFilePath string) []Note {

	content, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var notes []Note
	err = json.Unmarshal(content, &notes)
	if err != nil {
		log.Fatal(err)
	}
	return notes
}

func DumpNotesToJson(notes []Note, jsonFilePath string) {

	sort.Sort(NoteList(notes))

	buffer, err := json.Marshal(notes)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(jsonFilePath, buffer, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
