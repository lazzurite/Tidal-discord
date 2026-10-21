package discord

import (
	"fmt"
	"net/url"
	"tidal-discord/internal/media"
	"time"

	"github.com/hugolgst/rich-go/client"
)

type Client struct{}

func New(appID string) error {
	return client.Login(appID)
}

func SetTrack(track *media.Track) error {
	activity := client.Activity{
		Details:    track.Title,
		State:      track.Artist,
		LargeImage: "tidal_logo",
		LargeText:  "Listening to TIDAL",
	}

	query := url.QueryEscape(
		track.Artist + " " + track.Title + " official audio",
	)

	activity.Buttons = []*client.Button{
		{
			Label: "YouTube",
			Url:   "https://www.youtube.com/results?search_query=" + query,
		},
	}

	if track.Playing {
		start := time.Now().Add(-track.Position)

		activity.Timestamps = &client.Timestamps{
			Start: &start,
		}
	} else {
		activity.State = fmt.Sprintf(
			"%s • Paused at %s",
			track.Artist,
			formatPosition(track.Position),
		)
	}

	return client.SetActivity(activity)
}

func formatPosition(position time.Duration) string {
	seconds := int(position.Seconds())
	return fmt.Sprintf("%d:%02d", seconds/60, seconds%60)
}

func Close() {
	client.Logout()
}
