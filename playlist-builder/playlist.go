package playlist

import (
	"context"
	"fmt"
	"log"

	spotify "github.com/zmb3/spotify/v2"
)

func BuildPlaylist(client *spotify.Client, ctx context.Context, uriList []spotify.ID, trackName, artistName string) {
	user, err := client.CurrentUser(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Logged in!")
	fmt.Println("Getting songs...")
	fmt.Println("Creating playlist...")

	descStr := fmt.Sprintf("Live versions of %s - %s", artistName, trackName)

	newPlaylist, err := client.CreatePlaylistForUser(ctx, user.ID, descStr, "List of live performances", true, false)
	if err != nil {
		log.Fatal(err)
	}
	newPlaylistID := newPlaylist.SimplePlaylist.ID

	version, err := client.AddTracksToPlaylist(ctx, newPlaylistID, uriList...)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Succesfully created playlist!")
		fmt.Println(version)
	}
}
