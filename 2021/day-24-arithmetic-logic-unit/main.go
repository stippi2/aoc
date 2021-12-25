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

type ModelNumber struct {
	value []int
}

func newModelNumber() *ModelNumber {
	// 99999145946229
	n := &ModelNumber{}
	n.value = make([]int, 14)
	for i := 0; i < len(n.value); i++ {
		n.value[i] = 9
	}
	return n
}

func (n *ModelNumber) setFrom(input string) {
	for i := 0; i < len(input) && i < len(n.value); i++ {
		digit, err := strconv.Atoi(input[i:i+1])
		if err != nil {
			panic(fmt.Sprintf("failed to parse digit: %s", err))
		}
		n.value[i] = digit
	}
}

func (n *ModelNumber) decrement() {
	digit := len(n.value) - 1
	for digit > 0 {
		n.value[digit] = n.value[digit] - 1
		if n.value[digit] != 0 {
			break
		}
		n.value[digit] = 9
		digit--
	}
}

func (n *ModelNumber) increment() {
	digit := len(n.value) - 1
	for digit > 0 {
		n.value[digit] = n.value[digit] + 1
		if n.value[digit] != 10 {
			break
		}
		n.value[digit] = 1
		digit--
	}
}

func (n *ModelNumber) String() string {
	result := ""
	for i := 0; i < len(n.value); i++ {
		result += strconv.Itoa(n.value[i])
	}
	return result
}

func main() {
	start := time.Now()
	alu, program := parseInput(loadInput("puzzle-input.txt"))
	modelNumber := newModelNumber()
	alu.input = modelNumber.value
	iteration := 0
	for {
		for _, instruction := range program {
			instruction.Execute()
			fmt.Printf("ALU: w: %v, x: %v, y: %v, z: %v\n", alu.w, alu.x, alu.y, alu.z)
		}
		if iteration == 0 {
			break
		}
		if alu.z == 0 {
			break
		}
		iteration++
		if iteration % 100000 == 0 {
			duration := time.Since(start)
			durationPerPrint := duration / time.Duration(iteration / 100000)
			durationAll := durationPerPrint * 999999999
			fmt.Printf("model number %s: invalid, %s per 100000 iterations, %s for all\n", modelNumber, durationPerPrint, durationAll)
		}
		modelNumber.decrement()
		alu.reset()
	}
	fmt.Printf("largest valid model number: %v (%v)\n", modelNumber, time.Since(start))
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

func parseInput(input string) (alu *ALU, program []Instruction){
	alu = &ALU{}
	for _, line := range strings.Split(input, "\n") {
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
