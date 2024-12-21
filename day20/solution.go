package day20

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

var logger = log.Default()

func Part1() {
	solve(1, 2, 100) //1 is not a real shortcut, 2 is makes it across
}

func Part2() {
	solve(2, 20, 100)
}

type coord struct {
	x, y int
}

var sizeX int

func solve(part int, dist int, minimumSaved int) {
	baseCase := solveBaseCase()
	history := map[coord]int{}
	for s, sco := range baseCase.historyOrd {
		history[sco] = s
	}
	//cheaters := solveCheaters(baseCase, 64, 1)
	logger.Printf("Day 20, part %d: %d", part, findWithDist(baseCase.historyOrd, history, dist, minimumSaved))
}
func solveBaseCase() horse {
	area, start, end := parse()
	horsies := []horse{
		{
			p: start,
			historyOrd: []coord{
				start,
			},
			seen: map[coord]bool{start: true},
		},
	}

	for len(horsies) > 0 {
		h := horsies[0]
		horsies = horsies[1:]

		for _, nh := range h.move(area) {
			if nh.p.x == end.x && nh.p.y == end.y {
				return nh
			}
			horsies = append(horsies, nh)
		}
	}
	return horse{}
}

var dirs = []coord{
	{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1},
}

type horse struct {
	p          coord
	historyOrd []coord
	seen       map[coord]bool
}

type cheater struct {
	p     coord
	saved int
}

func findWithDist(ord []coord, history map[coord]int, dist int, minimumSaved int) int {

	counter := 0
	for _, start := range ord {
		for _, end := range ord {
			if history[end] <= history[start] {
				continue
			}
			//if history[end]-history[start] < minimumSaved {
			//	continue
			//}
			normalRoute := history[end] - history[start]
			shortcutDist := int(math.Abs(float64(end.x-start.x)) + math.Abs(float64(end.y-start.y)))
			saved := normalRoute-shortcutDist >= minimumSaved
			inDistance := shortcutDist <= dist
			if saved && inDistance {
				//logger.Printf("%d,%d -> %d,%d", start.x, start.y, end.x, end.y)
				counter++
			}
		}
	}
	return counter
}

func (h horse) move(area map[coord]string) []horse {
	nextH := make([]horse, 0)
	for _, dir := range dirs {
		nextP := coord{h.p.x + dir.x, h.p.y + dir.y}
		if h.seen[nextP] {
			continue
		}
		block, found := area[nextP]
		if !found {
			continue
		}
		if block != "." {
			continue
		}
		nSeen := make(map[coord]bool)
		for k, v := range h.seen {
			nSeen[k] = v
		}
		nSeen[nextP] = true

		nextH = append(nextH, horse{
			p:          nextP,
			historyOrd: append(h.historyOrd, nextP),
			seen:       nSeen,
		})
	}
	return nextH
}

//func (h horse) print(area map[coord]string) string {
//	m := fmt.Sprintf("Steps %d\n", len(h.history))
//	for y := 0; y < sizeY; y++ {
//		for x := 0; x < sizeX; x++ {
//			p := coord{x, y}
//			block := area[p]
//			_, ok := h.history[p]
//			if ok {
//				if block == "?" {
//					m += "@"
//				} else {
//					m += "*"
//				}
//			} else {
//				if block == "." {
//					m += "."
//				} else {
//					m += "#"
//				}
//			}
//		}
//		m += "\n"
//	}
//
//	//println(m)
//	return m
//}

func parse() (map[coord]string, coord, coord) {
	file, _ := os.Open("day20/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	var start coord
	var end coord
	area := map[coord]string{}
	for scanner.Scan() {
		line := scanner.Text()
		for x, ch := range strings.Split(line, "") {
			if ch == "S" {
				start = coord{x, y}
				area[coord{x: x, y: y}] = "."
			} else if ch == "E" {
				end = coord{x, y}
				area[coord{x: x, y: y}] = "."
			} else {
				area[coord{x: x, y: y}] = ch
			}
			if sizeX < x {
				sizeX = x
			}
		}
		y++
	}

	return area, start, end
}
