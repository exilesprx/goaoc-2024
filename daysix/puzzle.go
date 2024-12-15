package daysix

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Guard struct {
	obstructions map[string]int
	position     []int
	direction    rune
}

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	var startingLocation []int
	var startingDirection rune
	guard := Guard{
		position:     []int{0, 0},
		direction:    '^',
		obstructions: make(map[string]int),
	}
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		grid = append(grid, []rune{})
		for x, r := range line {
			grid[y] = append(grid[y], r)
			if r == '^' || r == 'v' || r == '<' || r == '>' {
				startingDirection = r
				startingLocation = []int{y, x}
			}
		}
	}

	loopsDetected := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			offgrid := false
			guard.direction = startingDirection
			guard.position = startingLocation
			if y == guard.position[0] && x == guard.position[1] || grid[y][x] == '#' {
				continue
			}
			grid[y][x] = '0'
		outer:
			for !offgrid {
				switch guard.direction {
				case '^':
					offgrid, err = movenorth(grid, &guard, guard.position[0], guard.position[1])
					if err != nil {
						loopsDetected += 1
						break outer
					}
				case 'v':
					offgrid, err = movesouth(grid, &guard, guard.position[0], guard.position[1])
					if err != nil {
						loopsDetected += 1
						break outer
					}
				case '>':
					offgrid, err = moveeast(grid, &guard, guard.position[0], guard.position[1])
					if err != nil {
						loopsDetected += 1
						break outer
					}
				case '<':
					offgrid, err = movewest(grid, &guard, guard.position[0], guard.position[1])
					if err != nil {
						loopsDetected += 1
						break outer
					}
				}
			}
			grid[y][x] = '.'
			guard.obstructions = make(map[string]int)
		}
	}

	return loopsDetected, nil
}

// movenorth moves the guard north until it hits a wall or the edge of the grid.
// It returns true if the guard is off the grid
func movenorth(grid [][]rune, guard *Guard, y, x int) (bool, error) {
	for ; y >= 0; y-- {
		key := strconv.Itoa(y) + ":" + strconv.Itoa(x)
		if grid[y][x] == '#' || grid[y][x] == '0' {
			guard.direction = '>'
			guard.position = []int{y + 1, x}

			guard.obstructions[key] += 1
			if guard.obstructions[key] >= 4 {
				return false, errors.New("loop detected")
			}

			return false, nil
		}

		grid[y][x] = 'X'
	}

	return true, nil
}

// movesouth moves the guard south until it hits a wall or the edge of the grid.
// It returns true if the guard is off the grid
func movesouth(grid [][]rune, guard *Guard, y, x int) (bool, error) {
	for ; y < len(grid); y++ {
		key := strconv.Itoa(y) + ":" + strconv.Itoa(x)
		if grid[y][x] == '#' || grid[y][x] == '0' {
			guard.direction = '<'
			guard.position = []int{y - 1, x}

			guard.obstructions[key] += 1
			if guard.obstructions[key] >= 4 {
				return false, errors.New("loop detected")
			}
			return false, nil
		}
		grid[y][x] = 'X'
	}

	return true, nil
}

// movewest moves the guard west until it hits a wall or the edge of the grid.
// It returns true if the guard is off the grid
func movewest(grid [][]rune, guard *Guard, y, x int) (bool, error) {
	for ; x >= 0; x-- {
		key := strconv.Itoa(y) + ":" + strconv.Itoa(x)
		if grid[y][x] == '#' || grid[y][x] == '0' {
			guard.direction = '^'
			guard.position = []int{y, x + 1}

			guard.obstructions[key] += 1
			if guard.obstructions[key] >= 4 {
				return false, errors.New("loop detected")
			}
			return false, nil
		}
		grid[y][x] = 'X'
	}

	return true, nil
}

// moveeast moves the guard east until it hits a wall or the edge of the grid.
// It returns true if the guard is off the grid
func moveeast(grid [][]rune, guard *Guard, y, x int) (bool, error) {
	for ; x < len(grid[y]); x++ {
		key := strconv.Itoa(y) + ":" + strconv.Itoa(x)
		if grid[y][x] == '#' || grid[y][x] == '0' {
			guard.direction = 'v'
			guard.position = []int{y, x - 1}

			guard.obstructions[key] += 1
			if guard.obstructions[key] >= 4 {
				return false, errors.New("loop detected")
			}
			return false, nil
		}
		grid[y][x] = 'X'
	}

	return true, nil
}

// calculateDistinctVisits counts the number of distinct visits in the grid.
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

// printGrid prints the grid to the console.
func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

// printObstructions prints the obstructions to the console.
func printObstructions(guard Guard) {
	for key, count := range guard.obstructions {
		fmt.Printf("Positions %v and Count %v", key, count)
		fmt.Println()
	}
}
