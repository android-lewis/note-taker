package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/android-lewis/note-taker/pkg/config"
	"github.com/android-lewis/note-taker/pkg/util"
)

// Cleans all special characters from the note name,
// so we don't violate file naming convention or cause side-effects
func cleanNoteName(name string) string {
	name = strings.ToLower(name)
	replacer := strings.NewReplacer(
		" ", "",
		"/", "",
		"\\", "",
		".", "",
		"*", "",
		"?", "",
		",", "")

	return replacer.Replace(name)
}

func generateFileName(name string) (string, int64) {
	key := time.Now().UnixMicro()        // Uses unix microseconds time as the "key"
	formattedName := cleanNoteName(name) // Remove any spaces and special characters from the note name
	if formattedName == "" {
		fmt.Println("Empty name")
		formattedName = "new_note"
	}
	filename := fmt.Sprintf("%d_%s.%s", key, formattedName, config.Extension)
	dstPath := filepath.Join(config.DataDir, filename)

	return dstPath, key
}

func createFile(name string) *os.File {
	dstPath, key := generateFileName(name)
	file, err := os.Create(dstPath)

	if err != nil {
		fmt.Printf("Error %s while creating file", err)
	}

	util.AddToIndex(key, dstPath) //Add an index to our list with the "key" and destination path

	return file
}

func CreateInlineNote(name, note string) error {

	file := createFile(name)
	_, err := file.WriteString(note)
	if err != nil {
		return err
	}
	return nil
}

func CreateNote(name string) error {
	file := createFile(name)
	cmd := util.OpenEditor(file.Name(), false)
	err := util.HandleCmd(cmd)
	if err != nil {
		return err
	}

	return nil
}
