package flighttracker

import "sort"

type Path []string

type FlightTracker interface {
	TrackFlights(paths []Path) []string
}

type tracker struct{}

func NewTracker() *tracker {
	return &tracker{}
}

func (t *tracker) TrackFlights(paths []Path) []string {
	var fligths []string
	for i := 0; i < len(paths); i++ {
		fligths = append(fligths, paths[i]...)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(fligths)))

	return fligths
}
