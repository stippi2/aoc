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
