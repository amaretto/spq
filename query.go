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

var client *spotify.Client

func getClient() error {
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
	client = spotify.New(httpClient)
	return nil
}

func executeQuery(c *cli.Context) error {
	ctx := context.Background()
	if err := getClient(); err != nil {
		return err
	}

	// check args
	var args = c.Args().Slice()
	if len(args) == 0 {
		return fmt.Errorf("no target args specified. see `spq ${kind} -h` for more details")
	} else if len(args) > 1 {
		return fmt.Errorf("too many args specified. see `spq ${kind} -h` for more details")
	}

	qs := args[0]

	if c.Command.Name == "album" {
		results, err := client.Search(ctx, qs, spotify.SearchTypeAlbum)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Albums:")
		for _, item := range results.Albums.Albums {
			fmt.Println("    ", item.Name)
		}
	} else {
		results, err := client.Search(ctx, qs, spotify.SearchTypeTrack)
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

func listArtist(c *cli.Context) error {
	ctx := context.Background()
	if err := getClient(); err != nil {
		return err
	}
	var args = c.Args().Slice()
	//ToDo: check arg length

	qs := args[0]
	results, err := client.Search(ctx, qs, spotify.SearchTypeArtist)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Artists:")
	for _, item := range results.Artists.Artists {
		fmt.Println("    ", item.Name)
	}
	return nil
}

func describeArtist(c *cli.Context) error {
	ctx := context.Background()
	if err := getClient(); err != nil {
		return err
	}
	var args = c.Args().Slice()
	//ToDo: check arg length

	qs := args[0]
	results, err := client.Search(ctx, qs, spotify.SearchTypeArtist)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Searched Artist info:")
	for _, item := range results.Artists.Artists {
		if item.Name == qs {
			fmt.Println("    ID: ", item.ID)
			fmt.Println("    Name: ", item.Name)
			fmt.Println("    Genre: ", item.Genres)
		}
	}
	return nil
}

func listAlbum(c *cli.Context) error {
	ctx := context.Background()
	if err := getClient(); err != nil {
		return err
	}
	var args = c.Args().Slice()
	//ToDo: check arg length

	qs := args[0]
	results, err := client.Search(ctx, qs, spotify.SearchTypeAlbum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Albums:")
	for _, item := range results.Albums.Albums {
		fmt.Println("    ", item.Name)
	}
	return nil
}

func describeAlbum(c *cli.Context) error {
	ctx := context.Background()
	if err := getClient(); err != nil {
		return err
	}
	var args = c.Args().Slice()
	//ToDo: check arg length

	qs := args[0]
	results, err := client.Search(ctx, qs, spotify.SearchTypeAlbum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Searched Album info:")
	for _, item := range results.Albums.Albums {
		if item.Name == qs {
			fmt.Println("    ID: ", item.ID)
			fmt.Println("    Name: ", item.Name)
			fmt.Println("    Artists: ", item.Artists)
			fmt.Println("    ReleaseDate: ", item.ReleaseDate)
		}
	}
	return nil
}
