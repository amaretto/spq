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

	// parse flags
	artist := c.String("artist")
	album := c.String("album")
	track := c.String("track")

	if artist != "" {
		results, err := client.Search(ctx, artist, spotify.SearchTypeArtist|spotify.SearchTypeAlbum|spotify.SearchTypeTrack)
		if err != nil {
			log.Fatal(err)
		}
		if results.Artists != nil {
			fmt.Println("Artists:")
			for _, item := range results.Artists.Artists {
				fmt.Println("    ", item.Name)
			}
		}
	}

	if album != "" {
		results, err := client.Search(ctx, album, spotify.SearchTypeAlbum)
		if err != nil {
			log.Fatal(err)
		}
		if results.Albums != nil {
			fmt.Println("albums:")
			for _, item := range results.Albums.Albums {
				artistString := ""
				for i := 0; i < len(item.Artists); i++ {
					artistString += item.Artists[i].Name
				}
				fmt.Println("artists:", artistString, "\t\t\talbum name:", item.Name)
			}
		}
	}

	if track != "" {
		results, err := client.Search(ctx, track, spotify.SearchTypeTrack)
		if err != nil {
			log.Fatal(err)
		}
		if results.Tracks != nil {
			for _, item := range results.Tracks.Tracks {
				artistString := ""
				albumString := ""
				for i := 0; i < len(item.Artists); i++ {
					artistString += item.Artists[i].Name
					albumString += item.Album.Name
				}
				fmt.Println("artists:", artistString, "\talbum:", albumString, "\ttrack:", item.Name)
			}
		}
	}

	//// handle album results
	//if results.Albums != nil {
	//	fmt.Println("Albums:")
	//	for _, item := range results.Albums.Albums {
	//		fmt.Println("   ", item.Name)
	//	}
	//}
	//// handle playlist results
	//if results.Tracks != nil {
	//	fmt.Println("Tracks:")
	//	for _, item := range results.Tracks.Tracks {
	//		fmt.Println("   ", item.Name)
	//	}
	//}
	return nil
}
