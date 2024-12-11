package dayX

import (
	"bufio"
	"log"
	"os"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day X, part 1: %d", solve())
}

func Part2() {
	logger.Printf("Day X, part 2: %d", solve())
}

func solve() int {
	file, _ := os.Open("dayX/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		_ = scanner.Text()

	}
	return result
}
