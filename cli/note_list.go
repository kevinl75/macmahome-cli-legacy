package cli

import (
	"log"
	"time"

	"github.com/kevinl75/macmahome/core"
	"github.com/spf13/cobra"
)

// noteListCmd represents the 'note list' command to list note(s).
var noteListCmd = &cobra.Command{
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

		if inputAllFlag {
			listAllNotes()
		} else {
			listNotes(inputDayFlag)
		}
	},
}

func listNotes(day string) {
	var notes []core.Note = core.LoadNotesFromJson("notes.json")
	for i := 0; i < len(notes); i++ {
		if notes[i].CreationDate == day {
			notes[i].PrintNote()
		}
	}
}

func listAllNotes() {
	var notes []core.Note = core.LoadNotesFromJson("notes.json")
	for i := 0; i < len(notes); i++ {
		notes[i].PrintNote()
	}
}

func init() {
	noteCmd.AddCommand(noteListCmd)
	noteListCmd.Flags().StringP("day", "d", time.Now().Format("2006-01-02"), "Print this day's note(s). (format YYYY-MM-DD)")
	noteListCmd.Flags().BoolP("all", "a", false, "List all your notes.")
}
