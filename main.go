package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	if err := newApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "spq"
	app.Usage = "call spotify API and retrieve album/track info"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "artist",
			Aliases: []string{"r"},
			Usage:   "artist name",
		},
		&cli.StringFlag{
			Name:    "album",
			Aliases: []string{"l"},
			Usage:   "show query result",
		},
		&cli.StringFlag{
			Name:    "track",
			Aliases: []string{"t"},
			Usage:   "show query result",
		},
	}
	app.Action = executeQuery
	return app
}
