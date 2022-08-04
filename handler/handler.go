package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fligth-tracker/flighttracker"
)

type Handler struct {
	tracker flighttracker.FlightTracker
}

func NewHandler(t flighttracker.FlightTracker) *Handler {
	return &Handler{
		tracker: t,
	}
}

func (h *Handler) TrackFlight(w http.ResponseWriter, r *http.Request) {
	// parse body
	var paths []flighttracker.Path
	err := json.NewDecoder(r.Body).Decode(&paths)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error on try to read request body"))
		return
	}

	sortedPath := h.tracker.TrackFlights(paths)
	resp, err := json.Marshal(sortedPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error on response data"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
