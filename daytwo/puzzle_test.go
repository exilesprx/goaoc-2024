package daytwo

import (
	"log"
	"testing"
)

func TestExampleInput(t *testing.T) {
	result, _ := Puzzle("example.txt")
	if result != 31 {
		t.Errorf("Expected 31, but got %d", result)
	}
}

func TestPuzzleInput(t *testing.T) {
	result, err := Puzzle("puzzle.txt")
  if err != nil {
    t.Fatal(err)
  }

	log.Print("Puzzle result: ", result)
}
