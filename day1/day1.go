package day1

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	file, _ := os.Open("day1/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var leftList []int
	var rightList []int
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		leftList = append(leftList, l)
		rightList = append(rightList, r)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)
	dist := 0
	for i, l := range leftList {
		r := rightList[i]
		if l < r {
			dist = dist + r - l
		} else {
			dist = dist + l - r
		}
	}

	log.Default().Printf("1: %d", dist)
}

func Part2() {
	file, _ := os.Open("day1/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var leftList []int
	var rightCounts map[int]int = make(map[int]int)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		leftList = append(leftList, l)
		count, found := rightCounts[r]
		if !found {
			rightCounts[r] = 1
		} else {
			rightCounts[r] = count + 1
		}
	}

	dist := 0
	for _, l := range leftList {
		count, found := rightCounts[l]
		if found {
			dist = dist + (l * count)
		}
	}

	log.Default().Printf("2: %d", dist)
}
