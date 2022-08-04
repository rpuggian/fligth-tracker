package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fligth-tracker/flighttracker"
	"github.com/stretchr/testify/assert"
)

func Test_TrackFlight(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		tr       flighttracker.FlightTracker
		wantCode int
		wantBody string
	}{
		{
			name: "Should return 200 ok",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/track", strings.NewReader(`[["ATL", "EWR"], ["SFO", "ATL"]]`)),
			},
			tr: &mockTracker{
				trackFlights: func(paths []flighttracker.Path) []string {
					return []string{"SFO", "EWR", "ATL", "ATL"}
				},
			},
			wantCode: 200,
			wantBody: `["SFO","EWR","ATL","ATL"]`,
		},
		{
			name: "Should return 400 bad request",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/track", strings.NewReader(`{"wrong": "body"}`)),
			},
			tr: &mockTracker{
				trackFlights: func(paths []flighttracker.Path) []string {
					return []string{"SFO", "EWR", "ATL", "ATL"}
				},
			},
			wantCode: 400,
			wantBody: `error on try to read request body`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.tr)
			h.TrackFlight(tt.args.w, tt.args.r)
			assert.Equal(t, tt.wantCode, tt.args.w.Code)
			assert.Equal(t, tt.wantBody, tt.args.w.Body.String())
		})
	}
}
