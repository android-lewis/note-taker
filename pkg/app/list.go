package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/android-lewis/note-taker/pkg/config"
	"github.com/cheynewallace/tabby"
)

type FileFormat struct {
	preview       string
	formattedDate string
	noteId        string
	noteName      string
}

func stripNewLines(str string) string {
	return strings.Replace(str, "\n", " ", -1)
}

func firstN(str string, length int) string {
	symbols := []rune(str)
	if length >= len(symbols) {
		return stripNewLines(str)
	}

	short := string(symbols[:length])
	return stripNewLines(short)
}

func formatFileInfo(info os.FileInfo, fileContent []byte, messageLimit int) *FileFormat {
	splitNote := strings.Split(info.Name(), ".")[0]
	return &FileFormat{
		preview:       firstN(string(fileContent), messageLimit),
		formattedDate: info.ModTime().Format(time.RFC822),
		noteId:        strings.Split(splitNote, "_")[0],
		noteName:      strings.Join(strings.Split(splitNote, "_")[1:], ""),
	}
}

func ListNotes(messageLimit int) {
	writer := tabwriter.NewWriter(os.Stdout, 25, 2, 2, ' ', 0)
	tabs := tabby.NewCustom(writer)
	tabs.AddHeader("Created", "ID", "Name", "Preview")

	err := filepath.Walk(config.DataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !info.IsDir() {

			fileContent, err := os.ReadFile(path)

			if err != nil {
				fmt.Println(err)
				return err
			}

			fileInfo := formatFileInfo(info, fileContent, messageLimit)
			tabs.AddLine(fileInfo.formattedDate, fileInfo.noteId, fileInfo.noteName, fileInfo.preview)
		}

		tabs.Print()

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
