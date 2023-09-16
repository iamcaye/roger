package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/iamcaye/roger/models"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/spf13/cobra"
)


func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readNotesFile() ([]models.Note, error) {
	data_str, err := os.ReadFile("notes.json")
	check(err)
	var tmp_notes []models.Note

	err = json.Unmarshal([]byte(data_str), &tmp_notes)

	return tmp_notes, err
}

func printNote(note models.Note) {
    title  := fmt.Sprintf("[ %v ]", note.Title)
    Box := box.New(box.Config{Px: 1, Py: 0, Type: "Classic", Color: "White", TitlePos: "Top", AllowWrapping: true})
    Box.Print(title, note.Body)
}

/*func printNotes(notes []models.Note) {
    for _, note := range notes {
        printNote (note)
    }
}*/

func listNote (note models.Note) {
    fmt.Printf("[%v] %v\n", note.Id, note.Title)
}

func ListNotes (notes []models.Note) {
    for _, note := range notes {
        listNote (note)
    }
}

func getNoteById (notes []models.Note, id int) (models.Note, error) {
    for _, note := range notes {
        if note.Id == id {
            return note, nil
        }
    }
    return models.Note{}, errors.New("Note not found")
}

func getNoteBySlug (notes []models.Note, slug string) (models.Note, error) {
    for _, note := range notes {
        if note.Slug == slug {
            return note, nil
        }
    }
    return models.Note{}, errors.New("Note not found")
}

func getNotesByCategory (notes []models.Note, category string) ([]models.Note, error) {
    var tmp_notes []models.Note
    for _, note := range notes {
        if note.Category == category {
            tmp_notes = append(tmp_notes, note)
        }
    }
    if len(tmp_notes) > 0 {
        return tmp_notes, nil
    }
    return tmp_notes, errors.New("No notes found")
}

func ReadNotes(cmd *cobra.Command, args []string) {
    log.SetPrefix("notes.go: ")
    log.SetFlags(log.LstdFlags | log.Lshortfile)
	notes, err := readNotesFile()
	check(err)
	if len(args) > 0 {
		n, err := strconv.Atoi(args[0])
        check(err)
        note, err := getNoteById(notes, n)
        check(err)
        printNote(note)

	} else {
        ListNotes(notes)
	}
}
