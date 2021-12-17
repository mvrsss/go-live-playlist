package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	authorizeUser "github.com/mvrsss/go-live-playlist/auth"
	playlist "github.com/mvrsss/go-live-playlist/playlist-builder"
	search "github.com/mvrsss/go-live-playlist/search-track"
)

var (
	trackName  string
	artistName string
)

func readInput() string {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	return input
}

func main() {
	// Authorize user for Spotify
	client, ctx := authorizeUser.AuthUser()

	fmt.Println("Enter song name: ")
	trackName = readInput()

	fmt.Println("Enter artist name: ")
	artistName = readInput()

	time.Sleep(1 * time.Second)
	uriList := search.SearchTrack(client, ctx, trackName, artistName)

	time.Sleep(1 * time.Second)
	// Build and create playlist on user's account
	playlist.BuildPlaylist(client, ctx, uriList, trackName, artistName)
}
