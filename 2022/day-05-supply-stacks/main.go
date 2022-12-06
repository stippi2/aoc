package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Crate struct {
	label string
}

type CrateStack struct {
	crates []Crate
}

func (s *CrateStack) topLabel() string {
	last := len(s.crates) - 1
	if last == -1 {
		return ""
	}
	return s.crates[last].label
}

func (s *CrateStack) removeTop() (Crate, error) {
	last := len(s.crates) - 1
	if last == -1 {
		return Crate{}, errors.New("empty stack")
	}
	crate := s.crates[last]
	s.crates = s.crates[:last]
	return crate, nil
}

func (s *CrateStack) addTop(crate Crate) {
	s.crates = append(s.crates, crate)
}

func (s *CrateStack) removeStack(count int) ([]Crate, error) {
	if count > len(s.crates) {
		return nil, errors.New("not enough crates in stack")
	}
	crates := s.crates[len(s.crates)-count:]
	s.crates = s.crates[:len(s.crates)-count]
	return crates, nil
}

func (s *CrateStack) addStack(crates []Crate) {
	s.crates = append(s.crates, crates...)
}

type MoveInstruction struct {
	from  int
	to    int
	count int
}

func applyMoveInstructions(stacks []CrateStack, instructions []MoveInstruction, moveAtOnce bool) {
	for _, instruction := range instructions {
		from := &stacks[instruction.from-1]
		to := &stacks[instruction.to-1]
		if moveAtOnce {
			crates, err := from.removeStack(instruction.count)
			if err != nil {
				panic(err)
			}
			to.addStack(crates)
		} else {
			for i := 0; i < instruction.count; i++ {
				crate, err := from.removeTop()
				if err != nil {
					panic(err)
				}
				to.addTop(crate)
			}
		}
	}
}

func getTopCrates(stacks []CrateStack) string {
	var topCrates string
	for _, stack := range stacks {
		topCrates += stack.topLabel()
	}
	return topCrates
}

func main() {
	stacks, instructions := parseInput(loadInput("puzzle-input.txt"), 9)
	applyMoveInstructions(stacks, instructions, false)
	fmt.Printf("top crates part 1: %v\n", getTopCrates(stacks))

	stacks, instructions = parseInput(loadInput("puzzle-input.txt"), 9)
	applyMoveInstructions(stacks, instructions, true)
	fmt.Printf("top crates part 2: %v\n", getTopCrates(stacks))
}

func parseStacks(stacksInput string, stackCount int) []CrateStack {
	stacksInput = strings.Trim(stacksInput, "\n")
	stacks := make([]CrateStack, stackCount)
	lines := strings.Split(stacksInput, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		lineOffset := 1
		for s := 0; s < stackCount; s++ {
			if lineOffset >= len(lines[i]) {
				break
			}
			crateLabel := string(lines[i][lineOffset])
			if crateLabel != " " {
				stacks[s].addTop(Crate{crateLabel})
			}
			lineOffset += 4
		}
	}
	return stacks
}

func parseMoveInstructions(instructionsInput string) []MoveInstruction {
	var instructions []MoveInstruction
	instructionsInput = strings.TrimSpace(instructionsInput)
	lines := strings.Split(instructionsInput, "\n")
	for _, line := range lines {
		var instruction MoveInstruction
		count, err := fmt.Sscanf(line, "move %d from %d to %d", &instruction.count, &instruction.from, &instruction.to)
		if err != nil || count != 3 {
			panic("scanning instructions failed")
		}
		instructions = append(instructions, instruction)
	}
	return instructions
}

func parseInput(input string, stackCount int) ([]CrateStack, []MoveInstruction) {
	divider := " "
	for i := 1; i <= stackCount; i++ {
		divider += strconv.Itoa(i) + "   "
	}
	divider = strings.TrimRight(divider, " ")
	parts := strings.Split(input, divider)
	return parseStacks(parts[0], stackCount), parseMoveInstructions(parts[1])
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
