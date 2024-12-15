package dayseven

import (
	"testing"
)

func TestExample(t *testing.T) {
	result, err := Puzzle("example.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if result != 3749 {
		t.Fatalf("Expected 3749, got %v", result)
	}
}

func TestPuzzle(t *testing.T) {
	t.Skipf("Skipping test")
	result, err := Puzzle("puzzle.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Skipf("Result: %v", result)
}
