package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Bambank")
	want := "Hello, Bambank"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
