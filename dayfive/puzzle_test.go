package dayfive

import (
	"testing"
)

func TestExample(t *testing.T) {
	result, err := Puzzle("example.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if result != 2 {
		t.Fatalf("Expected 2, got %v", result)
	}
}

func TestPuzzle(t *testing.T) {
	result, err := Puzzle("puzzle.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Skip("Result: ", result)
}