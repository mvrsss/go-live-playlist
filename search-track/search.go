package search

import (
	"context"
	"log"
	"strings"

	spotify "github.com/zmb3/spotify/v2"
)

func SearchTrack(client *spotify.Client, ctx context.Context, trackName, artistName string) []spotify.ID {
	var uriList []spotify.ID

	results, err := client.Search(ctx, trackName, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	if results.Tracks != nil {
		for _, item := range results.Tracks.Tracks {
			for _, artist := range item.Artists {
				if artist.Name == artistName && strings.Contains(item.Name, trackName) && (strings.Contains(item.Album.Name, "Live") || strings.Contains(item.Album.Name, "live") || strings.Contains(item.Name, "Live") || strings.Contains(item.Name, "live")) {
					songURI := item.ID
					uriList = append(uriList, songURI)
				}
			}
		}
	}

	return uriList
}
