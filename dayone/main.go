package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var leftNumbers []int = nil
	var rightNumbers []int = nil
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sides := strings.Fields(scanner.Text())
		left, err := strconv.Atoi(sides[0])
		if err != nil {
			log.Fatal("Error converting string to int")
		}
    right, err := strconv.Atoi(sides[1])
    if err != nil {
      log.Fatal("Error converting string to int")
    }
		leftNumbers = append(leftNumbers, left)
    rightNumbers = append(rightNumbers, right)
	}

  sort.Ints(leftNumbers)
  sort.Ints(rightNumbers)

  sum := 0
  for i := 0; i <len(leftNumbers); i ++ {
    distance := leftNumbers[i] - rightNumbers[i]
    if distance < 0 {
      sum += distance * -1
    } else {
      sum += distance
    }
  }

  log.Print("Sum: ", sum)
}
