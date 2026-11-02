package media

import "time"

type Track struct {
	Title    string
	Artist   string
	Album    string
	Duration time.Duration
	Position time.Duration
	Playing  bool
}

type Provider interface {
	CurrentTrack() (*Track, error)
}
