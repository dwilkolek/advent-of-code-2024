package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var logger = log.Default()

var sizeX = 101
var sizeY = 103

func Part1() {
	rs := parse()
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	// q1|q2
	// -----
	// q3|q4
	for sec := 0; sec < 100; sec++ {
		for _, r := range rs {
			move(r, sizeX, sizeY)
		}

		printer(rs, sizeX, sizeY)
	}

	for _, r := range rs {
		if r.x == sizeX/2 {
			continue
		}
		if r.y == sizeY/2 {
			continue
		}
		if r.y < sizeY/2 {
			if r.x < sizeX/2 {
				q1++
			} else {
				q2++
			}
		} else {
			if r.x < sizeX/2 {
				q3++
			} else {
				q4++
			}
		}

	}

	logger.Printf("Day 14, part 1: %d", q1*q2*q3*q4)
}

func Part2() {
	rs := parse()

	sec := 0
	for {
		sec++
		for _, r := range rs {
			move(r, sizeX, sizeY)
		}

		m := mapper(rs, sizeX, sizeY, false, "*")
		//if sec%100 == 0 {
		//	logger.Printf("Sec: %d", sec)
		//	logger.Printf(m)
		//}
		if strings.Contains(m, "*******") {
			//logger.Printf(m)
			break
		}

	}
	logger.Printf("Day 14, part 2: %d", sec)
}

type robot struct{ x, y, vx, vy int }

func mapper(rs []*robot, sizeX, sizeY int, drawBlankSpot bool, robotChar string) string {
	m := ""
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if drawBlankSpot && (x == sizeX/2 || y == sizeY/2) {
				m += " "
				continue
			}
			rcount := 0
			for _, r := range rs {
				if r.x == x && r.y == y {
					rcount++
					continue
				}
			}
			if rcount == 0 {
				m += "."
			} else {
				if robotChar != "" {
					m += robotChar
				} else {
					if rcount > 9 {
						m += "*"
					} else {
						m += fmt.Sprintf("%d", rcount)
					}
				}
			}

		}
		m += "\n"
	}
	return m
}
func printer(rs []*robot, sizeX, sizeY int) {
	mapper(rs, sizeX, sizeY, true, "")
}

func move(r *robot, sizeX, sizeY int) {
	r.x += r.vx
	r.y += r.vy
	outsideBroken := true
	for outsideBroken {
		outsideBroken = false
		if r.x < 0 {
			r.x = sizeX + r.x
			outsideBroken = true
		}
		if r.x >= sizeX {
			r.x = r.x - sizeX
			outsideBroken = true
		}

		if r.y < 0 {
			r.y = sizeY + r.y
			outsideBroken = true
		}
		if r.y >= sizeY {
			r.y = r.y - sizeY
			outsideBroken = true
		}
	}
}

func parse() []*robot {
	file, _ := os.Open("day14/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//result := 0
	robots := make([]*robot, 0)
	for scanner.Scan() {
		line := scanner.Text()
		//"p=9,5 v=-3,-3"
		r := robot{
			x:  0,
			y:  0,
			vx: 0,
			vy: 0,
		}
		_, err := fmt.Fscanf(strings.NewReader(line), "p=%d,%d v=%d,%d", &r.x, &r.y, &r.vx, &r.vy)
		if err != nil {
			panic(err)
		}
		robots = append(robots, &r)
	}
	return robots
}
