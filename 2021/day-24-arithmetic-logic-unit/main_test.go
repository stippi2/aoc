package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = `inp z
inp x
mul z 3
eql z x`

func Test_parseInput(t *testing.T) {
	alu, program := parseInput(example)
	alu.input = []int{1,3}
	for _, instruction := range program {
		instruction.Execute()
	}
	assert.Equal(t, 1, alu.z)
}

func Test_example(t *testing.T) {
	alu, program := parseInput(loadInput("puzzle-input.txt"))
	modelNumber := newModelNumber()
	modelNumber.setFrom("11111111111111")
	alu.input = modelNumber.value
	iteration := 0
	for {
		for _, instruction := range program {
			instruction.Execute()
		}
		if alu.z == 0 {
			break
		}
		iteration++
		fmt.Printf("model number %s: invalid, z:%v\n", modelNumber, alu.z)
		if iteration == 100000 {
			break
		}
		modelNumber.increment()
		alu.reset()
	}
	fmt.Printf("valid model number: %v\n", modelNumber)
}

func Test_mod(t *testing.T) {
	alu, program := parseInput(loadInput("puzzle-input-orig.txt"))
	modelNumber := newModelNumber()
	numbers := []string {
		"11111111111111",
		"22222222222222",
		"33333333333333",
		"44444444444444",
		"55555555555555",
		"66666666666666",
		"77777777777777",
		"88888888888888",
		"99999999999999",
	}
	type evolvingZ struct {
		x, y, z int
	}

	aluStates := make([]map[int]int, 15)
	aluStates[0] = make(map[int]int)
	aluStates[0][0] = 0

	for digit := 0; digit < 14; digit++ {
		aluStates[digit + 1] = make(map[int]int)
		for i, number := range numbers {
			replaced := 0
			for state, numberPath := range aluStates[digit] {
				// Only register "z" carries over from previous digits
				alu.x = 0
				alu.y = 0
				alu.z = state
				modelNumber.setFrom(number)
				alu.index = digit
				alu.input = modelNumber.value
				alu.runProgramToNextInput(digit, program)
				//fmt.Printf("digit %v, z: %v\n", digit, alu.z)
				oldNumberPath := aluStates[digit + 1][alu.z]
				newNumberPath := numberPath * 10 + i + 1
				if oldNumberPath == 0 || newNumberPath < oldNumberPath {
					aluStates[digit + 1][alu.z] = newNumberPath
					if oldNumberPath > 0 {
						replaced++
					}
				}
			}
			fmt.Printf("states at digit %v: %v (replaced: %v)\n", digit, len(aluStates[digit + 1]), replaced)
		}
	}
	fmt.Printf("lowest valid model number: %v\n", aluStates[14][0])
}

func Test_compareStepVersusAll(t *testing.T) {
	alu, program := parseInput(loadInput("puzzle-input-orig.txt"))
	modelNumber := newModelNumber()

	modelNumber.setFrom("13579246899999")
	alu.input = modelNumber.value
	for _, instruction := range program {
		instruction.Execute()
	}

	fmt.Printf("model number %s: invalid, z:%v\n", modelNumber, alu.z)

	z := 0
	for digit := 0; digit < 14; digit++ {
		// Only register "z" carries over from previous digits
		alu.x = 0
		alu.y = 0
		alu.z = z
		alu.index = digit
		alu.runProgramToNextInput(digit, program)
		z = alu.z
	}
	fmt.Printf("z by running steps individually: %v\n", z)
}
