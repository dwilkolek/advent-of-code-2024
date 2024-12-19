package day18

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

var logger = log.Default()
var debug = false

func Part1() {
	bytes := parse()
	logger.Printf("Day 18, part 1: %d", solve(bytes[:12]))
}

func Part2() {
	bytes := parse()
	i := 0
	step := len(bytes) / 10
	for {
		b := bytes[:i]
		s := solve(b)
		if s == math.MaxInt64-1 {
			if step == 1 {
				logger.Printf("Day 18, part 2: %d,%d", b[len(b)-1].x, b[len(b)-1].y)
				return
			}
			i -= step
			if step < 10 {
				step = 1
			} else {
				step = step / 10
			}
		} else {
			i = min(len(bytes), step+i)
		}

	}
}

type coord struct{ x, y int }

var size = 70

func parse() []coord {
	file, _ := os.Open("day18/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bytes := []coord{}
	for scanner.Scan() {
		c := coord{}
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d", &c.x, &c.y)
		if err != nil {
			logger.Fatal(err)
		}
		bytes = append(bytes, c)
	}
	return bytes
}
func solve(inB []coord) int {
	bytes := map[coord]bool{}
	for _, in := range inB {
		bytes[in] = true
	}
	logger.SetOutput(os.Stdout)

	bestStepCount := math.MaxInt64
	toCheck := []wanderer{
		{
			pos: coord{0, 0},
			history: map[coord]bool{
				coord{0, 0}: true,
			},
		},
	}
	cachedBest := map[coord]int{}
	toCheck[0].print(bytes)

	for len(toCheck) > 0 {
		w := toCheck[0]
		toCheck = toCheck[1:]

		best, ok := cachedBest[w.pos]
		if !ok || best > len(w.history) {
			cachedBest[w.pos] = len(w.history)
		} else {
			continue
		}

		if w.pos.y == size && w.pos.x == size {
			if len(w.history) < bestStepCount {
				bestStepCount = len(w.history)
				//logger.Printf("Current Best: %d", bestStepCount)
				w.print(bytes)
			}
			continue
		}

		toCheck = append(toCheck, w.move(bytes)...)
	}

	return bestStepCount - 1
}

var dirs = []coord{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

type wanderer struct {
	pos     coord
	history map[coord]bool
	dist    int
}

func (w wanderer) move(errors map[coord]bool) []wanderer {
	nextW := []wanderer{}
	for _, dir := range dirs {
		nextPos := coord{w.pos.x + dir.x, w.pos.y + dir.y}

		if errors[nextPos] {
			continue
		}
		if nextPos.x < 0 || nextPos.y < 0 || nextPos.x > size || nextPos.y > size {
			continue
		}

		if w.history[nextPos] {
			continue
		}
		nHist := map[coord]bool{}
		for k, v := range w.history {
			nHist[k] = v
		}
		nHist[nextPos] = true

		dist := size - nextPos.x + size - nextPos.y
		nextW = append(nextW, wanderer{nextPos, nHist, dist})
	}

	return nextW
}

func (w *wanderer) print(errors map[coord]bool) {
	if !debug {
		return
	}
	m := "\n"
	for y := 0; y <= size; y++ {
		for x := 0; x <= size; x++ {
			if w.history[coord{x, y}] {
				m += "*"
			} else if errors[coord{x, y}] {
				m += "#"
			} else {
				m += "."
			}
		}
		m += "\n"
	}
	logger.Println(m)
}
