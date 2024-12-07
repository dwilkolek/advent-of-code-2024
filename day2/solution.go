package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var logger = log.Default()

func Part1() {
	solve(1, 0)
}
func Part2() {
	solve(2, 1)
}

func solve(part int, allowErrors int) {

	file, _ := os.Open("day2/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReports := 0
	for scanner.Scan() {
		levelsStr := strings.Split(scanner.Text(), " ")
		// logger.Printf("%s ", levelsStr)
		levels := make([]int, len(levelsStr))
		for i, levelStr := range levelsStr {
			levels[i], _ = strconv.Atoi(levelStr)
		}

		valid := checkValidity(levels, len(levels)-allowErrors, 1) || checkValidity(levels, len(levels)-allowErrors, -1)
		// logger.Printf(" is %t", valid)
		if valid {
			safeReports += 1
		}
	}

	logger.Printf("Day 2, part %d: %d", part, safeReports)
}

func checkValidity(levels []int, minLength int, dir int) bool {
	// logger.Printf("\t testing(%d) %v ", dir, levels)
	isValid := true
	stoppedAt := 0
	for i := 1; i < len(levels) && isValid; i++ {
		isValid = verifyPair(levels[i], levels[i-1], dir)
		stoppedAt = i
	}

	if !isValid && minLength < len(levels) {
		isValid = checkValidity(dropOne(levels, stoppedAt-1), minLength, dir) || checkValidity(dropOne(levels, stoppedAt), minLength, dir)
	}
	// logger.Printf("\t valid %t ", isValid)
	return isValid
}

func verifyPair(level int, prev_level int, dir int) bool {
	diffBetweenLevels := (level - prev_level) * dir
	isOK := diffBetweenLevels >= 1 && diffBetweenLevels <= 3
	// logger.Printf("comp(%d) %d ? %d = %t", dir, prev_level, level, isOK)
	return isOK
}

func dropOne(levels []int, exceptIndex int) []int {
	cp := make([]int, len(levels))
	copy(cp, levels)
	return append(cp[:exceptIndex], cp[exceptIndex+1:]...)
}
