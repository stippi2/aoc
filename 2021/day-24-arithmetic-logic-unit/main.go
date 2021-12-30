package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type ALU struct {
	w, x, y, z int
	input []int
	index int
}

func (alu *ALU) Register(name string) *int {
	switch name {
	case "w":
		return &alu.w
	case "x":
		return &alu.x
	case "y":
		return &alu.y
	case "z":
		return &alu.z
	}
	return nil
}

func (alu *ALU) Read() int {
	alu.index++
	return alu.input[alu.index-1]
}

func (alu *ALU) reset() {
	alu.index = 0
	alu.w = 0
	alu.x = 0
	alu.y = 0
	alu.z = 0
}

type Register struct {
	register *int
}

func (r *Register) Read() int {
	return *r.register
}

type Number struct {
	value int
}

func (n *Number) Read() int {
	return n.value
}

type Input interface {
	Read() int
}

type Instruction interface {
	Execute()
}

type INP struct {
	register *int
	input    Input
}

func (i *INP) Execute() {
	*i.register = i.input.Read()
}

type ADD struct {
	register *int
	input    Input
}

func (i *ADD) Execute() {
	*i.register += i.input.Read()
}

type MUL struct {
	register *int
	input    Input
}

func (i *MUL) Execute() {
	*i.register *= i.input.Read()
}

type DIV struct {
	register *int
	input    Input
}

func (i *DIV) Execute() {
	*i.register /= i.input.Read()
}

type MOD struct {
	register *int
	input    Input
}

func (i *MOD) Execute() {
	*i.register %= i.input.Read()
}

type EQL struct {
	register *int
	input    Input
}

func (i *EQL) Execute() {
	if *i.register == i.input.Read() {
		*i.register = 1
	} else {
		*i.register = 0
	}
}

// splitInstructions splits the program into 14 chunks that each start with an INP instruction.
// It is assumed that the program contains exactly 14 INP instructions and that it starts with an INP.
func splitInstructions(program []Instruction) [][]Instruction {
	programChunks := make([][]Instruction, 14)
	inputsSeen := -1
	for _, instruction := range program {
		if _, ok := instruction.(*INP); ok {
			inputsSeen++
		}
		programChunks[inputsSeen] = append(programChunks[inputsSeen], instruction)
	}
	return programChunks
}

// min is unusual in the sense that it ignores the zero value for a (0)
func min(a, b int) int {
	if a == 0 || b < a {
		return b
	}
	return a
}

// max is unusual in the sense that it ignores the zero value for a (0)
func max(a, b int) int {
	if a == 0 || b > a {
		return b
	}
	return a
}

// minMaxModelNumber is based on the realization/assumption that only the "z" register of the ALU
// carries over between each program chunk processing an INP instruction.
// It builds a map of 14 maps, one for each digit of a "model number".
// Each of the 14 maps contains the possible values for the "z" register up to the corresponding digit.
// It is based on the assumption that there are far less than 9^14 possible values for "z".
func minMaxModelNumber(alu *ALU, program []Instruction) (int, int) {
	type minMaxNumbers struct {
		min, max int
	}
	zeroValue := minMaxNumbers{}

	aluStates := make([]map[int]minMaxNumbers, 15)
	aluStates[0] = make(map[int]minMaxNumbers)
	aluStates[0][0] = zeroValue

	programChunks := splitInstructions(program)

	for digit := 0; digit < 14; digit++ {
		aluStates[digit + 1] = make(map[int]minMaxNumbers)
		replaced := 0
		for input := 1; input < 10; input++ {
			alu.input = []int{input}
			for state, numberPath := range aluStates[digit] {
				// Only register "z" carries over from previous digits
				alu.x = 0
				alu.y = 0
				alu.z = state
				alu.index = 0
				for _, instruction := range programChunks[digit] {
					instruction.Execute()
				}
				oldNumberPath := aluStates[digit + 1][alu.z]
				newNumberPath := minMaxNumbers{
					min(oldNumberPath.min, numberPath.min * 10 + input),
					max(oldNumberPath.max, numberPath.max * 10 + input),
				}
				if oldNumberPath != newNumberPath {
					aluStates[digit + 1][alu.z] = newNumberPath
					if oldNumberPath != zeroValue {
						replaced++
					}
				}
			}
		}
		fmt.Printf("states at digit %v: %v (replaced: %v)\n", digit, len(aluStates[digit + 1]), replaced)
	}

	result := aluStates[14][0]
	return result.min, result.max
}

func main() {
	alu, program := parseInput(loadInput("puzzle-input-orig.txt"))

	start := time.Now()
	minModelNumber, maxModelNumber := minMaxModelNumber(alu, program)
	duration := time.Since(start)

	fmt.Printf("lowest / highest valid model number: %v / %v (%v)\n", minModelNumber, maxModelNumber, duration)
}

func inputFor(input string, alu *ALU) Input {
	r := alu.Register(input)
	if r != nil {
		return &Register{r}
	}
	v, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("invalid input: %s", input))
	}
	return &Number{v}
}

func parseInput(input string) (alu *ALU, program []Instruction) {
	alu = &ALU{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		parts := strings.Split(line, " ")
		r := alu.Register(parts[1])
		var i Input
		if len(parts) == 3 {
			i = inputFor(parts[2], alu)
		} else {
			i = alu
		}
		switch parts[0] {
		case "inp":
			program = append(program, &INP{r, i})
		case "add":
			program = append(program, &ADD{r, i})
		case "mul":
			program = append(program, &MUL{r, i})
		case "div":
			program = append(program, &DIV{r, i})
		case "mod":
			program = append(program, &MOD{r, i})
		case "eql":
			program = append(program, &EQL{r, i})
		}
	}
	return
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
