package dayone

import (
	"log"
	"testing"
)

func TestExampleInput(t *testing.T) {
	result := Puzzle("example.txt")
	if result != 31 {
		t.Errorf("Expected 31, but got %d", result)
	}
}

func TestPuzzleInput(t *testing.T) {
	result := Puzzle("puzzle.txt")
	log.Print("Puzzle result: ", result)
}
