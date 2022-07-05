package cli

import (
	"strings"

	"github.com/kevinl75/macmahome/core"
	"github.com/spf13/cobra"
)

// noteAddCmd represent the 'note add' command to add a note.
var noteAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new notes.",
	Long:  `Add a new notes.`,
	Run: func(cmd *cobra.Command, args []string) {
		noteContent := strings.Join(args, " ")
		note := core.NewNotes(noteContent)
		note.PrintNote()

		saveNote(note)
	},
}

func saveNote(newNote core.Note) {
	var notes []core.Note = core.LoadNotesFromJson("notes.json")
	notes = append(notes, newNote)
	core.DumpNotesToJson(notes, "notes.json")
}

func init() {
	noteCmd.AddCommand(noteAddCmd)
}
