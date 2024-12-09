package daysix

import (
	"bufio"
	"os"
)

type Guard struct {
	position  []int
	direction rune
}

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	guard := Guard{
		position:  []int{0, 0},
		direction: '^',
	}
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		grid = append(grid, []rune{})
		for x, r := range line {
			grid[y] = append(grid[y], r)
			if r == '^' || r == 'v' || r == '<' || r == '>' {
				guard.direction = r
				guard.position = []int{y, x}
			}
		}
	}

	offgrid := false
	for !offgrid {
		switch guard.direction {
		case '^':
			offgrid = movenorth(grid, &guard, guard.position[0], guard.position[1])
		case 'v':
			offgrid = movesouth(grid, &guard, guard.position[0], guard.position[1])
		case '>':
			offgrid = moveeast(grid, &guard, guard.position[0], guard.position[1])
		case '<':
			offgrid = movewest(grid, &guard, guard.position[0], guard.position[1])
		}
	}

	return calculateDistinctVisits(grid), nil
}

// Movenorth moves the guard north until it hits a wall or the edge of the grid.
// It returns true if the guard is off the grid
func movenorth(grid [][]rune, guard *Guard, y, x int) bool {
	for ; y >= 0; y-- {
		if grid[y][x] == '#' {
			guard.direction = '>'
			guard.position = []int{y + 1, x}
			return false
		}
		grid[y][x] = 'X'
	}

	return true
}

// Movesouth moves the guard south until it hits a wall or the edge of the grid.
// It returns true if the guard is off the grid
func movesouth(grid [][]rune, guard *Guard, y, x int) bool {
	for ; y < len(grid); y++ {
		if grid[y][x] == '#' {
			guard.direction = '<'
			guard.position = []int{y - 1, x}
			return false
		}
		grid[y][x] = 'X'
	}

	return true
}

// Movewest moves the guard west until it hits a wall or the edge of the grid.
// It returns true if the guard is off the grid
func movewest(grid [][]rune, guard *Guard, y, x int) bool {
	for ; x >= 0; x-- {
		if grid[y][x] == '#' {
			guard.direction = '^'
			guard.position = []int{y, x + 1}
			return false
		}
		grid[y][x] = 'X'
	}

	return true
}

// Moveeast moves the guard east until it hits a wall or the edge of the grid.
// It returns true if the guard is off the grid
func moveeast(grid [][]rune, guard *Guard, y, x int) bool {
	for ; x < len(grid[y]); x++ {
		if grid[y][x] == '#' {
			guard.direction = 'v'
			guard.position = []int{y, x - 1}
			return false
		}
		grid[y][x] = 'X'
	}

	return true
}

// CalculateDistinctVisits counts the number of distinct visits in the grid.
func calculateDistinctVisits(grid [][]rune) int {
	visits := 0
	for _, row := range grid {
		for _, r := range row {
			if r == 'X' {
				visits += 1
			}
		}
	}

	return visits
}
