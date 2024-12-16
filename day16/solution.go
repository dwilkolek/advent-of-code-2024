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
	if bestScore == math.MaxInt32 {
		bestScore, bestSeats = solve()
	}
	logger.Printf("Day 16, part 1: %d", bestScore)
}

func Part2() {
	if bestSeats == math.MaxInt32 {
		bestScore, bestSeats = solve()
	}
	logger.Printf("Day 16, part 2: %d", bestSeats)
}

type coord struct {
	x, y int
}
type Reindeer struct {
	position coord
	score    int
	facing   coord
	track    []Reindeer
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
	newTrack := make([]Reindeer, len(r.track))
	for k, v := range r.track {
		newTrack[k] = v
	}
	newTrack = append(newTrack, r)
	if r.facing.x == step.x && r.facing.y == step.y {
		return Reindeer{
			position: nextPosition,
			score:    r.score + 1,
			facing:   r.facing,
			track:    newTrack,
		}, nil
	} else {
		return Reindeer{
			position: coord{r.position.x, r.position.y},
			score:    r.score + 1000,
			facing:   step,
			track:    newTrack,
		}, nil
	}

}

func cacheKey(r Reindeer) string {
	return fmt.Sprintf("pos=%d,%d, f=%d,%d", r.position.x, r.position.y, r.facing.x, r.facing.y)
}
func instantWinCacheKey(r Reindeer) string {
	return fmt.Sprintf("pos=%d,%d, f=%d,%d score=%d", r.position.x, r.position.y, r.facing.x, r.facing.y, r.score)
}
func solve() (int, int) {
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
	visited := map[string]int{}
	grantInstantWin := map[string]bool{}
	winners := []Reindeer{
		{position: end, score: math.MaxInt64, facing: coord{x: -1, y: -1}, track: []Reindeer{}},
	}

	for len(possibleCases) > 0 {
		nextPossibleCases := []Reindeer{}
		for _, possibleCase := range possibleCases {
			for _, dir := range dirs {
				newCase, err := possibleCase.move(area, dir)

				if err == nil {
					oldBestScore, ok := visited[cacheKey(newCase)]
					worthTrying := false
					if !ok || newCase.score <= oldBestScore {
						visited[cacheKey(newCase)] = newCase.score
						worthTrying = true
					}

					if grantInstantWin[instantWinCacheKey(newCase)] {
						winners = append(winners, newCase)
						for _, track := range newCase.track {
							grantInstantWin[instantWinCacheKey(track)] = true
						}
						worthTrying = false
					}

					if newCase.position.x == end.x && newCase.position.y == end.y {
						if winners[0].score > newCase.score {
							winners = make([]Reindeer, 0)
						}
						winners = append(winners, newCase)
						worthTrying = false
						grantInstantWin = map[string]bool{}
						for _, w := range winners {
							grantInstantWin[instantWinCacheKey(w)] = true
							for _, wr := range w.track {
								grantInstantWin[instantWinCacheKey(wr)] = true
							}
						}
					}

					if worthTrying {
						nextPossibleCases = append(nextPossibleCases, newCase)
					}
				}
			}
		}
		possibleCases = nextPossibleCases

	}

	bestSpots := map[coord]bool{}
	bestScore := math.MaxInt
	for _, w := range winners {
		if w.score < bestScore {
			bestScore = w.score
		}
	}
	for _, winner := range winners {
		if winner.score == bestScore {
			bestSpots[winner.position] = true
			for _, t := range winner.track {
				bestSpots[t.position] = true
			}
		}
	}
	return bestScore, len(bestSpots)
}
