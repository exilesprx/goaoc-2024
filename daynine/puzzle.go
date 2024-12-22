package daynine

import (
	"bufio"
	"fmt"
	"os"
)

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, nil
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range line {
			fmt.Printf(string(r))
		}
	}

	return 0, nil
}
