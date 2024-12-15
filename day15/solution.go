package day15

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day 15, part 1: %d", solve(false))
}

func Part2() {
	logger.Printf("Day 15, part 2: %d", solve(true))
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

func solve(isPart2 bool) int {
	file, _ := os.Open("day15/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	area := map[coord]string{}
	var robot coord
	var moves []string
	y := -1
	readMap := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readMap = false
			continue
		}

		chars := strings.Split(line, "")
		if readMap {
			if isPart2 {
				newChars := make([]string, 0)
				for _, char := range chars {
					//If the tile is #, the new map contains ## instead.
					//If the tile is O, the new map contains [] instead.
					//If the tile is ., the new map contains .. instead.
					//If the tile is @, the new map contains @. instead.
					if char == "#" {
						newChars = append(newChars, "#", "#")
					}
					if char == "O" {
						newChars = append(newChars, "[", "]")
					}
					if char == "." {
						newChars = append(newChars, ".", ".")
					}
					if char == "@" {
						newChars = append(newChars, "@", ".")
					}

				}
				chars = newChars
			}
			y += 1
			for x, ch := range chars {
				if ch == "." {
					continue
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

	for _, moveCh := range moves {
		move := dirs[moveCh]
		moveRobot(&robot, move, area)
	}

	return score(area)
}

func score(area map[coord]string) int {
	score := 0
	for coord, ch := range area {
		if ch == "O" || ch == "[" {
			score += 100*coord.y + coord.x
		}
	}
	return score
}

func printMap(r *coord, area map[coord]string, sizeX, sizeY int, me string) {
	m := ""
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			ch, ok := area[coord{x: x, y: y}]
			if ok {
				m += ch
			} else if r.y == y && r.x == x {
				m += me
			} else {
				m += "."
			}
		}
		m += "\n"
	}
	logger.Println("\n" + m)
}

func moveRobot(r *coord, move coord, area map[coord]string) {
	clonedMap := cloneMap(area)
	canRobotMove := prepForMove(coord{
		x: r.x + move.x,
		y: r.y + move.y,
	}, move, clonedMap)

	if canRobotMove {
		prepForMove(coord{
			x: r.x + move.x,
			y: r.y + move.y,
		}, move, area)
		r.x = r.x + move.x
		r.y = r.y + move.y
	}
}

func cloneMap(area map[coord]string) map[coord]string {
	clone := make(map[coord]string)
	for k, v := range area {
		clone[k] = v
	}
	return clone
}

func prepForMove(dest coord, move coord, area map[coord]string) bool {
	ch, ok := area[dest]
	if !ok {
		return true
	}
	if ch == "O" {
		to := coord{x: dest.x + move.x, y: dest.y + move.y}
		isTherePlace := prepForMove(to, move, area)
		if isTherePlace {
			doMove([]coord{dest}, move, area)
			return true
		}
		return isTherePlace
	}
	if ch == "[" {
		var destConn = coord{x: dest.x + 1, y: dest.y}
		return doMove([]coord{dest, destConn}, move, area)
	}
	if ch == "]" {
		var destConn = coord{x: dest.x - 1, y: dest.y}
		return doMove([]coord{destConn, dest}, move, area)
	}
	return false
}

func doMove(froms []coord, move coord, area map[coord]string) bool {
	mapSizeBefore := len(area)
	isFree := true
	moveMap := make(map[coord]coord)
	wtfAmIMoving := movingStr(froms, area)
	//logger.Printf("Moving %v [%d, %d]: %s %t", froms, move.x, move.y, wtfAmIMoving)

	if wtfAmIMoving == ".." {
		return true
	}
	if strings.Contains(wtfAmIMoving, "#") {
		return false
	}

	for _, from := range froms {
		moveMap[from] = coord{x: from.x + move.x, y: from.y + move.y}
	}
	for _, to := range moveMap {
		_, isPartOfFrom := moveMap[to]
		if isPartOfFrom {
			continue
		}
		_, occupied := area[to]
		isFree = isFree && !occupied
	}
	if isFree {

		mapping := map[coord]string{}
		for from, to := range moveMap {
			mapping[to] = area[from]
			delete(area, from)
		}

		for to, v := range mapping {
			area[to] = v
		}
		mapSizeAfter := len(area)
		if mapSizeBefore != mapSizeAfter {
			logger.Panicf("After moving %d -> %d", mapSizeBefore, mapSizeAfter)
		}
		return true
	}

	if move.y == 0 {
		var nextFroms []coord
		for _, from := range froms {
			nextFroms = append(nextFroms, coord{x: from.x + 2*move.x, y: from.y})
		}
		if "[]" == movingStr(nextFroms, area) && doMove(nextFroms, move, area) {
			return doMove(froms, move, area)
		}
		return false
	}

	var nextFroms []coord

	for _, from := range froms {
		nextFroms = append(nextFroms, coord{x: from.x, y: from.y + move.y})
	}
	nextFromStr := movingStr(nextFroms, area)

	if "[]" == nextFromStr && doMove(nextFroms, move, area) {
		return doMove(froms, move, area)
	}
	if nextFromStr == "]." {
		nextFroms = []coord{{x: nextFroms[0].x - 1, y: nextFroms[0].y}, nextFroms[0]}
		//logger.Printf("Attempt from %s to %s", nextFromStr, movingStr(nextFroms, area))
		if doMove(nextFroms, move, area) {
			return doMove(froms, move, area)
		}
		return false
	}
	if nextFromStr == ".[" {
		nextFroms = []coord{nextFroms[1], {x: nextFroms[1].x + 1, y: nextFroms[1].y}}
		//logger.Printf("Attempt from %s to %s", nextFromStr, movingStr(nextFroms, area))
		if doMove(nextFroms, move, area) {
			return doMove(froms, move, area)
		}
		return false
	}
	if "][" == nextFromStr {
		nextFromsL := []coord{{x: nextFroms[0].x - 1, y: nextFroms[0].y}, nextFroms[0]}
		nextFromsR := []coord{nextFroms[1], {x: nextFroms[1].x + 1, y: nextFroms[1].y}}
		//logger.Printf("][ -> %s ; %s", movingStr(nextFromsL, area), movingStr(nextFromsR, area))
		lSucc := doMove(nextFromsL, move, area)
		rSucc := doMove(nextFromsR, move, area)

		if lSucc && rSucc {
			return doMove(froms, move, area)
		}
		return false
	}

	return false
}

func movingStr(from []coord, area map[coord]string) string {
	nextFromStr := ""
	for _, n := range from {
		c, ok := area[n]
		if ok {
			nextFromStr += c
		} else {
			nextFromStr += "."
		}

	}
	return nextFromStr
}
