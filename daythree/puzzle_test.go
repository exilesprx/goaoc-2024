package daythree

import (
	"fmt"
	"testing"
)

func TestExampleInput(t *testing.T) {
	result, err := Puzzle("example.txt")
	if err != nil {
		t.Fatal(err)
	}

  if result != 1 {
    t.Error("Expected 1, but got", result)
  }
}

func TestPuzzleInput(t *testing.T) {
  result, err := Puzzle("puzzle.txt")
  if err != nil {
    t.Fatal(err)
  }

  fmt.Println("Puzzle result: ", result)
}
