package day19

import (
	"bytes"
	"log"
	"os"
	"strings"
)

var logger = log.Default()

func Part1() {
	towels, patterns := parse()
	logger.Printf("Day 19, part 1: %d", solve(towels, patterns))
}

func Part2() {
	//towels, patterns := parse()
	//searchDict := buildDict(towels)
	logger.Printf("Day X, part 2: %d", 1)
}

func solve(towels []string, patterns []string) int {
	count := 0
	for _, pattern := range patterns {
		if isPossible(pattern, towels) {
			count++
		}
	}
	return count
}

func isPossible(pattern string, towels []string) bool {
	if pattern == "" {
		return true
	}
	for _, towel := range towels {
		nextPatter := strings.TrimPrefix(pattern, towel)
		if nextPatter != pattern {
			if isPossible(nextPatter, towels) {
				return true
			}
		}
	}
	return false
}

func parse() ([]string, []string) {
	data, _ := os.ReadFile("day19/input.txt")
	parts := bytes.Split(data, []byte("\n\n"))
	var towels []string
	patterns := strings.Split(string(parts[1]), "\n")
	for _, towel := range strings.Split(string(parts[0]), ", ") {
		towels = append(towels, towel)
	}
	return towels, patterns
}
