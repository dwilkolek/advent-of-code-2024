package day13

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day 13, part 1: %d", solve(0, true))
}

func Part2() {
	logger.Printf("Day 13, part 2: %d", solve(10000000000000, false))
}

func solve(offset int, hasClickLimit bool) int {
	prizes := getPrizes(offset)

	sum := 0
	for _, p := range prizes {
		cost := findCheapestWin(p, hasClickLimit)
		if cost != math.MaxInt {
			sum += cost
		}
	}
	return sum
}

func findCheapestWin(p prize, hasClickLimit bool) int {
	var best = state{
		coordinates: coordinates{
			x: 0,
			y: 0,
		},
		cacheKey: cacheKey{
			a: 0,
			b: 0,
		},
		cost: math.MaxInt,
		isOk: true,
	}
	maxAClicks := int(math.Min(float64(p.prize.y/p.a.y), float64(p.prize.x/p.a.x)))
	logger.Printf("max A clicks: %d", maxAClicks)
	for aClicks := maxAClicks; aClicks > 0; aClicks-- {
		logger.Printf("max A clicks left: %d", aClicks)
		afterA := click(state{
			coordinates: coordinates{
				x: 0,
				y: 0,
			},
			cacheKey: cacheKey{
				a: 0,
				b: 0,
			},
			cost: 0,
			isOk: true,
		}, p.a, cacheKey{
			a: aClicks,
			b: 0,
		}, aClicks)
		bClicks := int(math.Min(float64((p.prize.y-afterA.coordinates.y)/p.b.y), float64((p.prize.x-afterA.coordinates.x)/p.b.x)))
		s := click(afterA, p.b, cacheKey{
			a: aClicks,
			b: bClicks,
		}, bClicks)

		if hasClickLimit && (s.cacheKey.a > 100 || s.cacheKey.b > 100) {
			continue
		}
		if s.y > p.prize.y || s.x > p.prize.x {
			continue
		}
		if s.y == p.prize.y && s.x == p.prize.x {
			if s.cost < best.cost {
				logger.Printf("Found success %d", s.cost)
				best = s
			}
		}

	}

	//logger.Printf("a=%d b=%d cost=%d", best.cacheKey.a, best.cacheKey.b, best.cost)
	return best.cost
}

func nww(a int, b int) []int {
	var v []int
	i := 2
	for {
		if a == 0 && b == 0 {
			break
		}
		if a%i == 0 && b%i == 0 {
			a = a / i
			b = b / i
			v = append(v, i)
		} else {
			i++
		}
	}
	return v
}

func cachedLookDeeper(p prize, s state, c map[cacheKey]state, hasClickLimit bool) state {
	cState, ok := c[s.cacheKey]
	if ok {
		return cState
	}
	ns := lookDeeper(p, s, hasClickLimit)
	c[s.cacheKey] = ns
	return ns
}
func lookDeeper(p prize, s state, hasClickLimit bool) state {
	if hasClickLimit && (s.cacheKey.a > 100 || s.cacheKey.b > 100) {
		s.cost = math.MaxInt64
		return s
	}
	if s.y > p.prize.y || s.x > p.prize.x {
		s.isOk = false
		s.cost = math.MaxInt64
		return s
	}
	if s.y == p.prize.y && s.x == p.prize.x {
		return s
	}

	aClick := click(s, p.a, cacheKey{
		a: s.cacheKey.a + 1,
		b: s.cacheKey.b + 0,
	}, 1)
	bClick := click(s, p.b, cacheKey{
		a: s.cacheKey.a + 0,
		b: s.cacheKey.b + 1,
	}, 1)

	if aClick.cost < bClick.cost {
		return aClick
	} else {
		return bClick
	}
}

func click(s state, btn button, cacheKey cacheKey, count int) state {
	return state{
		coordinates: coordinates{
			x: s.coordinates.x + count*btn.x,
			y: s.coordinates.y + count*btn.y,
		},
		cost:     s.cost + (count * btn.cost),
		cacheKey: cacheKey,
		isOk:     true,
	}
}

type state struct {
	coordinates
	cost int
	cacheKey
	isOk bool
}

func getPrizes(offset int) []prize {
	file, _ := os.Open("day13/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prizes []prize
	prizeInProgress := prize{}
	doA := true
	doB := false
	doPrize := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if doA {
			coords := strings.Split(strings.TrimPrefix(line, "Button A: "), ", ")
			xa, _ := strconv.Atoi(strings.TrimPrefix(coords[0], "X"))
			ya, _ := strconv.Atoi(strings.TrimPrefix(coords[1], "Y"))
			prizeInProgress.a = button{coordinates: coordinates{x: xa, y: ya}, cost: 3}
			doB = true
			doA = false
		} else if doB {
			coords := strings.Split(strings.TrimPrefix(line, "Button B: "), ", ")
			xa, _ := strconv.Atoi(strings.TrimPrefix(coords[0], "X"))
			ya, _ := strconv.Atoi(strings.TrimPrefix(coords[1], "Y"))
			prizeInProgress.b = button{coordinates: coordinates{x: xa, y: ya}, cost: 1}
			doPrize = true
			doB = false
		} else if doPrize {
			coords := strings.Split(strings.TrimPrefix(line, "Prize: "), ", ")
			xa, _ := strconv.Atoi(strings.TrimPrefix(coords[0], "X="))
			ya, _ := strconv.Atoi(strings.TrimPrefix(coords[1], "Y="))
			prizeInProgress.prize = coordinates{x: xa + offset, y: ya + offset}
			doPrize = true
			prizes = append(prizes, prizeInProgress)
			prizeInProgress = prize{}
			doA = true
			doPrize = false
		}
	}
	return prizes
}

type coordinates struct{ x, y int }
type button struct {
	coordinates
	cost int
}

type prize struct {
	a     button
	b     button
	prize coordinates
}

type cacheKey struct {
	a int
	b int
}
