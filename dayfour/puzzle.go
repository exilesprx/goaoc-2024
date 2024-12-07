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

	grid := [][]rune{}
	scanner := bufio.NewScanner(file)
	y := 0
	for ; scanner.Scan(); y += 1 {
		line := scanner.Text()
		grid = append(grid, []rune{})
		for _, r := range line {
			grid[y] = append(grid[y], r)
		}
	}

	y = 0
	count := 0
	for ; y < len(grid); y += 1 {
		for x := 0; x < len(grid[y]); x += 1 {
			if grid[y][x] == 'A' {
				if backslash(grid, y, x) && forwardslash(grid, y, x) {
					count += 1
				}
			}
		}
	}

	return count, nil
}

func forwardslash(grid [][]rune, y, x int) bool {
	if y-1 < 0 || x+1 >= len(grid[y]) || y+1 >= len(grid) || x-1 < 0 {
		return false
	}
	word := string(grid[y-1][x+1]) + string(grid[y][x]) + string(grid[y+1][x-1])
	return word == "MAS" || word == "SAM"
}

func backslash(grid [][]rune, y, x int) bool {
	if y-1 < 0 || x-1 < 0 || y+1 >= len(grid) || x+1 >= len(grid[y]) {
		return false
	}
	word := string(grid[y-1][x-1]) + string(grid[y][x]) + string(grid[y+1][x+1])
	return word == "MAS" || word == "SAM"
}
