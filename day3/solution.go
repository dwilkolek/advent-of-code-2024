package day3

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var logger = log.Default()

func Part2() {
	file, _ := os.Open("day3/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	enabled := true
	r := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|(don\'t\(\))|(do\(\))`)
	for scanner.Scan() {
		line := scanner.Text()

		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			if strings.HasPrefix(match, "mul") {
				if !enabled {
					continue
				}
				muls := strings.Split(string([]byte(match)[4:len(match)-1]), ",")
				a, _ := strconv.Atoi(muls[0])
				b, _ := strconv.Atoi(muls[1])
				sum += a * b
			} else if strings.HasPrefix(match, "don't") {
				enabled = false
			} else if strings.HasPrefix(match, "do(") {
				enabled = true
			} else {
				log.Panicf("unknown command? %s\n", match)
			}

		}
	}

	log.Printf("Day 3, part 2: %d", sum)
}
func Part1() {

	file, _ := os.Open("day3/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	r := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	for scanner.Scan() {
		line := scanner.Text()

		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			muls := strings.Split(string([]byte(match)[4:len(match)-1]), ",")
			a, _ := strconv.Atoi(muls[0])
			b, _ := strconv.Atoi(muls[1])
			sum += a * b
		}
	}

	log.Printf("Day 3, part 1: %d", sum)
}
