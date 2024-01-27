package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"syscall"

	"github.com/android-lewis/note-taker/pkg/app"
	"github.com/urfave/cli/v2"
)

func main() {
	var noteName string
	var noteMessage string
	var messageLength int

	syscall.Umask(0)
	os.Mkdir("./data", 0744)

	app := &cli.App{
		Name:  "note-taker",
		Usage: "Take some quick notes and store them",
		Commands: []*cli.Command{
			{
				Name:  "new",
				Usage: "add a new note",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Usage:       "Name for the note",
						Aliases:     []string{"n"},
						Destination: &noteName,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "message",
						Usage:       "Inline message for note",
						Aliases:     []string{"m"},
						Destination: &noteMessage,
					},
				},
				Action: func(cCtx *cli.Context) error {

					if cCtx.String("message") != "" {
						// Process note message from specified text
						err := app.CreateInlineNote(noteName, noteMessage)

						if err != nil {
							return err
						}
					} else {
						// Launches Vim
						err := app.CreateNote(noteName)

						if err != nil {
							return err
						}
					}

					fmt.Println("New note created")
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "view all notes",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "length",
						Usage:       "Note Preview Length",
						Aliases:     []string{"l"},
						Value:       50,
						Destination: &messageLength,
						DefaultText: "50 characters",
					},
				},
				Action: func(cCtx *cli.Context) error {
					// Function to list all notes
					app.ListNotes(messageLength)
					return nil
				},
			},
			{
				Name:  "view",
				Usage: "view note with given ID",
				Action: func(cCtx *cli.Context) error {
					// Function to list all notes
					err := app.ViewNote(cCtx.Args().Get(0))
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "edit",
				Usage: "edit note with given ID",
				Action: func(cCtx *cli.Context) error {
					// Function to edit notes
					i, err := strconv.ParseInt(cCtx.Args().Get(0), 10, 64)
					if err != nil {
						return fmt.Errorf("%s is not a valid ID", cCtx.Args().Get(0))
					}
					app.EditNote(i)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "delete note with given ID",
				Action: func(cCtx *cli.Context) error {
					// Function to list all notes
					i, err := strconv.ParseInt(cCtx.Args().Get(0), 10, 64)
					if err != nil {
						return fmt.Errorf("%s is not a valid ID", cCtx.Args().Get(0))
					}
					err = app.DeleteNote(i)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
