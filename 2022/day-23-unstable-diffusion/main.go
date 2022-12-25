package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Pos struct {
	x, y int
}

func (p Pos) add(vector Pos) Pos {
	return Pos{p.x + vector.x, p.y + vector.y}
}

func (p Pos) neighbors() []Pos {
	return []Pos{
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y + 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y + 1},
		{p.x - 1, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y + 1},
	}
}

func (p Pos) String() string {
	return fmt.Sprintf("[%v, %v]", p.x, p.y)
}

type Map struct {
	elves map[Pos]bool
}

func (m *Map) boundingBox() (xMin, xMax, yMin, yMax int) {
	xMin = math.MaxInt
	xMax = math.MinInt
	yMin = math.MaxInt
	yMax = math.MinInt
	for elf := range m.elves {
		xMin = min(xMin, elf.x)
		xMax = max(xMax, elf.x)
		yMin = min(yMin, elf.y)
		yMax = max(yMax, elf.y)
	}
	return
}

func (m *Map) String() string {
	xMin, xMax, yMin, yMax := m.boundingBox()
	result := ""
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if m.elves[Pos{x, y}] {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return strings.TrimSpace(result)
}

func simulateOneDiffusionRound(m *Map, roundIndex int) bool {
	proposedPositions := make(map[Pos]int)
	intendedMoves := make(map[Pos]Pos)
	for elf := range m.elves {
		allEmpty := true
		foundValidDirection := false
		for _, n := range elf.neighbors() {
			if m.elves[n] {
				allEmpty = false
				break
			}
		}
		if !allEmpty {
			directions := [][]Pos{
				{{elf.x - 1, elf.y - 1}, {elf.x, elf.y - 1}, {elf.x + 1, elf.y - 1}},
				{{elf.x - 1, elf.y + 1}, {elf.x, elf.y + 1}, {elf.x + 1, elf.y + 1}},
				{{elf.x - 1, elf.y - 1}, {elf.x - 1, elf.y}, {elf.x - 1, elf.y + 1}},
				{{elf.x + 1, elf.y - 1}, {elf.x + 1, elf.y}, {elf.x + 1, elf.y + 1}},
			}
			directionShift := roundIndex % len(directions)

			for i := 0; i < len(directions); i++ {
				direction := directions[(i+directionShift)%len(directions)]
				directionEmpty := true
				for _, n := range direction {
					if m.elves[n] {
						directionEmpty = false
						break
					}
				}
				if directionEmpty {
					count := proposedPositions[direction[1]]
					proposedPositions[direction[1]] = count + 1
					intendedMoves[elf] = direction[1]
					foundValidDirection = true
					break
				}
			}
		}
		if allEmpty || !foundValidDirection {
			// Elf does not move
			count := proposedPositions[elf]
			proposedPositions[elf] = count + 1
			intendedMoves[elf] = elf
		}
	}
	resultingElves := make(map[Pos]bool)
	elvesMoved := 0
	for elf, newElf := range intendedMoves {
		if proposedPositions[newElf] == 1 {
			if elf != newElf {
				elvesMoved++
			}
			resultingElves[newElf] = true
		} else {
			resultingElves[elf] = true
		}
	}
	m.elves = resultingElves
	return elvesMoved > 0
}

func simulateDiffusion(m *Map, rounds int) {
	for i := 0; i < rounds; i++ {
		simulateOneDiffusionRound(m, i)
	}
}

func simulateDiffusionUntilSettled(m *Map) int {
	round := 0
	for {
		if !simulateOneDiffusionRound(m, round) {
			break
		}
		round++
	}
	return round + 1
}

func countEmptyPositions(m *Map) int {
	xMin, xMax, yMin, yMax := m.boundingBox()
	emptyPositions := 0
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if !m.elves[Pos{x, y}] {
				emptyPositions++
			}
		}
	}
	return emptyPositions
}

func main() {
	m := parseInput(loadInput("puzzle-input.txt"))
	simulateDiffusion(m, 10)
	fmt.Printf("empty positions in bounding box: %v\n", countEmptyPositions(m))

	m = parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("rounds needed until no elf moves: %v\n", simulateDiffusionUntilSettled(m))
}

func parseInput(input string) *Map {
	m := &Map{elves: make(map[Pos]bool)}
	for y, line := range strings.Split(input, "\n") {
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				m.elves[Pos{x, y}] = true
			}
		}
	}
	return m
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
