package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello when name is empty.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello to people.", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello, John"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello to people in spanish.", func(t *testing.T) {
		got := Hello("John", spanish)
		want := "Hola, John"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello to people in french.", func(t *testing.T) {
		got := Hello("John", french)
		want := "Bonjour, John"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello to people in gaelic.", func(t *testing.T) {
		got := Hello("John", gaelic)
		want := "Dia dhuit, John"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
