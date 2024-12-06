package dayfour

import (
	"bufio"
	"fmt"
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
			if grid[y][x] == 'X' {
				if north(grid, x, y) {
          fmt.Println("Found at ", x, y)
					count += 1
				}
				if northeast(grid, x, y) {
          fmt.Println("Found at ", x, y)
					count += 1
				}
				if east(grid, x, y) {
          fmt.Println("Found at ", x, y)
					count += 1
				}
				if southeast(grid, x, y) {
          fmt.Println("Found at ", x, y)
					count += 1
				}
				if south(grid, x, y) {
          fmt.Println("Found at ", x, y)
					count += 1
				}
				if southwest(grid, x, y) {
          fmt.Println("Found at ", x, y)
					count += 1
				}
				if west(grid, x, y) {
          fmt.Println("Found at ", x, y)
					count += 1
				}
				if northwest(grid, x, y) {
          fmt.Println("Found at ", x, y)
					count += 1
				}
			}
		}
	}

	return count, nil
}

func north(grid [][]rune, x, y int) bool {
	if y <= 3 {
		return false
	}
	word := ""
	for i := y - 1; i >= y-3; i -= 1 {
		word += string(grid[i][x])
	}

	return word == "MAS"
}

func south(grid [][]rune, x, y int) bool {
	if y >= len(grid)-4 {
		return false
	}
	word := ""
	for i := y + 1; i <= y+3; i += 1 {
		word += string(grid[i][x])
	}
	return word == "MAS"
}

func east(grid [][]rune, x, y int) bool {
	if x >= len(grid[y])-4 {
		return false
	}
	word := ""
	for i := x + 1; i <= x+3; i += 1 {
		word += string(grid[y][i])
	}
	return word == "MAS"
}

func west(grid [][]rune, x, y int) bool {
	if x <= 3 {
		return false
	}
	word := ""
	for i := x - 1; i >= x-3; i -= 1 {
		word += string(grid[y][i])
	}
	return word == "MAS"
}

func northeast(grid [][]rune, x, y int) bool {
	if y <= 3 || x >= len(grid[y])-4 {
		return false
	}
	word := ""
	for i := 1; i <= 3; i += 1 {
		word += string(grid[y-i][x+i])
	}
	return word == "MAS"
}

func northwest(grid [][]rune, x, y int) bool {
	if y <= 3 || x <= 3 {
		return false
	}
	word := ""
	for i := 1; i <= 3; i += 1 {
		word += string(grid[y-i][x-i])
	}
	return word == "MAS"
}

func southeast(grid [][]rune, x, y int) bool {
	if y >= len(grid)-4 || x >= len(grid[y])-4 {
		return false
	}
	word := ""
	for i := 1; i <= 3; i += 1 {
		word += string(grid[y+i][x+i])
	}
	return word == "MAS"
}

func southwest(grid [][]rune, x, y int) bool {
	if y >= len(grid)-4 || x <= 3 {
		return false
	}
	word := ""
	for i := 1; i <= 3; i += 1 {
		word += string(grid[y+i][x-i])
	}
	return word == "MAS"
}
