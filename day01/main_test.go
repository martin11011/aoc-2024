package main

import (
	_ "embed"
	"testing"
)

//go:embed testinput.txt
var testInput string

func Test_part1(t *testing.T) {
	result, err := part1(testInput)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result != 11 {
		t.Errorf("Sum was incorrect, got: %d, want: %d", result, 11)
	}
}
