package day13

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day 13, part 1: %d", solve(0, true))
}

func Part2() {
	logger.Printf("Day 13, part 2: %d", solve(10000000000000, false))
}

func solve(offset int, hasClickLimit bool) int {
	prizes := getPrizes(offset)

	sum := 0
	for _, p := range prizes {
		sum += calculateTokenSpent(p, hasClickLimit)
	}
	return sum
}

func calculateTokenSpent(p prize, hasClickLimit bool) int {
	a := (p.prize.x*p.b.y - p.prize.y*p.b.x) / (p.a.x*p.b.y - p.a.y*p.b.x)
	b := (p.prize.y*p.a.x - p.prize.x*p.a.y) / (p.a.x*p.b.y - p.a.y*p.b.x)
	if hasClickLimit && (a > 100 || b > 100) {
		return 0
	}

	if a*p.a.x+b*p.b.x == p.prize.x && a*p.a.y+b*p.b.y == p.prize.y {
		return a*p.a.cost + b*p.b.cost
	}

	return 0
}

func getPrizes(offset int) []prize {
	file, _ := os.Open("day13/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prizes []prize
	prizeInProgress := prize{}
	doA := true
	doB := false
	doPrize := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if doA {
			coords := strings.Split(strings.TrimPrefix(line, "Button A: "), ", ")
			xa, _ := strconv.Atoi(strings.TrimPrefix(coords[0], "X"))
			ya, _ := strconv.Atoi(strings.TrimPrefix(coords[1], "Y"))
			prizeInProgress.a = button{coordinates: coordinates{x: xa, y: ya}, cost: 3}
			doB = true
			doA = false
		} else if doB {
			coords := strings.Split(strings.TrimPrefix(line, "Button B: "), ", ")
			xa, _ := strconv.Atoi(strings.TrimPrefix(coords[0], "X"))
			ya, _ := strconv.Atoi(strings.TrimPrefix(coords[1], "Y"))
			prizeInProgress.b = button{coordinates: coordinates{x: xa, y: ya}, cost: 1}
			doPrize = true
			doB = false
		} else if doPrize {
			coords := strings.Split(strings.TrimPrefix(line, "Prize: "), ", ")
			xa, _ := strconv.Atoi(strings.TrimPrefix(coords[0], "X="))
			ya, _ := strconv.Atoi(strings.TrimPrefix(coords[1], "Y="))
			prizeInProgress.prize = coordinates{x: xa + offset, y: ya + offset}
			doPrize = true
			prizes = append(prizes, prizeInProgress)
			prizeInProgress = prize{}
			doA = true
			doPrize = false
		}
	}
	return prizes
}

type coordinates struct{ x, y int }
type button struct {
	coordinates
	cost int
}

type prize struct {
	a     button
	b     button
	prize coordinates
}
