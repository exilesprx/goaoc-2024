package daythree

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`mul\(\d+,\d+\)`)
		matches := re.FindAllString(line, -1)
		if matches == nil {
			continue
		}

		for _, match := range matches {
			re := regexp.MustCompile(`\d+`)
			numbers := re.FindAllString(match, -1)
			sum += stringToInt(numbers[0]) * stringToInt(numbers[1])
		}
	}

	return sum, nil
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
