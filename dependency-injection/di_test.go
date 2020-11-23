package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Bambank")

	got := buffer.String()
	want := "Hello, Bambank"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
