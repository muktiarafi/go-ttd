package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to someone", func(t *testing.T) {
		got := Hello("Bambank")
		want := "Hello, Bambank"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say `Hello, World' when passing empty string argument", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
