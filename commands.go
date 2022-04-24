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
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "format",
					Aliases: []string{"f"},
				},
			},
		},
	},
}

var searchTracks = &cli.Command{
	Name:  "track",
	Usage: "search track by keyword",
	Subcommands: []*cli.Command{
		{
			Name:   "list",
			Usage:  "list search results",
			Action: listTrack,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "format",
					Aliases: []string{"f"},
				},
			},
		},
		{
			Name:   "describe",
			Usage:  "list search results",
			Action: describeTrack,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "format",
					Aliases: []string{"f"},
				},
			},
		},
	},
}
