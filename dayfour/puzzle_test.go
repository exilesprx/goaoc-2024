package dayfour

import (
	"testing"
)

func TestExampleInput(t *testing.T) {
	result, err := Puzzle("example.txt")
	if err != nil {
		t.Fatalf("Puzzle errored: %v", err)
	}

	if result != 2 {
		t.Fatalf("Expected 2, got %v", result)
	}
}

func TestPuzzleInput(t *testing.T) {
	result, err := Puzzle("puzzle.txt")
	if err != nil {
		t.Fatalf("Puzzle errored: %v", err)
	}

	t.Skipf("Puzzle result: %v", result)
}
