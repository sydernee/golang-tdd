package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Andrea")
	want := "Hello, Andrea"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
