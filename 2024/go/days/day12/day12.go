package day12

import (
	"aoc/2024/go/lib"
	"fmt"
	"strings"
)

type Perimeter interface {
	extend(line Line)
	getLength() int
}

type RealPerimeter struct {
	length int
}

func (p *RealPerimeter) extend(_ Line) {
	p.length++
}

func (p *RealPerimeter) getLength() int {
	return p.length
}

type Line struct {
	a lib.Vec2
	b lib.Vec2
}

type DiscountedPerimeter struct {
	lines []Line
}

func (p *DiscountedPerimeter) extend(line Line) {
	p.lines = append(p.lines, line)
}

func (p *DiscountedPerimeter) getLength() int {
	if len(p.lines) == 0 {
		return 0
	}

	mergedLines := make([]Line, len(p.lines))
	copy(mergedLines, p.lines)

	for {
		merged := false

		for i := 0; i < len(mergedLines); i++ {
			for j := i + 1; j < len(mergedLines); j++ {
				if canMerge(mergedLines[i], mergedLines[j]) {
					mergedLines[i] = mergeTwoLines(mergedLines[i], mergedLines[j])
					mergedLines = append(mergedLines[:j], mergedLines[j+1:]...)
					merged = true
					break
				}
			}
			if merged {
				break
			}
		}

		if !merged {
			break
		}
	}

	return len(mergedLines)
}

func canMerge(l1, l2 Line) bool {
	if l1.b != l2.a {
		return false
	}

	if l1.a.Y == l1.b.Y && l2.a.Y == l2.b.Y {
		return true
	}

	if l1.a.X == l1.b.X && l2.a.X == l2.b.X {
		return true
	}

	return false
}

func mergeTwoLines(l1, l2 Line) Line {
	return Line{
		a: l1.a,
		b: l2.b,
	}
}

func calculateRegionPrice(input string, withDiscount bool) int {
	garden := lib.NewGrid(input)
	price := 0

	visited := make(map[lib.Vec2]bool)
	for len(visited) < garden.Width()*garden.Height() {
		// Find the next start of a not yet visited region
		var queue []lib.Vec2
		var plant byte
		for y := 0; y < garden.Height(); y++ {
			for x := 0; x < garden.Width(); x++ {
				start := lib.Vec2{X: x, Y: y}
				if !visited[start] {
					queue = append(queue, start)
					visited[start] = true
					plant = garden.Get(x, y)
					break
				}
			}
			if len(queue) == 1 {
				break
			}
		}
		// Explore the region
		area := 0
		var perimeter Perimeter
		if withDiscount {
			perimeter = &DiscountedPerimeter{}
		} else {
			perimeter = &RealPerimeter{}
		}
		for len(queue) > 0 {
			pos := queue[len(queue)-1]
			queue = queue[:len(queue)-1]

			area++

			type Neighbor struct {
				pos  lib.Vec2
				line Line
			}

			neighbors := []Neighbor{
				{
					pos:  lib.Vec2{X: pos.X - 1, Y: pos.Y},
					line: Line{a: lib.Vec2{X: pos.X, Y: pos.Y}, b: lib.Vec2{X: pos.X, Y: pos.Y + 1}},
				},
				{
					pos:  lib.Vec2{X: pos.X + 1, Y: pos.Y},
					line: Line{a: lib.Vec2{X: pos.X + 1, Y: pos.Y}, b: lib.Vec2{X: pos.X + 1, Y: pos.Y + 1}},
				},
				{
					pos:  lib.Vec2{X: pos.X, Y: pos.Y - 1},
					line: Line{a: lib.Vec2{X: pos.X, Y: pos.Y}, b: lib.Vec2{X: pos.X + 1, Y: pos.Y}},
				},
				{
					pos:  lib.Vec2{X: pos.X, Y: pos.Y + 1},
					line: Line{a: lib.Vec2{X: pos.X, Y: pos.Y + 1}, b: lib.Vec2{X: pos.X + 1, Y: pos.Y + 1}},
				},
			}
			for _, neighbor := range neighbors {
				if garden.Contains(neighbor.pos.X, neighbor.pos.Y) {
					neighborPlant := garden.Get(neighbor.pos.X, neighbor.pos.Y)
					if neighborPlant != plant {
						perimeter.extend(neighbor.line)
					} else {
						if !visited[neighbor.pos] {
							queue = append(queue, neighbor.pos)
							// Set visited here, to avoid adding neighbors already in the queue
							visited[neighbor.pos] = true
						}
					}
				} else {
					perimeter.extend(neighbor.line)
				}
			}
		}
		perimeterLength := perimeter.getLength()
		fmt.Printf("Region '%s' area: %v, perimeter: %v\n", string(plant), area, perimeterLength)
		price += area * perimeterLength
	}

	return price
}

func Part1() any {
	input, _ := lib.ReadInput(12)
	return calculateRegionPrice(strings.TrimSpace(input), false)
}

func Part2() any {
	input, _ := lib.ReadInput(12)
	return calculateRegionPrice(strings.TrimSpace(input), true)
}
