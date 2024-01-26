package util

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/android-lewis/note-taker/pkg/config"
)

type Index struct {
	ID   int64
	Path string
}

func SearchIndex(fileName string, ID int64) (Index, error) {
	indexMap := ReadJSON(fileName)
	i, j := 0, len(indexMap)-1

	for i <= j {
		h := int(uint(i+j) >> 1)
		if indexMap[h].ID == ID {
			return indexMap[h], nil
		} else if indexMap[h].ID < ID {
			i = h + 1
		} else {
			j = h - 1
		}
	}

	return Index{ID: -1, Path: ""}, fmt.Errorf("could not find index")
}

func ReadJSON(fileName string) []Index {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var filteredData []Index
	var data Index
	decoder.Token()

	for decoder.More() {
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println(err)
		}

		filteredData = append(filteredData, data)
	}

	return filteredData
}

func WriteJSON(fileName string, data []Index) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println(err)
	}
}

func ClearIndex(fileName string) {
	if err := os.Remove(fileName); err != nil {
		fmt.Printf("Failed to delete: %v", err)
	}
}

func AddToIndex(id int64, path string) {
	// Get the existing index list
	indexMap := ReadJSON(config.IndexPath)

	newindex := Index{ID: id, Path: path}
	indexMap = append(indexMap, newindex)

	WriteJSON(config.IndexPath, indexMap)
}

func OpenEditor(filePath string, readonly bool) *exec.Cmd {
	vi := "vim"
	var cmd *exec.Cmd

	path, err := exec.LookPath(vi)

	if err != nil {
		fmt.Printf("Error %s while looking up for %s!! Have you got Vim installed?", path, vi)
	}

	if readonly {
		cmd = exec.Command(path, "-M", filePath) // -M opens Vim in Readonly mode
	} else {
		cmd = exec.Command(path, filePath)
	}

	return cmd
}

func HandleCmd(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()

	if err != nil {
		fmt.Printf("Start failed: %s", err)
	}

	fmt.Printf("Waiting for command to finish.\n")

	err = cmd.Wait()

	if err != nil {
		return fmt.Errorf("command finished with error %w", err)
	}

	return nil
}
