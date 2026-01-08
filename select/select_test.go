package myselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 20)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowServerURL := slowServer.URL
	fastServerURL := fastServer.URL

	want := fastServerURL
	got := Racer(slowServerURL, fastServerURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}
