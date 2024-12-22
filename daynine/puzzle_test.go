package daynine

import "testing"

func TestExample(t *testing.T) {
	result, err := Puzzle("example.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if result != 1928 {
		t.Fatalf("Expected 1928, got %v", result)
	}
}

func TestPuzzle(t *testing.T) {
	t.Skip("Skipping test")
	result, err := Puzzle("puzzle.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Skipf("Result: %v", result)
}
