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
	solve()
	logger.Printf("Day 11, part 1: %d", solve())
}

func Part2() {

	logger.Printf("Day 11, part 2: %d", solve())
}

func solve() int {
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

		for i := 0; i < 25; i++ {
			var newStones []int
			//logger.Printf("Loop %d", i)
			for _, stone := range stones {
				if stone == 0 {
					newStones = append(newStones, 1)
				} else if len(fmt.Sprintf("%d", stone))%2 == 0 {
					ns := fmt.Sprintf("%d", stone)
					a, _ := strconv.Atoi(ns[:len(ns)/2])
					newStones = append(newStones, a)
					b, _ := strconv.Atoi(ns[len(ns)/2:])
					newStones = append(newStones, b)
				} else {
					newStones = append(newStones, stone*2024)
				}
			}
			stones = newStones
		}

		return len(stones)
	}
	return 0
}
