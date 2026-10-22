//go:build windows

package main

import (
	"fmt"
	"runtime"
	"time"

	"tidal-discord/internal/discord"
	"tidal-discord/internal/media"
	"tidal-discord/internal/tracker"

	"github.com/go-ole/go-ole"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED); err != nil {
		fmt.Println("Error initializing COM:", err)
		return
	}
	defer ole.CoUninitialize()

	provider := media.WindowsProvider{}
	t := tracker.New(provider)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	const appID = "1557046038450405537"

	if err := discord.New(appID); err != nil {
		fmt.Println("Discord error:", err)
		return
	}
	defer discord.Close()

	for {
		track, changed, err := t.Update()

		if err != nil {
			fmt.Println("Error:", err)
		} else if track != nil {
			fmt.Printf(
				"%s — %s | position=%v | playing=%v\n",
				track.Artist,
				track.Title,
				track.Position,
				track.Playing,
			)

			if changed {
				if err := discord.SetTrack(track); err != nil {
					fmt.Println("Discord error:", err)
				}

				if track.Playing {
					fmt.Printf("Playback started or track changed: %s — %s\n",
						track.Artist,
						track.Title,
					)
				} else {
					fmt.Printf("Playback paused: %s — %s\n",
						track.Artist,
						track.Title,
					)
				}
			}
		}

		<-ticker.C
	}
}
