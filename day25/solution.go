package day25

import (
	"log"
	"os"
	"slices"
	"strings"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day X, part 1: %d", solve())
}

func solve() int {
	file, _ := os.ReadFile("day25/input.txt")

	entries := strings.Split(string(file), "\n\n")
	var keys [][]int
	var locks [][]int
	for _, entry := range entries {
		rows := strings.Split(entry, "\n")

		if rows[0][0] == '#' {
			key := make([]int, 5, 5)
			for _, row := range rows[1:] {
				key[0] += toInt(row[0])
				key[1] += toInt(row[1])
				key[2] += toInt(row[2])
				key[3] += toInt(row[3])
				key[4] += toInt(row[4])
			}
			keys = append(keys, key)
		} else {
			lock := make([]int, 5, 5)
			lockRows := rows[:len(rows)-1]
			slices.Reverse(lockRows)
			for _, row := range lockRows {
				lock[0] += toInt(row[0])
				lock[1] += toInt(row[1])
				lock[2] += toInt(row[2])
				lock[3] += toInt(row[3])
				lock[4] += toInt(row[4])
			}
			locks = append(locks, lock)
		}
	}

	var result int
	for _, key := range keys {
		for _, lock := range locks {
			if isMatching(key, lock) {
				result++
			}
		}
	}

	return result
}

func isMatching(key, lock []int) bool {
	for i := 0; i < len(lock); i++ {
		if key[i]+lock[i] > 5 {
			return false
		}
	}

	return true
}

func toInt(c uint8) int {
	if c == '#' {
		return 1
	}
	return 0
}
