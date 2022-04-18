package main

import "github.com/urfave/cli/v2"

var commands = []*cli.Command{
	searchArtists,
	searchAlbums,
	searchTracks,
}

var searchArtists = &cli.Command{
	Name:   "artist",
	Usage:  "search artist by keyword",
	Action: executeQuery,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
		},
	},
}

var searchAlbums = &cli.Command{
	Name:   "album",
	Usage:  "search album by keyword",
	Action: executeQuery,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
		},
	},
}

var searchTracks = &cli.Command{
	Name:   "artist",
	Usage:  "search track by keyword",
	Action: executeQuery,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
		},
	},
}
