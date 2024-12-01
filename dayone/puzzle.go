package dayone

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Puzzle(name string) int {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var leftNumbers []int = nil
	var rightNumbers []int = nil
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    left, right := parseSides(scanner)
		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	sum := 0
	for i := 0; i < len(leftNumbers); i++ {
    number := leftNumbers[i]
    sum += number * appearences(number, rightNumbers)
	}

	return sum
}

func parseSides(scanner *bufio.Scanner) (int, int) {
	sides := strings.Fields(scanner.Text())
	left, err := strconv.Atoi(sides[0])
	if err != nil {
		log.Fatal("Error converting string to int")
	}
	right, err := strconv.Atoi(sides[1])
	if err != nil {
		log.Fatal("Error converting string to int")
	}

	return left, right
}

func appearences(number int, numbers []int) int {
  count := 0
  for _, n := range numbers {
    if n == number {
      count += 1
    }
  }

  return count
}
