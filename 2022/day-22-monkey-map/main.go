package main

import (
	"fmt"
	"os"
	"strconv"
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

func (p Pos) rotateLeft() Pos {
	return Pos{p.y, -p.x}
}

func (p Pos) rotateRight() Pos {
	return Pos{-p.y, p.x}
}

func (p Pos) negate() Pos {
	return Pos{-p.x, -p.y}
}

func (p Pos) isBetween(a, b Pos) bool {
	xMin := min(a.x, b.x)
	xMax := max(a.x, b.x)
	yMin := min(a.y, b.y)
	yMax := max(a.y, b.y)
	return p.x >= xMin && p.x <= xMax && p.y >= yMin && p.y <= yMax
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

type WrappingFunction func(m *Map, e *Explorer) (Pos, Pos)

func executeInstructions(m *Map, i *Instructions, e *Explorer, handleWrap WrappingFunction) {
	for {
		instruction := i.next()
		if instruction == "" {
			break
		}
		switch instruction {
		case "L":
			e.facing = e.facing.rotateLeft()
			e.tracePath()
		case "R":
			e.facing = e.facing.rotateRight()
			e.tracePath()
		default:
			distance, _ := strconv.Atoi(instruction)
			for distance > 0 {
				newLocation, newFacing := handleWrap(m, e)
				switch m.getLocation(newLocation) {
				case ".":
					e.tracePath()
					e.location = newLocation
					e.facing = newFacing
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

func handleWrapPartOne(m *Map, e *Explorer) (Pos, Pos) {
	newLocation := e.location.add(e.facing)
	if m.getLocation(newLocation) == " " {
		facingOpposite := e.facing.negate()
		for m.getLocation(newLocation.add(facingOpposite)) != " " {
			newLocation = newLocation.add(facingOpposite)
		}
	}
	return newLocation, e.facing
}

type Edge struct {
	a, b     Pos
	facing   Pos
	rotation string
}

type CubeSeam struct {
	label string
	edges []Edge
}

var seams = []CubeSeam{
	{
		label: "FC",
		edges: []Edge{
			{
				a:        Pos{149, 0},
				b:        Pos{149, 49},
				facing:   Pos{1, 0},
				rotation: "RR",
			},
			{
				a:        Pos{99, 149},
				b:        Pos{99, 100},
				facing:   Pos{1, 0},
				rotation: "LL",
			},
		},
	},
	{
		label: "FD",
		edges: []Edge{
			{
				a:        Pos{100, 49},
				b:        Pos{149, 49},
				facing:   Pos{0, 1},
				rotation: "R",
			},
			{
				a:        Pos{99, 50},
				b:        Pos{99, 99},
				facing:   Pos{1, 0},
				rotation: "L",
			},
		},
	},
	{
		label: "CA",
		edges: []Edge{
			{
				a:        Pos{50, 149},
				b:        Pos{99, 149},
				facing:   Pos{0, 1},
				rotation: "R",
			},
			{
				a:        Pos{49, 150},
				b:        Pos{49, 199},
				facing:   Pos{1, 0},
				rotation: "L",
			},
		},
	},
	{
		label: "AF",
		edges: []Edge{
			{
				a:        Pos{0, 199},
				b:        Pos{49, 199},
				facing:   Pos{0, 1},
				rotation: "",
			},
			{
				a:        Pos{100, 0},
				b:        Pos{149, 0},
				facing:   Pos{0, -1},
				rotation: "",
			},
		},
	},
	{
		label: "AE",
		edges: []Edge{
			{
				a:        Pos{0, 150},
				b:        Pos{0, 199},
				facing:   Pos{-1, 0},
				rotation: "L",
			},
			{
				a:        Pos{99, 0},
				b:        Pos{50, 0},
				facing:   Pos{0, -1},
				rotation: "R",
			},
		},
	},
	{
		label: "BE",
		edges: []Edge{
			{
				a:        Pos{0, 100},
				b:        Pos{0, 149},
				facing:   Pos{-1, 0},
				rotation: "RR",
			},
			{
				a:        Pos{50, 49},
				b:        Pos{50, 0},
				facing:   Pos{-1, 0},
				rotation: "LL",
			},
		},
	},
	{
		label: "BD",
		edges: []Edge{
			{
				a:        Pos{0, 100},
				b:        Pos{49, 100},
				facing:   Pos{0, -1},
				rotation: "R",
			},
			{
				a:        Pos{50, 50},
				b:        Pos{50, 99},
				facing:   Pos{-1, 0},
				rotation: "L",
			},
		},
	},
}

func project(location, from, to Pos, rotation string) Pos {
	location.x -= from.x
	location.y -= from.y
	for i := 0; i < len(rotation); i++ {
		switch rotation[i : i+1] {
		case "L":
			location = location.rotateLeft()
		case "R":
			location = location.rotateRight()
		}
	}
	location.x += to.x
	location.y += to.y
	return location
}

func handleWrapPartTwo(_ *Map, e *Explorer) (Pos, Pos) {
	for _, seam := range seams {
		for i, edge := range seam.edges {
			if e.facing == edge.facing && e.location.isBetween(edge.a, edge.b) {
				otherEdge := seam.edges[(i+1)%2]
				newFacing := otherEdge.facing.negate()
				newLocation := project(e.location, edge.a, otherEdge.a, edge.rotation)
				return newLocation, newFacing
			}
		}
	}
	return e.location.add(e.facing), e.facing
}

func main() {
	m, instructions := parseInput(loadInput("puzzle-input.txt"))
	explorer := m.startingPos()
	fmt.Printf("start pos: %s\n", explorer.location)
	executeInstructions(m, instructions, explorer, handleWrapPartOne)
	fmt.Printf("part 1 end pos: %s, password is %v\n", explorer.location, explorer.getPassword())

	explorer = m.startingPos()
	instructions.pos = 0
	executeInstructions(m, instructions, explorer, handleWrapPartTwo)
	fmt.Printf("part 2 end pos: %s, password is %v\n", explorer.location, explorer.getPassword())
}

func parseInput(input string) (*Map, *Instructions) {
	parts := strings.Split(input, "\n\n")
	return &Map{lines: strings.Split(parts[0], "\n")}, &Instructions{input: parts[1]}
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
