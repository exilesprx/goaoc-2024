package dayfour

import (
	"bufio"
	"os"
)

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
    // TODO: read each line of the file
	}

	return 0, nil
}
