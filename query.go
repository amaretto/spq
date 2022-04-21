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

	var results *spotify.SearchResult

	if c.Command.Name == "artist" {
		results, err = client.Search(ctx, qs, spotify.SearchTypeArtist)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Artists:")
		for _, item := range results.Artists.Artists {
			fmt.Println("    ", item.Name)
		}
	} else if c.Command.Name == "album" {
		results, err = client.Search(ctx, qs, spotify.SearchTypeAlbum)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Albums:")
		for _, item := range results.Albums.Albums {
			fmt.Println("    ", item.Name)
		}
	} else {
		results, err = client.Search(ctx, qs, spotify.SearchTypeTrack)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Tracks")
		for _, item := range results.Tracks.Tracks {
			fmt.Println("    ", item.Name)
		}

	}
	return nil
}

func listArtist(client spotify.Client, qs string) {
	ctx := context.Background()
	results, err := client.Search(ctx, qs, spotify.SearchTypeArtist)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Artists:")
	for _, item := range results.Artists.Artists {
		fmt.Println("    ", item.Name)
	}

}
