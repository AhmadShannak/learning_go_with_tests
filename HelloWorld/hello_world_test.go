package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "Hello World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("Saying hi to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hi to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("Saying hi in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("Saying hi in frenc", func(t *testing.T) {
		got := Hello("Abass", "french")
		want := "Bonjour, Abass"

		AssertCorrectMessage(t, got, want)
	})
}

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // this will prevent showing this helper as the problem when test fails, and instead shows the one that called it

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
