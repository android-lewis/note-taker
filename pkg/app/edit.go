package app

import (
	"github.com/android-lewis/note-taker/pkg/config"
	"github.com/android-lewis/note-taker/pkg/util"
)

func EditNote(id int64) error {
	index, _, err := util.SearchIndex(config.IndexPath, id)
	if err != nil {
		return err
	}

	cmd := util.OpenEditor(index.Path, false)
	err = util.HandleCmd(cmd)
	if err != nil {
		return err
	}

	return nil
}
