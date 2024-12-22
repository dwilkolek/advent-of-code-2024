package day22

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var logger = log.Default()

func Part1() {
	secrets := parse()
	var sum int64 = 0
	for _, secret := range secrets {
		sum += nextNSecret(secret, 2000)
	}
	logger.Printf("Day 22, part 1: %d", sum)
}

func fillSequences(secret int64, popularSequences map[string]int64) {
	lagging4 := []int64{}
	lastSecret := secret
	lastPrice := secret % 10
	seen := map[string]bool{}
	for i := 1; i <= 2000; i++ {
		newSecret := nextSecret(lastSecret)
		price := newSecret % 10
		diff := price - lastPrice
		lagging4 = append(lagging4, diff)
		if len(lagging4) > 4 {
			lagging4 = lagging4[1:]
		}
		if len(lagging4) == 4 {
			key := fmt.Sprintf("%d,%d,%d,%d", lagging4[0], lagging4[1], lagging4[2], lagging4[3])
			if _, ok := seen[key]; !ok {
				seen[key] = true
				popularSequences[key] += price
			}
		}
		lastPrice = price
		lastSecret = newSecret
	}
}

func Part2() {
	secrets := parse()
	popularSequences := map[string]int64{}
	for _, secret := range secrets {
		fillSequences(secret, popularSequences)
	}
	bestOutcome := int64(0)
	for _, v := range popularSequences {
		bestOutcome = max(bestOutcome, v)
	}

	logger.Printf("Day 22, part 2: %d", bestOutcome)
}

func nextSecret(secret int64) int64 {
	number := secret * 64
	secret = mixAndPrune(secret, number)

	number = secret / 32
	secret = mixAndPrune(secret, number)

	number = secret * 2048
	secret = mixAndPrune(secret, number)
	return secret
}

func nextNSecret(secret int64, n int) int64 {
	for i := 0; i < n; i++ {
		secret = nextSecret(secret)
	}
	return secret
}

func mixAndPrune(secret, number int64) int64 {
	secret = mix(secret, number)
	secret = prune(secret)
	return secret
}
func mix(secret, b int64) int64 {
	return secret ^ b
}
func prune(secret int64) int64 {
	return secret % 16777216
}
func parse() []int64 {
	file, _ := os.Open("day22/input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)

	var lines []int64

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			s, _ := strconv.ParseInt(line, 10, 64)
			lines = append(lines, s)
		}
	}

	return lines
}
