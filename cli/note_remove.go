package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/kevinl75/macmahome/core"
	"github.com/spf13/cobra"
)

// noteRmCmd represents the 'note rm' command to remove note(s).
var noteRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or multiple notes.",
	Long: `The 'rm' note subcommand enable you to remove one or multiple note. Notes can 
be removed by three differents ways: based on the creation date of the Note, based on its
id, or simply removed all the note. Be carreful, Note removed are erased permanently.`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: improve the code from here to line 48
		nbFlagSet := 0
		inputDayFlag, err := cmd.Flags().GetString("day")
		if err != nil {
			log.Fatal(err)
		}
		if inputDayFlag != "" {
			nbFlagSet += 1
		}

		inputAllFlag, err := cmd.Flags().GetBool("all")
		if err != nil {
			log.Fatal(err)
		}
		if inputAllFlag == true {
			nbFlagSet += 1
		}

		inputIdFlag, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Fatal(err)
		}
		if inputIdFlag != "" {
			nbFlagSet += 1
		}

		if nbFlagSet > 1 {
			fmt.Println("You cannot use multiple flags at the same time.")
			os.Exit(1)
		}

		if !askForUserValidation() {
			os.Exit(1)
		}

		if inputAllFlag {
			var emptyNoteList []core.Note = make([]core.Note, 0)
			core.DumpNotesToJson(emptyNoteList, "notes.json")
		}
		if inputDayFlag != "" {
			removeNotesOfDay(inputDayFlag)
		}
		if inputIdFlag != "" {
			removeNote(inputIdFlag)
		}
	},
}

// Remove Note(s) from NoteList based on the creation date of the note(s).
func removeNotesOfDay(dayFlag string) {
	var notes core.NoteList = core.LoadNotesFromJson("notes.json")
	notes.RemoveNotesByDay(dayFlag)
	core.DumpNotesToJson(notes, "notes.json")
}

// Remove Note from NoteList based on Note id.
func removeNote(idFlag string) {
	var notes core.NoteList = core.LoadNotesFromJson("notes.json")
	if notes.RemoveNotesById(idFlag) {
		fmt.Printf("Note with id %s was removed.", idFlag)
	}
	core.DumpNotesToJson(notes, "notes.json")
}

// Ask for user validation before deleting items.
func askForUserValidation() bool {
	fmt.Printf("You're going to remove one ore more notes. Do you confirm? ('y' or 'yes')\n")
	var userAck string
	fmt.Scanln(&userAck)
	if userAck == "y" || userAck == "yes" {
		return true
	} else {
		return false
	}
}

func init() {
	noteCmd.AddCommand(noteRmCmd)
	noteRmCmd.Flags().StringP("id", "i", "", "Remove note with the given id.")
	noteRmCmd.Flags().StringP("day", "d", "", "Remove notes for a particular days.")
	noteRmCmd.Flags().BoolP("all", "a", false, "Remove all your notes.")
}
