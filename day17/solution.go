package day17

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var logger = log.Default()

func Part1() {
	parse()
	logger.Printf("Day 17, part 1: %s", solve())
}

func Part2() {
	parse()
	logger.Printf("Day 17, part 2: %d", solve())
}

var registers = map[string]int{}
var program []int
var instrPointer = 0
var output string

func parse() {
	registers = map[string]int{}
	program = make([]int, 0)
	data, _ := os.ReadFile("day17/input.txt")
	parts := bytes.Split(data, []byte("\n\n"))
	for _, part := range strings.Split(string(parts[0]), "\n") {
		rs := strings.Split(strings.TrimPrefix(part, "Register "), ": ")
		rv, _ := strconv.Atoi(rs[1])
		registers[rs[0]] = rv
	}
	for _, opcode := range strings.Split(strings.TrimPrefix(string(parts[1]), "Program: "), ",") {
		op, _ := strconv.Atoi(opcode)
		program = append(program, op)
	}
}

func comboValue(operand int) int {
	switch operand {
	case 0:
		return operand
	case 1:
		return operand
	case 2:
		return operand
	case 3:
		return operand
	case 4:
		return registers["A"]
	case 5:
		return registers["B"]
	case 6:
		return registers["C"]
	case 7:
		panic("should not show up in valid programs")
	default:
		panic("unknown")
	}
	return -1
}

func loadOpCode() bool {
	if instrPointer >= len(program) {
		return false
	}
	opCode := program[instrPointer]

	switch opCode {
	case 0:
		registers["A"] = registers["A"] / int(math.Pow(2.0, float64(comboValue(program[instrPointer+1]))))
		break
	case 1:
		registers["B"] = registers["B"] ^ program[instrPointer+1]
		break
	case 2:
		registers["B"] = comboValue(program[instrPointer+1]) % 8
		break
	case 3:
		if registers["A"] == 0 {
			break
		}
		instrPointer = program[instrPointer+1]
		return true
	case 4:
		registers["B"] = registers["B"] ^ registers["C"]
		break
	case 5:
		nVal := comboValue(program[instrPointer+1]) % 8
		if len(output) == 0 {
			output = fmt.Sprintf("%d", nVal)
		} else {
			output += fmt.Sprintf(",%d", nVal)
		}
		break
	case 6:
		registers["B"] = registers["A"] / int(math.Pow(2.0, float64(comboValue(program[instrPointer+1]))))
		break
	case 7:
		registers["C"] = registers["A"] / int(math.Pow(2.0, float64(comboValue(program[instrPointer+1]))))
		break
	}
	instrPointer += 2
	return true
}

func solve() string {
	parse()
	programLoop()
	return output
}

func resetProgram() {
	registers = map[string]int{
		"A": 0,
		"C": 0,
		"B": 0,
	}
	program = make([]int, 0)
	instrPointer = 0
	output = ""
}

func programLoop() {
	doNext := true
	for doNext {
		doNext = loadOpCode()
	}
}
