package day6

import (
	"bufio"
	"log"
	"maps"
	"os"
	"strings"
)

var logger = log.Default()

func Part1() {
	person, area := read()
	for moveAndMark(&person, area) == CONTINUE {
	}

	logger.Printf("Day 6, part 1: %d", len(person.visited))
}

func Part2() {
	person, areaObj := read()
	looped := 0

	for y := 0; y <= areaObj.maxY; y++ {
		for x := 0; x <= areaObj.maxX; x++ {
			personCopy := pers{
				dir:     person.dir,
				x:       person.x,
				y:       person.y,
				visited: make(map[pos]dir),
			}
			areaObjCopy := area{
				obstacles: maps.Clone(areaObj.obstacles),
				maxX:      areaObj.maxX,
				maxY:      areaObj.maxY,
			}
			_, has := areaObjCopy.obstacles[pos{
				x, y,
			}]
			if has {
				continue
			}
			areaObjCopy.obstacles[pos{
				x, y,
			}] = true
			keepRunning := true
			for keepRunning {
				result := moveAndMark(&personCopy, areaObjCopy)
				switch result {
				case STOP:
					keepRunning = false
					break
				case STOP_LOOPED:
					looped += 1
					keepRunning = false
					break
				case CONTINUE:
					break
				}
			}
		}
	}

	logger.Printf("Day 6, part 2: %d", looped)
}

func moveAndMark(person *pers, area area) res {
	diffX := 0
	diffY := 0
	nextDir := person.dir

	switch person.dir {
	case UP:
		diffY = -1
		nextDir = RIGHT
	case DOWN:
		diffY = 1
		nextDir = LEFT
	case LEFT:
		diffX = -1
		nextDir = UP
	case RIGHT:
		diffX = 1
		nextDir = DOWN
	}

	currX := person.x
	currY := person.y

	for {
		nextX := currX + diffX
		nextY := currY + diffY
		if nextX < 0 || nextY < 0 || nextX > area.maxX || nextY > area.maxY {
			return STOP
		}
		visitedDir, visited := person.visited[pos{
			x: nextX, y: nextY,
		}]
		if visited && visitedDir == person.dir {
			return STOP_LOOPED
		}

		_, ok := area.obstacles[pos{
			x: nextX,
			y: nextY,
		}]

		if ok {
			person.x = currX
			person.y = currY
			person.dir = nextDir
			return CONTINUE
		}

		currY = nextY
		currX = nextX

		person.visited[pos{
			x: currX,
			y: currY,
		}] = person.dir
	}

}

type area struct {
	obstacles map[pos]bool
	maxX      int
	maxY      int
}

func read() (pers, area) {
	file, _ := os.Open("day6/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	obstacles := make(map[pos]bool)
	person := pers{}
	maxX := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		for x, char := range chars {
			if char == "#" {
				obstacles[pos{
					x, y,
				}] = true
			} else if char == "^" {
				person.x = x
				person.y = y
				person.dir = UP
				person.visited = make(map[pos]dir)
				person.visited[pos{x, y}] = UP
			}
			if maxX < x {
				maxX = x
			}
		}

		y += 1

	}

	return person, area{
		obstacles: obstacles,
		maxX:      maxX,
		maxY:      y - 1,
	}
}

type pers struct {
	dir     dir
	x       int
	y       int
	visited map[pos]dir
}
type pos struct {
	x int
	y int
}
type dir int

const (
	UP dir = iota
	DOWN
	LEFT
	RIGHT
)

type res int

const (
	CONTINUE res = iota
	STOP
	STOP_LOOPED
)
