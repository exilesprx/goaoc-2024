package dayone

import (
	"log"
	"testing"
)

func TestExampleInput(t *testing.T) {
	result := Puzzle("example.txt")
	if result != 11 {
		t.Errorf("Expected 5, but got %d", result)
	}
}

func TestPuzzleInput(t *testing.T) {
	result := Puzzle("puzzle.txt")
	log.Print("Puzzle result: ", result)
}
