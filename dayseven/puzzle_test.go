package dayseven

import (
	"testing"
)

func TestExample(t *testing.T) {
	result, err := Puzzle("example.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if result != 11387 {
		t.Fatalf("Expected 11387, got %v", result)
	}
}

func TestPuzzle(t *testing.T) {
	result, err := Puzzle("puzzle.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Skipf("Result: %v", result)
}
