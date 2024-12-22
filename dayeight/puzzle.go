package dayeight

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

type Point struct {
	x int
	y int
}

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}

	frequencyMap := make(map[rune]map[string]Point)
	antinodes := map[string]Point{}
	rows := 0
	columns := 0
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, r := range line {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				if _, exists := frequencyMap[r]; !exists {
					frequencyMap[r] = map[string]Point{}
				}
				frequencyMap[r][key(x, y)] = Point{x, y}
			}
		}
		if columns == 0 {
			columns = utf8.RuneCountInString(line)
		}
		rows++
	}

	for r, points := range frequencyMap {
		for _, pointA := range points {
			for _, pointB := range points {
				x := pointA.x + (pointA.x - pointB.x)
				y := pointA.y + (pointA.y - pointB.y)

				if x < 0 || y < 0 || x >= columns || y >= rows {
					continue
				}

				key := key(x, y)
				// antinodes cannot overlap same frequency nodes
				if _, ok := frequencyMap[r][key]; ok {
					continue
				}

				// find unique antinodes
				antinodes[key] = Point{x, y}
			}
		}
	}
	return len(antinodes), nil
}

func key(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}
