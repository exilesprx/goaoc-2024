package daysix

import (
	"testing"
)

func TestExample(t *testing.T) {
	result, err := Puzzle("example.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if result != 6 {
		t.Fatalf("Expected 6, got %v", result)
	}
}

func TestPuzzle(t *testing.T) {
  t.Skipf("Skipping test. Takes ~9 seconds to run")
	result, err := Puzzle("puzzle.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Skipf("Result: %v", result)
}
