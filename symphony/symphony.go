package symphony

import (
	"github.com/nate-droid/core/chords"
	"github.com/nate-droid/core/notes"
	"github.com/nate-droid/core/scales"
)

type Symphony struct {
	SongStructure
	Sections []Section
	ID       string
}

type Section struct {
	Type      string
	GroupSize int
}

type SongStructure struct {
	Key              notes.Note
	ChordProgression []chords.ChordInterval
	Mode             scales.ModeName
	SymphonyID       string
	Section
}

type Song struct {
	ChordProgression []chords.Chord
	Section
	SymphonyID string
}
