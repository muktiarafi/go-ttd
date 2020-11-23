package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://facebook.com"
	fastURL := "http://anjay.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
