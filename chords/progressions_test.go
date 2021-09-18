package chords

import (
	"github.com/nate-droid/core/notes"
	"github.com/nate-droid/core/scales"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgression(t *testing.T) {
	type args struct {
		mode      scales.ModeName
		intervals []ChordInterval
		key       notes.Note
	}
	tests := []struct {
		name string
		args args
		want []Chord
	}{
		{
			name: "C Major I - IV -V",
			args: args{
				mode:      scales.Ionian,
				intervals: []ChordInterval{I, IV, V},
				key:       notes.C,
			},
			want: []Chord{
				{Root: notes.C.Name},
				{Root: notes.F.Name},
				{Root: notes.G.Name},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			progression, err := Progression(tt.args.mode, tt.args.intervals, tt.args.key)
			assert.NoError(t, err)
			assert.Equal(t, tt.want[0].Name, progression[0].Name)
			assert.Equal(t, tt.want[1].Name, progression[1].Name)
			assert.Equal(t, tt.want[2].Name, progression[2].Name)
		})
	}
}
