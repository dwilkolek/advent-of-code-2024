package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day 11, part 1: %d", solve2(25))
}

func Part2() {
	logger.Printf("Day 11, part 2: %d", solve2(75))
}

func nextStone(stone int) []int {
	if stone == 0 {
		return []int{1}
	} else if len(fmt.Sprintf("%d", stone))%2 == 0 {
		ns := fmt.Sprintf("%d", stone)
		a, _ := strconv.Atoi(ns[:len(ns)/2])
		b, _ := strconv.Atoi(ns[len(ns)/2:])
		return []int{a, b}
	} else {
		return []int{stone * 2024}
	}
}

type cacheKey struct {
	stone int
	iter  int
}

var cache = map[cacheKey]int{}

func countStones(stone int, iterLeft int) int {
	counted, ok := cache[cacheKey{stone, iterLeft}]
	if ok {
		return counted
	}
	if 0 == iterLeft {
		return 1
	}
	newStones := nextStone(stone)
	count := 0
	for _, nStone := range newStones {
		count += countStones(nStone, iterLeft-1)
	}
	cache[cacheKey{stone, iterLeft}] = count
	return count
}
func solve2(maxIter int) int {
	file, _ := os.Open("day11/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var stones []int
		for _, stoneStr := range strings.Split(line, " ") {
			stone, _ := strconv.Atoi(stoneStr)
			stones = append(stones, stone)
		}

		count := 0
		for _, stone := range stones {
			count += countStones(stone, maxIter)
		}

		return count
	}
	return -1
}
