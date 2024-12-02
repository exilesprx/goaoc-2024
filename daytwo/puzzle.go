package daytwo

import (
	"bufio"
	"errors"
	"log"
	"os"
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
		increments := true
		previous := 0
		numbers := strings.Split(scanner.Text(), " ")
		for _, r := range numbers {
			number, err := strconv.Atoi(r)
			if err != nil {
				return 0, errors.New("Invalid number")
			}

			if previous == 0 {
				previous = number
				next, err := strconv.Atoi(numbers[1])
				if err != nil {
					return 0, errors.New("Invalid number")
				}

				if next > previous {
					increments = true
				} else {
					increments = false
				}
				continue
			}

			if !validDistance(previous - number) {
				valid = false
			}

			if number < previous && increments {
				valid = false
			}

			if number > previous && !increments {
				valid = false
			}

			previous = number
		}

		if valid {
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
