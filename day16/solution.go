package day16

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var logger = log.Default()
var bestScore, bestSeats = math.MaxInt32, math.MaxInt32

func Part1() {
	bests := solve1(true, math.MaxInt32)
	logger.Printf("Day 16, part 1: %d", bests[0].score)
}

func Part2() {
	best1 := solve1(true, math.MaxInt32)
	bests := solve1(false, best1[0].score)
	seats := map[coord]bool{}
	for _, b := range bests {
		seats[b.position] = true
		for k, _ := range b.track {
			seats[k] = true
		}
	}

	logger.Printf("Day 16, part 2: %d", len(seats))
}

type coord struct {
	x, y int
}
type Reindeer struct {
	position coord
	score    int
	facing   coord
	track    map[coord]Reindeer
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
	_, beenThere := r.track[nextPosition]
	if beenThere {
		return r.track[nextPosition], fmt.Errorf("been there")
	}
	ch, ok := area[nextPosition]
	if !ok {
		return Reindeer{}, fmt.Errorf("step coordinate not found")
	}

	if ch == "#" {
		return Reindeer{}, fmt.Errorf("step coordinate is a wall")
	}
	newTrack := map[coord]Reindeer{}
	for k, v := range r.track {
		newTrack[k] = v
	}
	newTrack[nextPosition] = r
	if r.facing.x == step.x && r.facing.y == step.y {
		return Reindeer{
			position: nextPosition,
			score:    r.score + 1,
			facing:   r.facing,
			track:    newTrack,
		}, nil
	} else {
		return Reindeer{
			position: nextPosition,
			score:    r.score + 1000 + 1,
			facing:   step,
			track:    newTrack,
		}, nil
	}
}

func (r Reindeer) uniqueId() string {
	uniqueId := cacheKey(r)
	for _, track := range r.track {
		uniqueId += ";" + track.uniqueId()
	}
	return uniqueId
}

func cacheKey(r Reindeer) string {
	return fmt.Sprintf("pos=%d,%d , f=%d,%d", r.position.x, r.position.y, r.facing.x, r.facing.y)
}

func parse() (map[coord]string, coord, coord) {
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

	return area, start, end
}
func solve1(lookForTheBest bool, skipAbove int) []Reindeer {
	area, start, end := parse()
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
	visited := map[string]int{}
	var bestReindeers []Reindeer
	for len(possibleCases) > 0 {
		//logger.Printf("Queue size %d", len(possibleCases))
		nextPossibleCases := []Reindeer{}
		for _, possibleCase := range possibleCases {
			if possibleCase.score > skipAbove {
				continue
			}
			for _, dir := range dirs {
				newCase, err := possibleCase.move(area, dir)

				if err != nil {
					continue
				}
				oldBestScore, ok := visited[cacheKey(newCase)]
				if !ok || (lookForTheBest && newCase.score < oldBestScore) || (!lookForTheBest && newCase.score <= oldBestScore) {
					visited[cacheKey(newCase)] = newCase.score
				} else {
					continue
				}

				if newCase.position.x == end.x && newCase.position.y == end.y {
					bestReindeers = append(bestReindeers, newCase)
					if len(bestReindeers) == 0 || bestReindeers[0].score > newCase.score {
						bestReindeers = []Reindeer{newCase}
					} else {
						bestReindeers = append(bestReindeers, newCase)
					}
					//logger.Printf("Wining %d", len(bestReindeers))
					continue
				}
				if newCase.score < skipAbove {
					nextPossibleCases = append(nextPossibleCases, newCase)
				}
			}
		}
		possibleCases = nextPossibleCases

	}

	return bestReindeers
}

func solve2() int {
	//exlusions := map[string]bool{}
	//firstBest := solve1(exlusions)
	//solutions := []Reindeer{
	//	firstBest,
	//}
	//for {
	//	solved := solve1(exlusions)
	//	if solved.score == firstBest.score {
	//		solutions = append(solutions, solved)
	//	} else {
	//		break
	//	}
	//}
	//
	//coolSpots := map[coord]bool{}
	//for _, solution := range solutions {
	//	coolSpots[solution.position] = true
	//	for _, past := range solution.track {
	//		coolSpots[past.position] = true
	//	}
	//}
	//return len(coolSpots)
	return 1
}
