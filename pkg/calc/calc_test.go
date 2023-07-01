package calc

import (
	"testing"
)

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
