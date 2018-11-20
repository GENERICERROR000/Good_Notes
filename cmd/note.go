// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: note,
}

func getNoteDir() string {
	if noteDir == "" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path := filepath.Clean(home + "/" + DEFAULT_NOTEPATH)
		return path
	}
	return noteDir
}

func editNote(editCommand string, noteFile string) error {
	editor := exec.Command(editCommand, noteFile)
	editor.Stdin = os.Stdin
    editor.Stdout = os.Stdout
	err := editor.Run()
	return err 
}

func note(cmd *cobra.Command, args []string) {
	workspace := getNoteDir()
	editFiles := args
	if len(args) == 0 {
		stamp := time.Now().Format("2006-01-02")
		newNoteFile := "note-" + stamp
		editFiles = append(editFiles, newNoteFile)
	}
	for _, file := range editFiles {
		noteFile := filepath.Clean(workspace + "/" + file)
		err := editNote(editCmd, noteFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("Noted")

}

func init() {
	rootCmd.AddCommand(noteCmd)

}
