package daytwo

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Print("Error opening file")
	}
	defer file.Close()

	safeReports := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		numbers, err := convertToInt(text)
		if err != nil {
			return 0, err
		}

		if validList(numbers) {
			safeReports += 1
		} else {
			valid := false
			for i := 0; i < len(numbers); i++ {
				list := []int{}
				list = append(append(list, numbers[:i]...), numbers[i+1:]...)
				if validList(list) {
					valid = true
					break
				}
			}

			if valid {
				safeReports += 1
			}
		}
	}
	return safeReports, nil
}

func convertToInt(text []string) ([]int, error) {
	numbers := []int{}
	for _, r := range text {
		n, err := strconv.Atoi(r)
		if err != nil {
			return nil, errors.New("Invalid number")
		}
		numbers = append(numbers, n)
	}

	return numbers, nil
}

func validList(numbers []int) bool {
	iod := sort.IsSorted(sort.IntSlice(numbers)) || sort.IsSorted(sort.Reverse(sort.IntSlice(numbers)))
	if !iod {
		return false
	}

	previous := 0
	for _, n := range numbers {
		if previous == 0 {
			previous = n
			continue
		}

		if !validDistance(previous - n) {
			return false
		}

		previous = n
	}

	return true
}

func validDistance(distance int) bool {
	if distance < 0 {
		distance = distance * -1
	}
	return distance >= 1 && distance <= 3
}
