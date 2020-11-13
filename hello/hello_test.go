package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to someone", func(t *testing.T) {
		got := Hello("Bambank", "")
		want := "Hello, Bambank"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say `Hello, World' when passing empty string argument", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in japanese", func(t *testing.T) {
		got := Hello("Udin", "Japanese")
		want := "Konnichiwa, Udin"
		assertCorrectMessage(t, got, want)

	})
}
