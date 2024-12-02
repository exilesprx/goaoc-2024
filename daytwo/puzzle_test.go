package daytwo

import (
	"log"
	"testing"
)

func TestExampleInput(t *testing.T) {
	result, _ := Puzzle("example.txt")
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}

func TestPuzzleInput(t *testing.T) {
	result, err := Puzzle("puzzle.txt")
  if err != nil {
    t.Fatal(err)
  }

	log.Print("Puzzle result: ", result)
}
