package day21

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day X, part 1: %d", solve())
}

func Part2() {
	logger.Printf("Day X, part 2: %d", solve())
}

type coord struct {
	x, y int
}

var numpad = map[coord]string{
	{0, 0}: "7",
	{1, 0}: "8",
	{2, 0}: "9",
	{0, 1}: "4",
	{1, 1}: "5",
	{2, 1}: "6",
	{0, 2}: "1",
	{1, 2}: "2",
	{2, 2}: "3",
	{1, 3}: "0",
	{2, 3}: "A",
}
var dirBoard = map[coord]string{
	{1, 0}: "^",
	{2, 0}: "A",
	{0, 1}: "<",
	{1, 1}: "v",
	{2, 1}: ">",
}

var moves = map[string]coord{
	"v": {x: 0, y: 1},
	"^": {x: 0, y: -1},
	"<": {x: -1, y: 0},
	">": {x: 1, y: 0},
}

func findCoordOfKey(key string, keybaord map[coord]string) (coord, error) {
	for kKeyCoord, kKey := range keybaord {
		if key == kKey {
			return kKeyCoord, nil
		}
	}
	return coord{}, errors.New("key not found")
}
func findStepsToKey(fromKey, toKey string, keyboard map[coord]string, history []string, level int, seen map[string]bool) [][]string {
	if _, ok := seen[strings.Join(history, ",")]; ok {
		return [][]string{}
	}
	seen[strings.Join(history, ",")] = true
	hitCount := 0
	fromKeyCoord, err := findCoordOfKey(fromKey, keyboard)
	if err != nil {
		panic(err)
	}
	toKeyCoord, err := findCoordOfKey(toKey, keyboard)
	if err != nil {
		panic(err)
	}

	if fromKeyCoord == toKeyCoord {
		return [][]string{
			append(history, "A"),
		}
	}

	possibleSolutions := [][]string{}
	for moveSign, move := range moves {
		newFromKeyCoord := coord{
			x: fromKeyCoord.x + move.x,
			y: fromKeyCoord.y + move.y,
		}
		if newKey, ok := keyboard[newFromKeyCoord]; ok {
			hitCount++
			oldDist := dist(fromKeyCoord, toKeyCoord)
			newDist := dist(newFromKeyCoord, toKeyCoord)
			if newDist < oldDist {
				solutions := findStepsToKey(newKey, toKey, keyboard, append(history, moveSign), level+1, seen)
				possibleSolutions = append(possibleSolutions, solutions...)
			}
		}
	}
	//if level == 0 {
	//	logger.Printf("Find key(%d) from %s to %s returned %d, dist=%d", level, fromKey, toKey, len(possibleSolutions), len(possibleSolutions[0])-1)
	//}
	return possibleSolutions
}

func dist(a, b coord) float64 {
	return math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
}

func solve() int {
	file, _ := os.Open("day21/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalComplexity := 0
	for scanner.Scan() {
		line := scanner.Text()
		buttons := strings.Split(line, "")
		v := stage2(buttons, numpad)
		logger.Printf("Stage 1 combinations : %d,", len(v))

		var v2 [][]string
		for _, vi := range v {
			v2S := stage2(vi, dirBoard)
			logger.Printf("For %s got %d", strings.Join(vi, ""), len(v2S))
			for _, v2i := range v2S {
				v2 = append(v2, v2i)
			}
		}

		logger.Printf("Stage 2 combinations : %d, ", len(v2))
		//logger.Printf("Stage 2 %d: %v", len(v2), v2)
		var v3 [][]string
		for _, vi := range v2 {
			for _, v3i := range stage2(vi, dirBoard) {
				v3 = append(v3, v3i)
			}
		}

		//logger.Printf("Stage 3 comintation : %d", len(v3))

		logger.Printf("Stage 3 combinations : %d,", len(v3))
		bestV3 := math.MaxInt
		for _, v3i := range v3 {
			if len(v3i) <= bestV3 {
				bestV3 = len(v3i)
			}
		}

		logger.Printf("Stage 3: %d", bestV3)
		code, _ := strconv.Atoi(strings.Join(buttons[0:3], ""))
		totalComplexity += code * bestV3
	}
	return totalComplexity
}

func findBest(list [][]string) [][]string {
	bestSize := math.MaxInt
	for _, item := range list {
		if len(item) < bestSize {
			bestSize = len(item)
		}
	}
	//2024/12/21 18:49:03 Stage 1 12: [< A ^ ^ ^ A v A > v v A]
	//2024/12/21 18:49:03 Stage 2 28: [v < < A > > ^ A < A A A > A < v A > ^ A v A < A A > ^ A]
	//2024/12/21 18:49:03 Stage 3 70: [v < A < A A > ^ > A v A A ^ < A > A < v < A > > ^ A A A v A ^ A v < < A > A > ^ A v A < ^ A > A < v A ^ > A v < < A > > ^ A A v A < ^ A > A]

	//2024/12/21 18:49:27 Stage 1 12: [< A ^ ^ ^ A v A v v > A]
	//2024/12/21 18:49:27 Stage 2 28: [v < < A > > ^ A < A A A > A v < A ^ > A < v A A > A ^ A]
	//2024/12/21 18:49:27 Stage 3 66: [v < A < A A > ^ > A v A A ^ < A > A v < < A > > ^ A A A v A ^ A v < A < A > > ^ A < A > v A ^ A < v < A > A > ^ A A v A ^ A < A > A]
	best := make([][]string, 0)
	for _, item := range list {
		if len(item) == bestSize {
			best = append(best, item)
		}
	}

	return best
}

func stage1(buttons []string) [][]string {
	buttons = append([]string{"A"}, buttons...)
	atButton := 0
	solutions := map[int][][]string{}
	for atButton = 0; atButton < len(buttons)-1; atButton++ {
		v := findStepsToKey(buttons[atButton], buttons[atButton+1], numpad, []string{}, 0, map[string]bool{})
		if prev, prevOk := solutions[atButton-1]; prevOk {
			solutions[atButton] = [][]string{}
			for _, vi := range v {
				for _, previ := range prev {
					solutions[atButton] = append(solutions[atButton], append(previ, vi...))
				}
			}
		} else {
			solutions[atButton] = v
		}

	}

	return solutions[atButton-1]
}
func stage2(buttons []string, kb map[coord]string) [][]string {
	solutions := map[int][][]string{}
	buttons = append([]string{"A"}, buttons...)

	for atButton := 0; atButton < len(buttons)-1; atButton++ {
		v := findStepsToKey(buttons[atButton], buttons[atButton+1], kb, []string{}, 0, map[string]bool{})
		if atButton == 0 {
			solutions[atButton] = v
			continue
		}
		//prev, _ := solutions[atButton-1]
		for _, vi := range v {
			solutions[atButton] = append(solutions[atButton], vi)
		}

		//prev, prevOk := solutions[atButton-1]
		//if prevOk {
		//	for _, vi := range v {
		//		for _, previ := range prev {
		//			solutions[atButton] = append(solutions[atButton], append(previ, vi...))
		//		}
		//	}
		//} else {
		//	logger.Printf("Creating new %d for l=%d", atButton, len(v))
		//	solutions[atButton] = v
		//}
	}

	finalSolutions := solutions[0]
	for atButton := 1; atButton < len(buttons)-1; atButton++ {
		newFinalSolutions := make([][]string, 0)
		for _, p := range finalSolutions {
			for _, n := range solutions[atButton] {
				//logger.Printf("Creating from %s -> %s", strings.Join(n, ""), strings.Join(append(p, n...), ""))
				nv := []string{}
				for _, pi := range p {
					nv = append(nv, pi)
				}
				for _, ni := range n {
					nv = append(nv, ni)
				}
				newFinalSolutions = append(newFinalSolutions, nv)
				//logger.Printf("Result %s", strings.Join(newFinalSolutions[len(newFinalSolutions)-1], ""))

			}
		}
		finalSolutions = newFinalSolutions
	}

	//best := math.MaxInt
	//for _, x := range solutions[len(buttons)-2] {
	//	if len(x) < best {
	//		best = len(x)
	//	}
	//}
	//if best < len(solutions[len(buttons)-2][0]) {
	//	panic("There is shorted but returning Longer")
	//}
	//logger.Printf("shortes: %d, but %d", best, len(solutions[atButton-1][0]))
	return finalSolutions
}
