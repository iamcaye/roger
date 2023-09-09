/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/iamcaye/roger/models"
	"github.com/spf13/cobra"
)

var notes []models.Note = []models.Note{}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readNotesFile() ([]models.Note, error) {
	fmt.Println("Opening file")
	data_str, err := os.ReadFile("notes.json")
	check(err)
	var tmp_notes []models.Note

	err = json.Unmarshal([]byte(data_str), &tmp_notes)

	return tmp_notes, err
}

func readNotes(cmd *cobra.Command, args []string) {
	notes, err := readNotesFile()
	check(err)
	if len(args) > 0 {
		fmt.Println("args", args)
		n, err := strconv.Atoi(args[0])
		check(err)
		fmt.Println(notes[n])

	} else {
		fmt.Println(notes)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "roger [command] [--flags]",
	Short: "note taking cli tool",
	Long:  `This is a note taking cli app for the best developers in the world`, // this execute when running only the command
	Run:   readNotes,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
