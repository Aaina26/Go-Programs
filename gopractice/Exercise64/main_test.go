package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(5, 5)
	want := 1
	if got != want {
		t.Errorf("Error in Add")
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(6, 6)
	want := 0
	if got != want {
		t.Errorf("Error in Subtract")
	}
}

func TestCal(t *testing.T) {
	got := cal(6, 6, add)
	want := 1
	if got != want {
		t.Errorf("Error in cal")
	}
}
