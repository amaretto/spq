package main

import "github.com/urfave/cli/v2"

var commands = []*cli.Command{
	searchArtists,
	searchAlbums,
	searchTracks,
}

var searchArtists = &cli.Command{
	Name:  "artist",
	Usage: "search artist by keyword",
	Subcommands: []*cli.Command{
		{
			Name:   "list",
			Usage:  "list search results",
			Action: listArtist,
		},
		{
			Name:   "describe",
			Usage:  "list search results",
			Action: describeArtist,
		},
	},
}

var searchAlbums = &cli.Command{
	Name:  "album",
	Usage: "search album by keyword",
	Subcommands: []*cli.Command{
		{
			Name:   "list",
			Usage:  "list search results",
			Action: listAlbum,
		},
		{
			Name:   "describe",
			Usage:  "list search results",
			Action: describeAlbum,
		},
	},
}

var searchTracks = &cli.Command{
	Name:   "track",
	Usage:  "search track by keyword",
	Action: executeQuery,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
		},
	},
}
