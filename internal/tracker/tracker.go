package tracker

import "tidal-discord/internal/media"

type Tracker struct {
	provider media.Provider
	previous *media.Track
}

func New(provider media.Provider) *Tracker {
	return &Tracker{
		provider: provider,
	}
}

func (t *Tracker) Update() (*media.Track, bool, error) {
	current, err := t.provider.CurrentTrack()
	if err != nil {
		return nil, false, err
	}

	if current == nil {
		t.previous = nil
		return nil, false, nil
	}

	changed := t.previous == nil ||
		t.previous.Title != current.Title ||
		t.previous.Artist != current.Artist ||
		t.previous.Album != current.Album ||
		t.previous.Playing != current.Playing
	t.previous = current

	return current, changed, nil
}
