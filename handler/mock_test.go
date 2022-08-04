package handler

import "github.com/fligth-tracker/flighttracker"

type mockTracker struct {
	trackFlights func(paths []flighttracker.Path) []string
}

func (m *mockTracker) TrackFlights(paths []flighttracker.Path) []string {
	if m.trackFlights != nil {
		return m.trackFlights(paths)
	}
	return []string{}
}
