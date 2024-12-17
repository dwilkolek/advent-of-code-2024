package day17

import (
	"testing"
)

// If register C contains 9, the program 2,6 would set register B to 1.
func TestAoC1(t *testing.T) {
	resetProgram()
	registers["C"] = 9
	program = []int{2, 6}
	programLoop()
	if registers["B"] != 1 {
		t.Fatalf("Expected %d but was %d", 1, registers["B"])
	}
}

// If register A contains 10, the program 5,0,5,1,5,4 would output 0,1,2.
func TestAoC2(t *testing.T) {
	resetProgram()
	registers["A"] = 10
	program = []int{5, 0, 5, 1, 5, 4}
	programLoop()
	if output != "0,1,2" {
		t.Fatalf("Expected '%s' but was '%s'", "0,1,2", output)
	}
}

// If register A contains 2024, the program 0,1,5,4,3,0 would output 4,2,5,6,7,7,7,7,3,1,0 and leave 0 in register A.
func TestAoC3(t *testing.T) {
	resetProgram()
	registers["A"] = 2024
	program = []int{0, 1, 5, 4, 3, 0}
	programLoop()
	if output != "4,2,5,6,7,7,7,7,3,1,0" {
		t.Fatalf("Expected %s but was %s", "4,2,5,6,7,7,7,7,3,1,0", output)
	}
	if registers["A"] != 0 {
		t.Fatalf("Expected %d but was %d", 0, registers["A"])
	}
}

// If register B contains 29, the program 1,7 would set register B to 26.
func TestAoC4(t *testing.T) {
	resetProgram()
	registers["B"] = 29
	program = []int{1, 7}
	programLoop()
	if registers["B"] != 26 {
		t.Fatalf("Expected %d but was %d", 26, registers["B"])
	}
}

// If register B contains 2024 and register C contains 43690, the program 4,0 would set register B to 44354.
func TestAoC5(t *testing.T) {
	resetProgram()
	registers["B"] = 2024
	registers["C"] = 43690
	program = []int{4, 0}
	programLoop()
	if registers["B"] != 44354 {
		t.Fatalf("Expected %d but was %d", 44354, registers["B"])
	}
}
