/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/iamcaye/roger/local_storage"
	"github.com/iamcaye/roger/models"
	"github.com/spf13/cobra"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func readFromFile (filename string ) models.Note {
    var note models.Note
    full_path := fmt.Sprintf("%v/%v.md", local_storage.NotesDir, filename)
    body, err := os.ReadFile(full_path)
    check(err)

    note.Body = string(body)
    note.Title = strings.Replace(filename, "-", " ", -1)
    note.Slug = strings.Replace(filename, ".md", "", -1)
    note.Id = local_storage.GetNextId()

    check(err)

    return note
}

func checkNoteExists (filename string) bool {
    full_path := fmt.Sprintf("%v/%v.md", local_storage.NotesDir, filename)
    _, err := os.Stat(full_path)
    return !os.IsNotExist(err)
}

func addNote(filename string) {
    full_path := fmt.Sprintf("%v/%v.md", local_storage.NotesDir, filename)
    cmd := exec.Command("vim", full_path)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    check(err)

    note := readFromFile(filename)
    local_storage.AddNote (note)
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new note",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            fmt.Println("Please provide a filename")
            return
        }
        if  len(args) > 1 {
            fmt.Println("Please provide only one filename separated by a dash (-)")
            return
        }
        if checkNoteExists(args[0]) {
            fmt.Println("Note already exists")
            return
        }
        addNote(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
