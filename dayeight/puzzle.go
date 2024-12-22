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

	// grid := [][]rune{}
	frequencyMap := make(map[rune]map[string]Point)
	antinodes := map[string]Point{}
	rows := 0
	columns := 0
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		// grid = append(grid, []rune{})
		for x, r := range line {
			// grid[y] = append(grid[y], r)
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

	for _, points := range frequencyMap {
		for _, pointA := range points {
			for _, pointB := range points {
				distanceX := pointA.x - pointB.x
				distanceY := pointA.y - pointB.y

				x := pointA.x + distanceX
				y := pointA.y + distanceY

				if distanceX == 0 && distanceY == 0 {
					key := key(x, y)
					antinodes[key] = Point{x, y}
					continue
				}
				for x >= 0 && y >= 0 && x < columns && y < rows {
					key := key(x, y)
					// antinodes cannot overlap same frequency nodes - part one
					// if _, ok := frequencyMap[r][key]; ok {
					// 	continue
					// }

					// find unique antinodes
					antinodes[key] = Point{x, y}
					x += distanceX
					y += distanceY
				}
			}
		}
	}
	// printGrid(grid, antinodes)
	return len(antinodes), nil
}

func key(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func printGrid(grid [][]rune, antinodes map[string]Point) {
	for y, row := range grid {
		for x, column := range row {
			_, ok := antinodes[key(x, y)]
			if ok && column != '.' {
				fmt.Printf("*")
			} else if ok {
				fmt.Printf("#")
			} else {
				fmt.Printf("%v", string(column))
			}
		}
		fmt.Println()
	}
}
