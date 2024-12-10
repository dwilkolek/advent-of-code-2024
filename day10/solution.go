package day10

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var logger = log.Default()

func Part1() {
	trailStarts, area := prepare()
	sum := 0
	for start, _ := range trailStarts {

		unqTrailHeads := map[pos]bool{}

		for _, head := range append([]pos{}, findTrailHeads(start, make(map[pos]bool), area)...) {
			unqTrailHeads[head] = true
		}
		sum += len(unqTrailHeads)
	}

	logger.Printf("Day 10, part 1: %d", sum)
}

func Part2() {
	trailStarts, area := prepare()
	trailHeadsList := []pos{}
	for start, _ := range trailStarts {
		trailHeadsList = append(trailHeadsList, findTrailHeads(start, make(map[pos]bool), area)...)
	}
	logger.Printf("Day 10, part 2: %d", len(trailHeadsList))
}

type pos struct{ x, y int }

var dirs = []pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findTrailHeads(place pos, visited map[pos]bool, area map[pos]int) []pos {
	foundHeads := []pos{}
	visited[place] = true
	for _, dir := range dirs {
		ch := area[place]
		if ch == 9 {
			return []pos{place}
		}
		nextpos := pos{x: place.x + dir.x, y: place.y + dir.y}
		_, seen := visited[nextpos]
		if seen {
			continue
		}
		nch, ok := area[nextpos]

		if !ok || nch-1 != ch {
			continue
		}
		n_visited := map[pos]bool{}
		for k, v := range visited {
			n_visited[k] = v
		}
		foundHeads = append(foundHeads, findTrailHeads(nextpos, n_visited, area)...)
	}
	return foundHeads
}

func prepare() (map[pos]bool, map[pos]int) {
	file, _ := os.Open("day10/input.txt")
	defer file.Close()

	y := 0

	trailHeads := map[pos]bool{}
	trailStarts := map[pos]bool{}

	area := map[pos]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		chars := strings.Split(line, "")

		for x, hch := range chars {
			if hch == "." {
				continue
			}
			h, _ := strconv.Atoi(hch)
			if h == 9 {
				trailHeads[pos{x: x, y: y}] = true
			}
			if h == 0 {
				trailStarts[pos{x: x, y: y}] = true
			}
			area[pos{x: x, y: y}] = h
		}

		y += 1
	}

	return trailStarts, area
}
