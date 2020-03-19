package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello world when no string passed in", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("amigo", "Spanish")
		want := "Hola, amigo"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("mon amis", "French")
		want := "Bonjour, mon amis"
		assertCorrectMessage(t, got, want)

	})
}
