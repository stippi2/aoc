package day17

import (
	"aoc/2024/go/lib"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Computer struct {
	A                  int
	B                  int
	C                  int
	instructionPointer int
	program            []int
	output             string
}

func (c *Computer) comboOperand(value int) int {
	switch value {
	case 0, 1, 2, 3:
		return value
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	default:
		panic(fmt.Sprintf("combo operand with value %v", value))
	}
}

func (c *Computer) runInstruction() bool {
	if c.instructionPointer > len(c.program)-2 {
		return false
	}
	opCode := c.program[c.instructionPointer]
	operand := c.program[c.instructionPointer+1]
	c.instructionPointer += 2

	switch opCode {
	case 0: // adv
		fmt.Printf("adv %v\n", c.comboOperand(operand))
		c.A = int(float64(c.A) / math.Pow(2, float64(c.comboOperand(operand))))
	case 1: // bxl
		fmt.Printf("bxl %v\n", operand)
		c.B = c.B ^ operand
	case 2: // bst
		fmt.Printf("bst %v\n", operand)
		c.B = c.comboOperand(operand) % 8
	case 3: // jnz
		if c.A != 0 {
			fmt.Printf("jnz %v\n", operand)
			c.instructionPointer = operand
		} else {
			fmt.Printf("nop %v\n", operand)
		}
	case 4: // bxc
		fmt.Printf("bxc\n")
		c.B = c.B ^ c.C
	case 5: // out
		fmt.Printf("out %v\n", c.comboOperand(operand))
		if len(c.output) > 0 {
			c.output = fmt.Sprintf("%s,%v", c.output, c.comboOperand(operand)%8)
		} else {
			c.output = fmt.Sprintf("%v", c.comboOperand(operand)%8)
		}
	case 6: // bdv
		fmt.Printf("bdv %v\n", c.comboOperand(operand))
		c.B = int(float64(c.A) / math.Pow(2, float64(c.comboOperand(operand))))
	case 7: // cdv
		fmt.Printf("cdv %v\n", c.comboOperand(operand))
		c.C = int(float64(c.A) / math.Pow(2, float64(c.comboOperand(operand))))
	}

	return true
}

func (c *Computer) runProgram() {
	for c.runInstruction() {
	}
}

func parseInput(input string) *Computer {
	computer := &Computer{}
	parts := strings.Split(input, "\n\n")
	matches, _ := fmt.Sscanf(parts[0], "Register A: %d\nRegister B: %d\nRegister C: %d", &computer.A, &computer.B, &computer.C)
	if matches != 3 {
		panic("failed to parse registers")
	}
	program := strings.TrimPrefix(parts[1], "Program: ")
	for _, codeString := range strings.Split(program, ",") {
		code, _ := strconv.Atoi(codeString)
		computer.program = append(computer.program, code)
	}
	return computer
}

func Part1() any {
	input, _ := lib.ReadInput(17)
	computer := parseInput(input)
	computer.runProgram()
	return computer.output
}

func Part2() any {
	return "Not implemented"
}
