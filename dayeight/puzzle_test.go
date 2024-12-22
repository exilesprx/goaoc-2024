package dayeight

import (
	"testing"
)

func TestExample(t *testing.T) {
	result, err := Puzzle("example.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if result != 14 {
		t.Fatalf("Expected 14, got %v", result)
	}
}

func TestPuzzle(t *testing.T) {
  // 254 is too high
	result, err := Puzzle("puzzle.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Skipf("Result: %v", result)
}
