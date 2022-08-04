package main

import (
	"net/http"

	"github.com/fligth-tracker/flighttracker"
	"github.com/fligth-tracker/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	h := handler.NewHandler(flighttracker.NewTracker())
	r.Post("/track", h.TrackFlight)
	http.ListenAndServe(":8080", r)
}
