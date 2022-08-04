package flighttracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TrackFlights(t *testing.T) {
	tests := []struct {
		name  string
		want  []string
		paths []Path
	}{
		{
			name: "should return the correct list of flights",
			want: []string{"SFO", "EWR", "ATL", "ATL"},
			paths: []Path{
				{"ATL", "EWR"},
				{"SFO", "ATL"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tracker{}
			p := tr.TrackFlights(tt.paths)
			assert.Equal(t, p, tt.want)
		})
	}
}
