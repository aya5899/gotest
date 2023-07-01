package main

import "testing"

func TestAddint(t *testing.T) {
	got := Addint(1, 2)
	want := 3
	if got != want {
		t.Errorf("Addint(1, 2) == %d, want %d", got, want)
	}
}

func TestSubint(t *testing.T) {
	got := Subint(1, 2)
	want := -1
	if got != want {
		t.Errorf("Subint(1, 2) == %d, want %d", got, want)
	}
}

func TestMax(t *testing.T) {
	got := Max(1, 2)
	want := 2
	if got != want {
		t.Errorf("Max(1, 2) == %d, want %d", got, want)
	}
}

func TestMin(t *testing.T) {
	got := Min(1, 2)
	want := 1
	if got != want {
		t.Errorf("Min(1, 2) == %d, want %d", got, want)
	}
}
