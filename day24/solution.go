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
	s := defaultCase()
	logger.Printf("Day 24, part 1: %d", zValue(s))
}

func defaultCase() state {
	return parse(make(map[string]string)).eval()
}

func Part2() {
	swaps := findPossibleSwaps(defaultCase())
	for _, replacement := range swaps {
		ok, ans := testAndVerify(replacement)
		if ok {
			logger.Printf("Day 24, part 2: %s", ans)
			return
		}
	}
}

func testAndVerify(replacementPairs map[string]string) (bool, string) {
	replacements := map[string]string{}
	answer := []string{}
	for k, v := range replacementPairs {
		replacements[k] = v
		replacements[v] = k
		answer = append(answer, k)
		answer = append(answer, v)
	}

	slices.Sort(answer)

	return parse(replacements).eval().verify(), strings.Join(answer, ",")
}

type op struct {
	a, b string
	op   string
	out  string
}

func (o op) execute(s state) {
	if o.op == "AND" {
		s.gates[o.out] = s.gates[o.b] & s.gates[o.a]
	}
	if o.op == "OR" {
		s.gates[o.out] = s.gates[o.b] | s.gates[o.a]
	}

	if o.op == "XOR" {
		s.gates[o.out] = s.gates[o.b] ^ s.gates[o.a]
	}
}

type state struct {
	gates       map[string]int
	outputsDef  []op
	allZs       []string
	connections map[string][]string
}

func parse(replacements map[string]string) state {

	var gates = map[string]int{}
	var outputsDef = []op{}
	var allZs = []string{}
	var connections = map[string][]string{}

	data, _ := os.ReadFile("day24/input.txt")
	parts := strings.Split(string(data), "\n\n")

	for _, line := range strings.Split(parts[0], "\n") {
		def := strings.Split(line, ": ")
		gates[def[0]], _ = strconv.Atoi(def[1])
	}

	for _, line := range strings.Split(parts[1], "\n") {
		var gate1, gate2, operation, output string
		_, err := fmt.Sscanf(line, "%s %s %s -> %s", &gate1, &operation, &gate2, &output)
		if err != nil {
			panic(err)
		}
		if gate2O, ok := replacements[output]; ok {
			output = gate2O
		}
		outputsDef = append(outputsDef, op{gate1, gate2, operation, output})

		connections[gate1] = append(connections[gate1], output)
		connections[gate2] = append(connections[gate2], output)

		if strings.HasPrefix(output, "z") {
			allZs = append(allZs, output)
		}
		if _, ok := gates[output]; !ok {
			gates[output] = -1
		}
	}

	slices.Sort(allZs)

	s := state{
		gates:       gates,
		outputsDef:  outputsDef,
		allZs:       allZs,
		connections: connections,
	}

	return s

}

func (s state) eval() state {
	for evalNext(s) {
	}
	return s
}
func zValue(s state) int {

	v := 0
	for i, z := range s.allZs {
		v += s.gates[z] * int(math.Pow(2.0, float64(i)))
	}

	return v
}

func (s state) verify() bool {
	var xNum, yNum, zNum []int
	for _, z := range s.allZs {
		x := strings.ReplaceAll(z, "z", "x")
		y := strings.ReplaceAll(z, "z", "y")
		zNum = append(zNum, s.gates[z])
		if z == "z45" {
			continue
		}
		xNum = append(xNum, s.gates[x])
		yNum = append(yNum, s.gates[y])
	}
	return arrToNum(xNum)+arrToNum(yNum) == arrToNum(zNum)

}

func evalNext(s state) bool {
	wasUpdate := false
	for _, ope := range s.outputsDef {
		aSig, aIsReady := s.gates[ope.a]
		bSig, bIsReady := s.gates[ope.b]
		outSig, outIsReady := s.gates[ope.out]
		if (outSig == -1 || !outIsReady) && aIsReady && bIsReady && aSig != -1 && bSig != -1 {
			ope.execute(s)
			wasUpdate = true
		}
	}
	return wasUpdate
}

func arrToNum(c []int) int64 {
	var cn int64 = 0
	for i, ci := range c {
		cn += int64(ci) * int64(math.Pow(2.0, float64(i)))
	}
	return cn
}

func findPossibleSwaps(s state) []map[string]string {
	all := []map[string]string{}
	wrong := []string{}
	for _, op := range s.outputsDef {
		if op.out == "z45" {
			continue
		}
		if strings.HasPrefix(op.out, "z") && op.op != "XOR" {
			wrong = append(wrong, op.out)
			continue
		}
		if op.op == "XOR" && !hasKnownPrefix(op.a) && !hasKnownPrefix(op.b) && !strings.HasPrefix(op.out, "z") {
			wrong = append(wrong, op.out)
		}
	}
	for i := 0; i < len(s.allZs)-1; i++ {
		s1, s2 := xyXORhasANDLeadingToZ(s, fmt.Sprintf("z%02d", i))
		if s1 != "" {
			//might not work for all inputs, but seems like they are sorted...
			all = append(all, map[string]string{
				s1:       s2,
				wrong[0]: wrong[1],
				wrong[2]: wrong[3],
				wrong[4]: wrong[5],
			})
		}
	}
	return all
}

func xyXORhasANDLeadingToZ(s state, z string) (string, string) {
	x := strings.ReplaceAll(z, "z", "x")
	y := strings.ReplaceAll(z, "z", "y")
	xor := s.findConn(x, y, "XOR")
	if strings.HasPrefix(xor, "z") {
		return "", ""
	}
	maybeZ := s.findConn(xor, "?", "XOR")
	if !strings.HasPrefix(maybeZ, "z") {
		return xor, s.findConn(x, y, "AND")
	}
	return "", ""
}

func hasKnownPrefix(s string) bool {
	return strings.HasPrefix(s, "x") || strings.HasPrefix(s, "y") || strings.HasPrefix(s, "z")
}

func (s state) findConn(a, b, op string) string {
	for _, o := range s.outputsDef {
		if o.op == op {
			if b == "?" && o.a == a || o.b == a {
				return o.out
			} else if (o.a == a || o.a == b) && (o.b == a || o.b == b) {
				return o.out
			}
		}
	}
	return ""
}
