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
	antonodes := calculateAntiNodes(area)
	sum := draw(area, antonodes)
	//sum := 0
	//for _, node := range antonodes {
	//	sum += len(node)
	//}
	logger.Printf("Day 8, part 1: %d", sum)
}

func Part2() {
	area := readMap()
	antonodes := calculateAntiNodes(area)
	sum := draw(area, antonodes)
	//sum := 0
	//for _, node := range antonodes {
	//	sum += len(node)
	//}
	logger.Printf("Day 8, part 2: %d", sum)
}
func draw(area area, nodes map[string][]p) int {
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
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
	return count
}
func calculateAntiNodes(area area) map[string][]p {
	antinodes := map[string][]p{}
	for k, v := range area.antennas {
		logger.Printf("Doing %s", k)
		for ai, antena := range v {
			for oai, otherAntena := range v {
				if ai == oai {
					continue
				}
				distX := antena.x - otherAntena.x
				distY := antena.y - otherAntena.y
				//_, used := area.usedPlaces[p{antena.x + distX, antena.y + distY}]
				antinodes[k] = append(antinodes[k], p{antena.x + distX, antena.y + distY})
			}
		}
	}
	return antinodes
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
		antennas:   antennas,
		maxX:       maxX + 1,
		maxY:       y,
		usedPlaces: usedPlaces,
	}
}

type area struct {
	antennas   map[string][]p
	maxX       int
	maxY       int
	usedPlaces map[p]bool
}
type p struct {
	x int
	y int
}
