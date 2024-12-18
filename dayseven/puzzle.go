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
		target := parseInt(parts[0])
		res := parseInt(numbers[0])
		if evaluate(numbers, target, 1, res) {
			sum += target
		}
	}

	return sum, nil
}

func evaluate(numbers []string, target int, idx int, result int) bool {
	if idx == len(numbers) {
		return result == target
	}

	n := parseInt(numbers[idx])
	if evaluate(numbers, target, idx+1, result+n) {
		return true
	}

	if evaluate(numbers, target, idx+1, result*n) {
		return true
	}

	concat := concatenate(result, n)
	return evaluate(numbers, target, idx+1, concat)
}

func concatenate(result, number int) int {
	factor := 1
	for number/factor > 0 {
		factor *= 10
	}
	return result*factor + number
}

func parseInt(number string) int {
	num, _ := strconv.Atoi(number)
	return num
}
