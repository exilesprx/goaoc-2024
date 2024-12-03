package daytwo

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	safeReports := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		valid := true
		previous := 0
		text := strings.Split(scanner.Text(), " ")
		numbers := []int{}
		for _, r := range text {
			n, err := strconv.Atoi(r)
			if err != nil {
				return 0, errors.New("Invalid number")
			}
			numbers = append(numbers, n)
		}

		iod := sort.IsSorted(sort.IntSlice(numbers)) || sort.IsSorted(sort.Reverse(sort.IntSlice(numbers)))
		for _, n := range numbers {
			if previous == 0 {
				previous = n
				continue
			}

			if !validDistance(previous - n) {
				valid = false
				break
			}

			previous = n
		}

		if valid && iod {
			safeReports += 1
		}
	}
	return safeReports, nil
}

func validDistance(distance int) bool {
	if distance < 0 {
		distance = distance * -1
	}
	return distance >= 1 && distance <= 3
}
