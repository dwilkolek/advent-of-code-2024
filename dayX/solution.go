package dayX

import (
	"bufio"
	"log"
	"os"
)

var logger = log.Default()

func Part1() {

	logger.Printf("Day X, part 1: %d", 0)
}

func Part2() {

	logger.Printf("Day X, part 2: %d", 0)
}

func solve() {
	file, _ := os.Open("dayX/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_ = scanner.Text()

	}
}
