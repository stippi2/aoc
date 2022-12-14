package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	Rock       int = 1
	SandAtRest     = 2
)

type Pos struct {
	x, y int
}

type Cave struct {
	data   map[Pos]int
	bottom int
	floor  int
}

func (c *Cave) setRock(p Pos) {
	c.data[p] = Rock
	if p.y > c.bottom {
		c.bottom = p.y
	}
}

func simulateSand(start Pos, c *Cave) int {
	sandAtRest := 0
	for {
		p := start
		if c.data[p] != 0 {
			break
		}
		cameToRest := false
		for {
			if c.floor > c.bottom {
				if p.y == c.floor-1 {
					cameToRest = true
					break
				}
			} else if p.y == c.bottom {
				break
			}
			if c.data[Pos{p.x, p.y + 1}] == 0 {
				p.y++
			} else if c.data[Pos{p.x - 1, p.y + 1}] == 0 {
				p.x--
				p.y++
			} else if c.data[Pos{p.x + 1, p.y + 1}] == 0 {
				p.x++
				p.y++
			} else {
				cameToRest = true
				break
			}
		}
		if cameToRest {
			c.data[p] = SandAtRest
			sandAtRest++
		} else {
			break
		}
	}
	return sandAtRest
}

func main() {
	cave := parseInput(loadInput("puzzle-input.txt"))
	sandCount := simulateSand(Pos{500, 0}, cave)
	fmt.Printf("blocks with sand at rest: %v\n", sandCount)

	cave = parseInput(loadInput("puzzle-input.txt"))
	cave.floor = cave.bottom + 2
	sandCount = simulateSand(Pos{500, 0}, cave)
	fmt.Printf("with floor, blocks with sand at rest: %v\n", sandCount)
}

func parseCoords(coords string) Pos {
	xy := strings.Split(coords, ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	return Pos{x, y}
}

func parseInput(input string) *Cave {
	c := &Cave{
		data: make(map[Pos]int),
	}
	for _, line := range strings.Split(input, "\n") {
		coords := strings.Split(line, " -> ")
		for i := 1; i < len(coords); i++ {
			start := parseCoords(coords[i-1])
			end := parseCoords(coords[i])
			if start.x == end.x {
				y1 := int(math.Min(float64(start.y), float64(end.y)))
				y2 := int(math.Max(float64(start.y), float64(end.y)))
				for y := y1; y <= y2; y++ {
					c.setRock(Pos{start.x, y})
				}
			} else {
				x1 := int(math.Min(float64(start.x), float64(end.x)))
				x2 := int(math.Max(float64(start.x), float64(end.x)))
				for x := x1; x <= x2; x++ {
					c.setRock(Pos{x, start.y})
				}
			}
		}
	}
	return c
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
