package daynine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, nil
	}

	scanner := bufio.NewScanner(file)
	blocks := []int{}
	for scanner.Scan() {
		diskmap := scanner.Text()
		fileId := 0
		for i, r := range diskmap {
			n, _ := strconv.Atoi(string(r))
			// file
			if i%2 == 0 {
				for i := 1; i <= n; i++ {
					blocks = append(blocks, fileId)
				}
				fileId++
				continue
			}

			// free space
			for i := 1; i <= n; i++ {
				blocks = append(blocks, -1)
			}
		}
	}

	// printBlocks(blocks)
	// fmt.Println()

	for i := 0; i < len(blocks); {
		if blocks[i] == -1 {
			replacement := blocks[len(blocks)-1]
			if replacement == -1 {
				blocks = blocks[:len(blocks)-1]
				continue
			}
			blocks[i] = replacement
			blocks = blocks[:len(blocks)-1]
		}
		i++
	}

  sum := 0
  for i, n := range blocks {
    sum += i * n
  }

	// printBlocks(blocks)
	return sum, nil
}

func printBlocks(blocks []int) {
	for _, n := range blocks {
		if n == -1 {
			fmt.Print(".")
			continue
		}
		fmt.Printf("%v", n)
	}
}
