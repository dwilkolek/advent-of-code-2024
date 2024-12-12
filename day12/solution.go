package day12

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var logger = log.Default()

func Part1() {
	price := 0
	for _, g := range solve() {
		price += g.area * g.perimeter
	}
	logger.Printf("Day 12, part 1: %d", price)
}

func Part2() {
	price := 0
	for _, g := range solve() {
		price += g.area * g.sides
	}
	logger.Printf("Day 12, part 2: %d", price)
}

type position struct {
	x, y int
}

func solve() []*plantGroup {
	file, _ := os.Open("day12/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	garden := map[position]string{}
	y := 0
	sizeX := 0
	for scanner.Scan() {

		line := scanner.Text()
		plants := strings.Split(line, "")
		sizeX = len(plants)
		for x, plant := range plants {
			garden[position{x: x, y: y}] = plant
		}

		y += 1
	}
	sizeY := y

	accounted := map[position]int{}
	for pos, _ := range garden {
		accounted[position{x: pos.x, y: pos.y}] = -1
	}

	var groups []*plantGroup
	for pos, _ := range accounted {
		if accounted[pos] == -1 {
			group := createPlantGroup(pos, garden, accounted)
			groups = append(groups, &group)
		}
	}

	for _, group := range groups {
		group.sides = findSides(group.groupId, accounted, sizeX, sizeY)
	}

	return groups
}

func findSides(groupId int, accounted map[position]int, sizeX int, sizeY int) int {
	sides := 0

	for x := 0; x < sizeX; x++ {
		continueWallL := false
		continueWallR := false
		for y := 0; y < sizeY; y++ {
			c := accounted[position{x: x, y: y}]
			l := accounted[position{x: x - 1, y: y}]
			r := accounted[position{x: x + 1, y: y}]
			isWallFromL := c == groupId && l != groupId
			isWallFromR := c == groupId && r != groupId
			if isWallFromL && !continueWallL {
				continueWallL = true
				sides++
			}
			if isWallFromR && !continueWallR {
				continueWallR = true
				sides++
			}
			if !isWallFromL && continueWallL {
				continueWallL = false
			}
			if !isWallFromR && continueWallR {
				continueWallR = false
			}
		}
	}
	for y := 0; y < sizeY; y++ {
		continueWallT := false
		continueWallB := false

		for x := 0; x < sizeX; x++ {
			isWallFromT := accounted[position{x: x, y: y}] == groupId && accounted[position{x: x, y: y - 1}] != groupId
			isWallFromB := accounted[position{x: x, y: y}] == groupId && accounted[position{x: x, y: y + 1}] != groupId
			if isWallFromT && !continueWallT {
				continueWallT = true
				sides++
			}
			if isWallFromB && !continueWallB {
				continueWallB = true
				sides++
			}
			if !isWallFromT && continueWallT {
				continueWallT = false
			}
			if !isWallFromB && continueWallB {
				continueWallB = false
			}
		}
	}

	return sides
}

type plantGroup struct {
	groupId   int
	plant     string
	area      int
	perimeter int
	sides     int
}

func createPlantGroup(pos position, garden map[position]string, accounted map[position]int) plantGroup {
	groupId := 1
	for _, accountedPlant := range accounted {
		if accountedPlant+1 > groupId {
			groupId = accountedPlant + 1
		}

	}
	if groupId < 1 {
		groupId = 1
	}

	area, perimeter, visited := findGroup(groupId, pos, garden, map[position]int{}, 1, 0)

	for k, _ := range visited {
		_, ok := accounted[k]
		if ok {
			accounted[k] = groupId
		}
	}

	group := plantGroup{
		groupId:   groupId,
		plant:     garden[pos],
		area:      area,
		perimeter: perimeter,
	}

	return group
}

var dirs = []position{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func findGroup(groupId int, pos position, garden map[position]string, visited map[position]int,
	area int, perimeter int) (int, int, map[position]int) {
	if visited[pos] == groupId {
		return area, perimeter, visited
	} else {
		visited[pos] = groupId
	}
	for _, dir := range dirs {
		plant := garden[pos]
		nPos := position{
			x: pos.x + dir.x,
			y: pos.y + dir.y,
		}
		nPosG, done := visited[nPos]
		if done && nPosG == groupId {
			continue
		}
		neighbour, nok := garden[nPos]
		if nok && neighbour == plant {
			area += 1
			area, perimeter, visited = findGroup(groupId, nPos, garden, visited, area, perimeter)
		} else {
			perimeter += 1
		}
	}

	return area, perimeter, visited
}
