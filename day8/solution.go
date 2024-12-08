package day8

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var logger = log.Default()

func Part1() {
	area := readMap()
	antiNodes := calculateAntiNodes(area)
	sum := solve(area, antiNodes, false)
	logger.Printf("Day 8, part 1: %d", sum)
}

func Part2() {
	area := readMap()
	antiNodes := calculateAntiNodes2(area)
	sum := solve(area, antiNodes, false)
	logger.Printf("Day 8, part 2: %d", sum)
}
func solve(area area, nodes map[string][]p, draw bool) int {
	unqNodes := map[p]bool{}
	for _, nodesType := range nodes {
		for _, node := range nodesType {
			unqNodes[node] = true
		}
	}
	count := 0
	for y := 0; y < area.maxY; y++ {
		for x := 0; x < area.maxX; x++ {
			_, ok := unqNodes[p{x, y}]
			if ok {
				count++
				if draw {
					print("#")
				}
			} else {
				if draw {
					print(".")
				}
			}
		}
		if draw {
			println()
		}
	}
	return count
}
func calculateAntiNodes(area area) map[string][]p {
	antiNodes := map[string][]p{}
	for k, v := range area.antennas {
		for ai, antenna := range v {
			for oai, otherAntenna := range v {
				if ai == oai {
					continue
				}
				distX := antenna.x - otherAntenna.x
				distY := antenna.y - otherAntenna.y
				antiNodes[k] = append(antiNodes[k], p{antenna.x + distX, antenna.y + distY})
			}
		}
	}
	return antiNodes
}

func calculateAntiNodes2(area area) map[string][]p {
	antiNodes := map[string][]p{}
	for k, v := range area.antennas {
		for ai, antenna := range v {
			for oai, otherAntenna := range v {
				if ai == oai {
					continue
				}
				distX := antenna.x - otherAntenna.x
				distY := antenna.y - otherAntenna.y
				for times := 0; times < 1000; times++ {
					newPos := p{antenna.x + times*distX, antenna.y + times*distY}
					if newPos.x >= 0 && newPos.y >= 0 && newPos.x < area.maxX && newPos.y < area.maxY {
						antiNodes[k] = append(antiNodes[k], newPos)
					} else {
						break
					}
				}

			}
		}
	}
	return antiNodes
}

func readMap() area {
	file, _ := os.Open("day8/input.txt")
	defer file.Close()
	antennas := map[string][]p{}
	usedPlaces := map[p]bool{}
	scanner := bufio.NewScanner(file)
	y := 0
	maxX := 0
	for scanner.Scan() {
		text := scanner.Text()
		for x, c := range strings.Split(text, "") {
			if c != "." {
				antennas[c] = append(antennas[c], p{x, y})
				usedPlaces[p{x, y}] = true
			}
			if x > maxX {
				maxX = x
			}

		}
		y += 1
	}
	return area{
		antennas: antennas,
		maxX:     maxX + 1,
		maxY:     y,
	}
}

type area struct {
	antennas map[string][]p
	maxX     int
	maxY     int
}
type p struct {
	x int
	y int
}
