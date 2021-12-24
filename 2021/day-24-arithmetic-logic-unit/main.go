package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

func main() {
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
