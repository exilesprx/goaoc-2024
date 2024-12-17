package dayseven

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		numbers := strings.Split(parts[1], " ")
		possibleCombos := 1 << (len(numbers) - 1)
		for mask := 0; mask < possibleCombos; mask++ {
			result, _ := strconv.Atoi(numbers[0])
			for i := 0; i < len(numbers)-1; i++ {
				if mask&(1<<i) != 0 {
					n, _ := strconv.Atoi(numbers[i+1])
					result *= n
				} else {
					n, _ := strconv.Atoi(numbers[i+1])
					result += n
				}
			}

			if target, _ := strconv.Atoi(parts[0]); result == target {
				sum += result
				break
			}
		}
	}

	return sum, nil
}
