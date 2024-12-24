package day24

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var logger = log.Default()

func Part1() {
	logger.Printf("Day 24, part 1: %d", solve())
}

func Part2() {
	logger.Printf("Day 24, part 2: %d", solve())
}

type op struct {
	a, b string
	op   string
	out  string
}

func (o op) execute() {
	if o.op == "AND" {
		gates[o.out] = gates[o.b] & gates[o.a]
	}
	if o.op == "OR" {
		gates[o.out] = gates[o.b] | gates[o.a]
	}

	if o.op == "XOR" {
		gates[o.out] = gates[o.b] ^ gates[o.a]
	}
}

var gates = map[string]int{}
var outputSources = map[string]op{}
var outputsDef = []op{}
var allZs = []string{}

func solve() int {
	data, _ := os.ReadFile("day24/input.txt")
	parts := strings.Split(string(data), "\n\n")

	for _, line := range strings.Split(parts[0], "\n") {
		def := strings.Split(line, ": ")
		gates[def[0]], _ = strconv.Atoi(def[1])
	}
	//var operations []op

	for _, line := range strings.Split(parts[1], "\n") {
		var gate1, gate2, operation, output string
		// x00 AND y00 -> z00
		_, err := fmt.Sscanf(line, "%s %s %s -> %s", &gate1, &operation, &gate2, &output)
		if err != nil {
			return -1
		}
		outputSources[output] = op{gate1, gate2, operation, output}
		outputsDef = append(outputsDef, op{gate1, gate2, operation, output})
		//operations = append(operations, op{gate1, gate2, operation, output})
		if strings.HasPrefix(output, "z") {
			allZs = append(allZs, output)
		}
		if _, ok := gates[output]; !ok {
			gates[output] = -1
		}
	}

	slices.Sort(allZs)

	for evalNext(outputsDef) {
		if allZDone() {
			break
		}
	}

	v := 0
	for i, z := range allZs {
		v += gates[z] * int(math.Pow(2.0, float64(i)))
	}

	return v
}

func allZDone() bool {
	for _, z := range allZs {
		if v, ok := gates[z]; !ok || v == -1 {
			return false
		}
	}
	return true
}

func evalNext(ops []op) bool {
	wasUpdate := false
	for _, ope := range ops {
		aSig, aIsReady := gates[ope.a]
		bSig, bIsReady := gates[ope.b]
		outSig, outIsReady := gates[ope.out]
		if (outSig == -1 || !outIsReady) && aIsReady && bIsReady && aSig != -1 && bSig != -1 {
			ope.execute()
			wasUpdate = true
		}
	}
	return wasUpdate
}
