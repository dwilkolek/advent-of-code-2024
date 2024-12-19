package day18

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

var logger = log.Default()
var debug = true

func Part1() {
	logger.Printf("Day X, part 1: %d", solve())
}

func Part2() {
	logger.Printf("Day X, part 2: %d", solve())
}

type coord struct{ x, y int }

var size = 70
var limit = 1024

//var size = 6
//var limit = 12

func solve() int {

	logger.SetOutput(os.Stdout)
	file, _ := os.Open("day18/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bytes := map[coord]bool{}
	for scanner.Scan() {
		c := coord{}
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d", &c.x, &c.y)
		if err != nil {
			logger.Fatal(err)
		}
		bytes[c] = true
		if len(bytes) == limit {
			break
		}
	}

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
		//logger.Printf("Queue size: %d", len(toCheck))
		w := toCheck[0]
		toCheck = toCheck[1:]
		//w.print(bytes)
		//cb, ok := cachedBest[w.pos]
		//if ok && cb < len(w.history) {
		//	continue
		//}
		//if len(w.history) >= bestStepCount {
		//	continue
		//}
		best, ok := cachedBest[w.pos]
		if !ok || best > len(w.history) {
			cachedBest[w.pos] = len(w.history)
		} else {
			continue
		}

		if w.pos.y == size && w.pos.x == size {
			if len(w.history) < bestStepCount {
				bestStepCount = len(w.history)
				logger.Printf("Current Best: %d", bestStepCount)
				w.print(bytes)
			}
			continue
		}

		toCheck = append(toCheck, w.move(bytes)...)
		//slices.SortFunc(toCheck, func(a, b wanderer) int {
		//	return cmp.Compare(a.dist+len(a.history), b.dist+len(b.history))
		//})
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
