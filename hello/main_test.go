package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Ali!", "")
		want := "Hello, Ali!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello to world if empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in Italian", func(t *testing.T) {
		got := Hello("Ali!", "Italian")
		want := "Ciao, Ali!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in French", func(t *testing.T) {
		got := Hello("Ali!", "French")
		want := "Bonjour, Ali!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %q but got %q.", want, got)
	}
}
