package main

import "testing"

func TestNewDeck(t *testing.T) {
	expected := 16
	d := newDeck()
	if len(d) != expected {
		t.Errorf("We were expecting %v and we got %v", expected, len(d))
	}
}
