package day4

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var logger = log.Default()

func Part1() {

	file, _ := os.Open("day4/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	offset := 4
	size := 140 + 2*offset
	lines := make([][]string, size)
	for i := 0; i < size; i++ {
		lines[i] = make([]string, size)
	}
	lineIdx := offset
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		for i, c := range chars {
			lines[lineIdx][i+offset] = c
		}
		lineIdx += 1
	}
	sum := 0
	for y, line := range lines {
		if y < offset {
			continue
		}
		for x, _ := range line {
			if x < offset {
				continue
			}
			if t(y, x, lines) {
				sum += 1
			}
			if b(y, x, lines) {
				sum += 1
			}
			if l(y, x, lines) {
				sum += 1
			}
			if r(y, x, lines) {
				sum += 1
			}
			if tl(y, x, lines) {
				sum += 1
			}
			if tr(y, x, lines) {
				sum += 1
			}
			if bl(y, x, lines) {
				sum += 1
			}
			if br(y, x, lines) {
				sum += 1
			}
		}
	}
	log.Printf("Day 4, part 1: %d", sum)
}

func r(y int, x int, lines [][]string) bool {
	return lines[y][x] == "X" && lines[y][x+1] == "M" && lines[y][x+2] == "A" && lines[y][x+3] == "S"
}
func l(y int, x int, lines [][]string) bool {
	return lines[y][x] == "X" && lines[y][x-1] == "M" && lines[y][x-2] == "A" && lines[y][x-3] == "S"
}
func t(y int, x int, lines [][]string) bool {
	return lines[y][x] == "X" && lines[y-1][x] == "M" && lines[y-2][x] == "A" && lines[y-3][x] == "S"
}
func b(y int, x int, lines [][]string) bool {
	return lines[y][x] == "X" && lines[y+1][x] == "M" && lines[y+2][x] == "A" && lines[y+3][x] == "S"
}
func tr(y int, x int, lines [][]string) bool {
	return lines[y][x] == "X" && lines[y-1][x+1] == "M" && lines[y-2][x+2] == "A" && lines[y-3][x+3] == "S"
}
func tl(y int, x int, lines [][]string) bool {
	return lines[y][x] == "X" && lines[y-1][x-1] == "M" && lines[y-2][x-2] == "A" && lines[y-3][x-3] == "S"
}
func br(y int, x int, lines [][]string) bool {
	return lines[y][x] == "X" && lines[y+1][x+1] == "M" && lines[y+2][x+2] == "A" && lines[y+3][x+3] == "S"
}
func bl(y int, x int, lines [][]string) bool {
	return lines[y][x] == "X" && lines[y+1][x-1] == "M" && lines[y+2][x-2] == "A" && lines[y+3][x-3] == "S"
}

func Part2() {

	file, _ := os.Open("day4/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	offset := 4
	size := 140 + 2*offset
	lines := make([][]string, size)
	for i := 0; i < size; i++ {
		lines[i] = make([]string, size)
	}
	lineIdx := offset
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		for i, c := range chars {
			lines[lineIdx][i+offset] = c
		}
		lineIdx += 1
	}
	sum := 0
	for y, line := range lines {
		if y < offset {
			continue
		}
		for x, _ := range line {
			if x < offset {
				continue
			}
			if xmas1(y, x, lines) {
				sum += 1
			}
			if xmas2(y, x, lines) {
				sum += 1
			}
			if xmas3(y, x, lines) {
				sum += 1
			}
			if xmas4(y, x, lines) {
				sum += 1
			}
		}
	}
	log.Printf("Day 4, part 2: %d", sum)
}
func xmas1(y int, x int, lines [][]string) bool {
	// M . S
	// . A .
	// M . S
	return lines[y][x] == "A" && lines[y-1][x-1] == "M" && lines[y+1][x+1] == "S" && lines[y+1][x-1] == "M" && lines[y-1][x+1] == "S"
}
func xmas2(y int, x int, lines [][]string) bool {
	// S . M
	// . A .
	// S . M
	return lines[y][x] == "A" && lines[y-1][x-1] == "S" && lines[y+1][x+1] == "M" && lines[y+1][x-1] == "S" && lines[y-1][x+1] == "M"
}
func xmas3(y int, x int, lines [][]string) bool {
	// M . M
	// . A .
	// S . S
	return lines[y][x] == "A" && lines[y-1][x-1] == "M" && lines[y+1][x+1] == "S" && lines[y-1][x+1] == "M" && lines[y+1][x-1] == "S"
}
func xmas4(y int, x int, lines [][]string) bool {
	// S . S
	// . A .
	// M . M
	return lines[y][x] == "A" && lines[y-1][x-1] == "S" && lines[y+1][x+1] == "M" && lines[y-1][x+1] == "S" && lines[y+1][x-1] == "M"
}
