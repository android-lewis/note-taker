package app

import (
	"github.com/android-lewis/note-taker/pkg/config"
	"github.com/android-lewis/note-taker/pkg/util"
)

// Takes in the ID of the note from list and
// Removes it from the persistent storage and the index
func DeleteNote(id int64) error {
	// Find index with key
	// Delete from peristent storage
	// Remove index entry
	index, i, err := util.SearchIndex(config.IndexPath, id)
	if err != nil {
		return err
	}

	util.RemoveFile(index.Path)
	err = util.RemoveFromIndex(i)
	if err != nil {
		return err
	}

	return nil
}
