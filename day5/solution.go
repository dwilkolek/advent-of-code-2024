package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var logger = log.Default()

func Part1() {
	input := read()
	result := 0
	for _, update := range input.updates {
		pageMap := make(map[int]int)
		for p, page := range update {
			pageMap[page] = p
		}
		valid := true
		for _, rule := range input.rules {
			a, ok := pageMap[rule.a]
			if !ok {
				continue
			}
			b, ok := pageMap[rule.b]
			if !ok {
				continue
			}

			if a > b {
				valid = false
			}
		}

		if valid {
			middle := update[len(update)/2]
			result += middle
		}
	}

	log.Printf("Day 5, part 1: %d", result)
}

func Part2() {
	input := read()
	result := 0
	for _, update := range input.updates {
		pageMap := make(map[int]int)
		for p, page := range update {
			pageMap[page] = p
		}
		addToResult := false
		revalidate := true
		for revalidate {
			revalidate = false
			for _, rule := range input.rules {
				aIdx, ok := pageMap[rule.a]
				if !ok {
					continue
				}
				bIdx, ok := pageMap[rule.b]
				if !ok {
					continue
				}

				if aIdx > bIdx {
					addToResult = true
					revalidate = true
					pageMap[rule.a] = bIdx
					pageMap[rule.b] = aIdx

					update[aIdx] = rule.b
					update[bIdx] = rule.a
				}
			}
		}

		if addToResult {
			middle := update[len(update)/2]
			result += middle
		}
	}

	log.Printf("Day 5, part 2: %d", result)
}

type rule struct {
	a int
	b int
}
type update = []int

type input struct {
	rules   []rule
	updates []update
}

func read() input {
	file, _ := os.Open("day5/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	readsRules := true
	input := input{
		rules:   make([]rule, 0),
		updates: make([]update, 0),
	}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readsRules = false
			continue
		}

		if readsRules {
			parts := strings.Split(line, "|")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			input.rules = append(input.rules, rule{
				a: a, b: b,
			})
		} else {
			parts := strings.Split(line, ",")
			pages := make(update, len(parts))
			for i, part := range parts {
				pages[i], _ = strconv.Atoi(part)
			}
			input.updates = append(input.updates, pages)
		}

	}

	return input
}
