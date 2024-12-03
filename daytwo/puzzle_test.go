package daytwo

import (
	"testing"
)

func TestExampleInput(t *testing.T) {
	result, err := Puzzle("example.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}

func TestPuzzleInput(t *testing.T) {
	result, err := Puzzle("puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Skipf("Puzzle result: %d", result)
}
