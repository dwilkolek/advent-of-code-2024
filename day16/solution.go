package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day 16, part 1: %d", solve())
}

func Part2() {
	logger.Printf("Day 16, part 2: %d", solve())
}

type coord struct {
	x, y int
}
type Reindeer struct {
	position coord
	score    int
	facing   coord
}

var dirs = []coord{
	{
		x: 0,
		y: 1,
	}, {
		x: 1,
		y: 0,
	}, {
		x: 0,
		y: -1,
	}, {
		x: -1,
		y: 0,
	},
}

func (r Reindeer) move(area map[coord]string, step coord) (Reindeer, error) {
	nextPosition := coord{r.position.x + step.x, r.position.y + step.y}
	ch, ok := area[nextPosition]
	if !ok {
		return Reindeer{}, fmt.Errorf("step coordinate not found")
	}

	if ch == "#" {
		return Reindeer{}, fmt.Errorf("step coordinate is a wall")
	}

	if r.facing.x == step.x && r.facing.y == step.y {
		return Reindeer{
			position: nextPosition,
			score:    r.score + 1,
			facing:   r.facing,
		}, nil
	} else {
		return Reindeer{
			position: nextPosition,
			score:    r.score + 1000 + 1,
			facing:   step,
		}, nil
	}

}

func solve() int {
	file, _ := os.Open("day16/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	area := map[coord]string{}
	start := coord{}
	end := coord{}
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		for x, char := range chars {
			if char == "S" {
				start = coord{x, y}
			}
			if char == "E" {
				end = coord{x, y}
			}
			area[coord{x, y}] = char
		}
		y++
	}

	possibleCases := []Reindeer{
		{
			position: coord{
				x: start.x,
				y: start.y,
			},
			facing: coord{
				x: 1,
				y: 0,
			},
			score: 0,
		},
	}
	visited := map[coord]int{}
	for len(possibleCases) > 0 {
		nextPossibleCases := []Reindeer{}
		for _, possibleCase := range possibleCases {
			for _, dir := range dirs {
				if possibleCase.position.x == end.x && possibleCase.position.y == end.y {
					continue
				}
				newCase, err := possibleCase.move(area, dir)
				if err == nil {
					oldBestScore, ok := visited[newCase.position]
					if !ok || oldBestScore > newCase.score {
						visited[newCase.position] = newCase.score
					} else {
						continue
					}
					if newCase.position.x == end.x && newCase.position.y == end.y {
						continue
					}
					nextPossibleCases = append(nextPossibleCases, newCase)
				}
			}
		}
		possibleCases = nextPossibleCases

	}
	for k, v := range visited {
		logger.Printf("Reindeer %d, %d is visited at position %d", k.x, k.y, v)
	}
	printMap(area, visited)
	return visited[end]
}

func printMap(area map[coord]string, visited map[coord]int) {
	m := ""
	for y := 0; y < 15; y++ {
		for x := 0; x < 15; x++ {
			if visited[coord{x, y}] > 0 {
				m += "*"
			} else {
				m += area[coord{x, y}]
			}
		}
		m += "\n"
	}
	fmt.Println(m)
}
