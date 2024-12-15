package day15

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day X, part 1: %d", solve())
}

func Part2() {
	logger.Printf("Day X, part 2: %d", solve())
}

var dirs = map[string]coord{
	"v": {
		x: 0,
		y: 1,
	},
	">": {
		x: 1,
		y: 0,
	},
	"^": {
		x: 0,
		y: -1,
	},
	"<": {
		x: -1,
		y: 0,
	},
}

type coord struct {
	x, y int
}

func solve() int {
	file, _ := os.Open("day15/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	area := map[coord]string{}
	var robot coord
	moves := []string{}
	y := -1
	readMap := true
	sizeX, sizeY := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readMap = false
			continue
		}
		chars := strings.Split(line, "")
		if readMap {
			y += 1
			for x, ch := range chars {
				if ch == "." {
					continue
				}
				if x > sizeX {
					sizeX = x
				}
				if ch == "@" {
					robot = coord{x: x, y: y}
				} else {
					area[coord{x: x, y: y}] = ch
				}
			}

		} else {
			moves = append(moves, chars...)
		}
	}
	sizeY = y + 1
	sizeX = sizeX + 1
	printMap(&robot, area, sizeX, sizeY)

	for _, moveCh := range moves {
		move := dirs[moveCh]
		moveRobot(&robot, move, area)
		//printMap(&robot, area, sizeX, sizeY)
	}

	return score(area)
}

type DataStore struct {
	datastore map[coord]string
}

func score(area map[coord]string) int {
	score := 0
	for coord, ch := range area {
		if ch == "#" {
			continue
		}
		score += 100*coord.y + coord.x
	}
	return score
}

func printMap(r *coord, area map[coord]string, sizeX, sizeY int) {
	m := ""
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			ch, ok := area[coord{x: x, y: y}]
			if ok {
				m += ch
			} else if r.y == y && r.x == x {
				m += "@"
			} else {
				m += "."
			}
		}
		m += "\n"
	}
	log.Println("\n" + m)
}

func moveRobot(r *coord, move coord, area map[coord]string) {
	canRobotMove := prepForMove(coord{
		x: r.x + move.x,
		y: r.y + move.y,
	}, move, area)
	if canRobotMove {
		r.x = r.x + move.x
		r.y = r.y + move.y
	}
}

func prepForMove(dest coord, move coord, area map[coord]string) bool {
	//
	ch, ok := area[dest]
	if !ok {
		return true
	}
	if ch == "#" {
		return false
	}
	if ch == "O" {
		to := coord{x: dest.x + move.x, y: dest.y + move.y}
		isTherePlace := prepForMove(to, move, area)
		if isTherePlace {
			doMove(dest, to, area)
			return true
		}
		return isTherePlace
	}
	return false
}

func doMove(from coord, to coord, area map[coord]string) {
	toC, tOk := area[to]
	fromC, _ := area[from]
	if fromC == "O" && !tOk {
		area[to] = fromC
		delete(area, from)
	} else {
		logger.Panicf("illegal move: from(%d, %d)=%s, to(%d, %d)=%s", from.x, from.y, fromC, to.x, to.y, toC)
	}

}
