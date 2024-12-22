package dayeight

import (
	"testing"
)

func TestExample(t *testing.T) {
	result, err := Puzzle("example.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if result != 34 {
		t.Fatalf("Expected 34, got %v", result)
	}
}

func TestPuzzle(t *testing.T) {
	result, err := Puzzle("puzzle.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Skipf("Result: %v", result)
}
