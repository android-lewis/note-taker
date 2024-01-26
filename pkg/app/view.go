package app

import (
	"fmt"
	"strconv"

	"github.com/android-lewis/note-taker/pkg/config"
	"github.com/android-lewis/note-taker/pkg/util"
)

func ViewNote(id string) error {
	// Find id in index list and open path in vim
	// 1. Find ID in index
	// 2. Return path
	// 3. Open Editor in readonly mode
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("id provided is not valid %s", err)
	}

	index, _, err := util.SearchIndex(config.IndexPath, intId)
	if err != nil {
		return fmt.Errorf("could not find note with id: %v", index)
	}

	cmd := util.OpenEditor(index.Path, true)
	err = util.HandleCmd(cmd)
	if err != nil {
		return err
	}

	return nil
}
