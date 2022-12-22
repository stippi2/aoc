package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

func (p Pos) add(vector Pos) Pos {
	return Pos{p.x + vector.x, p.y + vector.y}
}

func (p Pos) rotateLeft() Pos {
	return Pos{p.y, -p.x}
}

func (p Pos) rotateRight() Pos {
	return Pos{-p.y, p.x}
}

func (p Pos) negate() Pos {
	return Pos{-p.x, -p.y}
}

func (p Pos) String() string {
	return fmt.Sprintf("[%v, %v]", p.x, p.y)
}

type Instructions struct {
	input string
	pos   int
}

func (s *Instructions) next() string {
	if s.pos == len(s.input) {
		return ""
	}
	length := 1
	defer func() {
		s.pos += length
	}()
	token := s.input[s.pos : s.pos+length]
	if token == "L" || token == "R" {
		return token
	}
	for s.pos+length < len(s.input) {
		token = s.input[s.pos+length : s.pos+length+1]
		if token == "L" || token == "R" {
			break
		}
		length++
	}
	return s.input[s.pos : s.pos+length]
}

type Explorer struct {
	location Pos
	facing   Pos
	path     map[Pos]string
}

func (e *Explorer) tracePath() {
	var trace string
	switch e.facing {
	case Pos{1, 0}:
		trace = ">"
	case Pos{0, -1}:
		trace = "^"
	case Pos{-1, 0}:
		trace = "<"
	case Pos{0, 1}:
		trace = "v"
	default:
		panic("invalid facing")
	}
	if e.path == nil {
		e.path = make(map[Pos]string)
	}
	e.path[e.location] = trace
}

func (e *Explorer) getPassword() int {
	var facing int
	switch e.facing {
	case Pos{1, 0}:
		facing = 0
	case Pos{0, -1}:
		facing = 3
	case Pos{-1, 0}:
		facing = 2
	case Pos{0, 1}:
		facing = 1
	}
	return 1000*(e.location.y+1) + 4*(e.location.x+1) + facing
}

type Map struct {
	lines []string
}

func (m *Map) width() int {
	width := 0
	for _, line := range m.lines {
		lineWidth := len(line)
		if lineWidth > width {
			width = lineWidth
		}
	}
	return width
}

func (m *Map) getLocation(p Pos) string {
	if p.y < 0 || p.y >= len(m.lines) {
		return " "
	}
	line := m.lines[p.y]
	if p.x < 0 || p.x >= len(line) {
		return " "
	}
	return line[p.x : p.x+1]
}

func (m *Map) startingPos() *Explorer {
	e := &Explorer{
		location: Pos{0, 0},
		facing:   Pos{1, 0},
	}
	for m.getLocation(e.location) == " " {
		e.location.x++
	}
	return e
}

func executeInstructions(m *Map, i *Instructions, e *Explorer) {
	for {
		instructions := i.next()
		if instructions == "" {
			break
		}
		switch instructions {
		case "L":
			e.facing = e.facing.rotateLeft()
			e.tracePath()
		case "R":
			e.facing = e.facing.rotateRight()
			e.tracePath()
		default:
			distance, _ := strconv.Atoi(instructions)
			for distance > 0 {
				newLocation := e.location.add(e.facing)
				if m.getLocation(newLocation) == " " {
					// Wrap around
					facingOpposite := e.facing.negate()
					for m.getLocation(newLocation.add(facingOpposite)) != " " {
						newLocation = newLocation.add(facingOpposite)
					}
				}
				switch m.getLocation(newLocation) {
				case ".":
					e.tracePath()
					e.location = newLocation
				case "#":
					// Nothing
				}
				distance--
			}
		}
	}
	e.tracePath()
}

func printMap(m *Map, e *Explorer) {
	for y := 0; y < len(m.lines); y++ {
		for x := 0; x < m.width(); x++ {
			trace := e.path[Pos{x, y}]
			if trace != "" {
				fmt.Print(trace)
			} else {
				fmt.Print(m.getLocation(Pos{x, y}))
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	m, instructions := parseInput(loadInput("puzzle-input.txt"))
	explorer := m.startingPos()
	fmt.Printf("start pos: %s\n", explorer.location)
	executeInstructions(m, instructions, explorer)
	fmt.Printf("end pos: %s, password is %v\n", explorer.location, explorer.getPassword())
}

func parseInput(input string) (*Map, *Instructions) {
	parts := strings.Split(input, "\n\n")
	return &Map{lines: strings.Split(parts[0], "\n")}, &Instructions{input: parts[1]}
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
