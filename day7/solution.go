package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type op int

const (
	ADD op = iota
	MUL
	CONCAT
)

var logger = log.Default()

func Part1() {
	sum := solve([]op{
		ADD, MUL,
	})
	logger.Printf("Day 7, part 1: %d", sum)
}

func Part2() {
	sum := solve([]op{
		ADD, MUL, CONCAT,
	})
	logger.Printf("Day 7, part 2: %d", sum)
}

func solve(ops []op) uint64 {
	file, _ := os.Open("day7/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ": ")
		resultStr := split[0]
		numbersPart := split[1]
		result, _ := strconv.ParseUint(resultStr, 10, 64)
		numbersStrs := strings.Split(numbersPart, " ")

		numbers := make([]uint64, len(numbersStrs))
		for i, numStr := range numbersStrs {
			numbers[i], _ = strconv.ParseUint(numStr, 10, 64)
		}

		valid := isValid(0, 0, numbers, result, ops)
		if valid {
			sum += result
		}
	}

	return sum
}

func isValid(index int, value uint64, numbers []uint64, result uint64, ops []op) bool {
	if index == len(numbers) {
		return value == result
	}

	if value > result {
		return false
	}
	ok := false
	for _, op := range ops {
		switch op {
		case ADD:
			ok = ok || isValid(index+1, value+numbers[index], numbers, result, ops)
			break
		case MUL:
			ok = ok || isValid(index+1, value*numbers[index], numbers, result, ops)
			break
		case CONCAT:
			strN := fmt.Sprintf("%d%d", value, numbers[index])
			n, _ := strconv.ParseUint(strN, 10, 64)
			ok = ok || isValid(index+1, n, numbers, result, ops)
			break
		}
	}
	return ok
}
