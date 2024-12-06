package day6

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var logger = log.Default()

func Part1() {
	person, area := read()
	logger.Printf("%v", person)
	for moveAndMark(&person, area) {
	}

	log.Printf("Day 6, part 1: %d", len(person.visited))
}

func moveAndMark(person *pers, area area) bool {
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
			return false
		}

		_, ok := area.obstacles[nextY][nextX]

		if ok {
			person.x = currX
			person.y = currY
			person.dir = nextDir
			return true
		}

		currY = nextY
		currX = nextX

		person.visited[pos{
			x: currX,
			y: currY,
		}] = true
	}

}

type area struct {
	obstacles map[int]map[int]bool
	maxX      int
	maxY      int
}

func read() (pers, area) {
	file, _ := os.Open("day6/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	obstacles := make(map[int]map[int]bool)
	person := pers{}
	maxX := 0
	for scanner.Scan() {
		line := scanner.Text()
		obstacles[y] = make(map[int]bool)
		chars := strings.Split(line, "")
		for x, char := range chars {
			if char == "#" {
				obstacles[y][x] = true
			} else if char == "^" {
				person.x = x
				person.y = y
				person.dir = UP
				person.visited = make(map[pos]bool)
				person.visited[pos{x, y}] = true
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
	visited map[pos]bool
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
