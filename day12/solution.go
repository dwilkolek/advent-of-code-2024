package day12

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day 12, part 1: %d", solve())
}

func Part2() {
	logger.Printf("Day 11, part 2: %d", 0)
}

type position struct {
	x, y int
}

func solve() int {
	file, _ := os.Open("day12/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	garden := map[position]string{}
	y := 0
	for scanner.Scan() {

		line := scanner.Text()
		for x, plant := range strings.Split(line, "") {
			garden[position{x: x, y: y}] = plant
		}

		y += 1
	}

	accounted := map[position]int{}
	for pos, _ := range garden {
		accounted[position{x: pos.x, y: pos.y}] = -1
	}
	price := 0
	for pos, _ := range accounted {
		if accounted[pos] == -1 {
			gr := createPlantGroup(pos, garden, accounted)
			logger.Printf("group=%s area=%d perimiter=%d", gr.plant, gr.area, gr.perimiter)
			price += gr.perimiter * gr.area
		}
	}

	return price
}

type plantGroup struct {
	groupId   int
	plant     string
	area      int
	perimiter int
}

func createPlantGroup(pos position, garden map[position]string, accounted map[position]int) plantGroup {
	groupId := 1
	for _, accountedPlant := range accounted {
		groupId = accountedPlant + 1
	}
	if groupId < 1 {
		groupId = 1
	}

	area, perimiter, visited := findGroup(groupId, pos, garden, map[position]int{}, 1, 0)

	for k, v := range visited {
		_, ok := accounted[k]
		if ok {
			accounted[k] += v
		}
	}

	group := plantGroup{
		groupId:   groupId,
		plant:     garden[pos],
		area:      area,
		perimiter: perimiter,
	}

	return group
}

var dirs = []position{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func findGroup(groupId int, pos position, garden map[position]string, visited map[position]int, area int, perimiter int) (int, int, map[position]int) {
	nArea := area
	nPerimiter := perimiter
	if visited[pos] == groupId {
		return nArea, nPerimiter, visited
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
			nArea += 1
			nArea, nPerimiter, visited = findGroup(groupId, nPos, garden, visited, nArea, nPerimiter)
		} else {
			nPerimiter += 1
		}
	}

	return nArea, nPerimiter, visited
}
