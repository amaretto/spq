package main

import "github.com/urfave/cli/v2"

var commands = []*cli.Command{
	commandQuery,
}

var commandQuery = &cli.Command{
	Name:   "query",
	Usage:  "Execute query to spotify",
	Action: executeQuery,
	Flags: []cli.Flag{
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
	},
}
