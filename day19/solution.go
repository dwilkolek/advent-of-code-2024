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
	count := 0
	for _, pattern := range patterns {
		if isPossible(pattern, towels) > 0 {
			count++
		}
	}
	logger.Printf("Day 19, part 1: %d", count)
}

func Part2() {
	towels, patterns := parse()
	count := 0
	for _, pattern := range patterns {
		count += isPossible(pattern, towels)
	}
	logger.Printf("Day 19, part 2: %d", count)
}

var cache = map[string]int{}

func isPossible(pattern string, towels []string) int {
	cacheValue, ok := cache[pattern]
	if ok {
		return cacheValue
	}
	if pattern == "" {
		return 1
	}
	sum := 0
	for _, towel := range towels {
		nextPatter := strings.TrimPrefix(pattern, towel)
		if nextPatter != pattern {
			sum += isPossible(nextPatter, towels)
		}
	}
	cache[pattern] = sum
	return sum
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
