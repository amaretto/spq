package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

func executeQuery(c *cli.Context) error {
	// authorizaiton
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	// check args
	var args = c.Args().Slice()
	if len(args) == 0 {
		return fmt.Errorf("no target args specified. see `spq ${kind} -h` for more details")
	} else if len(args) > 1 {
		return fmt.Errorf("too many args specified. see `spq ${kind} -h` for more details")
	}

	qs := args[0]

	// ToDo change kind by specified command
	results, err := client.Search(ctx, qs, spotify.SearchTypeArtist|spotify.SearchTypeAlbum|spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}
	if results.Artists != nil {
		fmt.Println("Artists:")
		for _, item := range results.Artists.Artists {
			fmt.Println("    ", item.Name)
		}
	}

	return nil
}
