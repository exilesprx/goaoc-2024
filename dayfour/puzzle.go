package dayfour

import (
	"bufio"
	"os"
	"sync"
)

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}

	grid := [][]rune{}
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y += 1 {
		line := scanner.Text()
		grid = append(grid, []rune{})
		for _, r := range line {
			grid[y] = append(grid[y], r)
		}
	}

	xmas := make(chan bool)
	wg := sync.WaitGroup{}
	for y := 0; y < len(grid); y += 1 {
		for x := 0; x < len(grid[y]); x += 1 {
			if grid[y][x] == 'A' {
				wg.Add(1)
				go calculate(grid, &wg, xmas, y, x)
			}
		}
	}

	go waitOnCalculate(&wg, xmas)

	return receive(xmas), nil
}

func calculate(grid [][]rune, wg *sync.WaitGroup, xmas chan<- bool, y, x int) {
	defer wg.Done()
	xmas <- forwardslash(grid, y, x) && backslash(grid, y, x)
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

func waitOnCalculate(wg *sync.WaitGroup, xmas chan<- bool) {
	wg.Wait()
	close(xmas)
}

func receive(xmas <-chan bool) int {
	count := 0
	for found := range xmas {
		if found {
			count += 1
		}
	}
	return count
}
