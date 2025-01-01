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

func (l *Line) direction() string {
	if l.a.Y == l.b.Y {
		if l.a.X < l.b.X {
			return "left"
		} else {
			return "right"
		}
	} else {
		if l.a.Y < l.b.Y {
			return "up"
		} else {
			return "down"
		}
	}
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

	mergeLines := func(a, b Line) *Line {
		if a.direction() == b.direction() {
			if a.a == b.b {
				return &Line{a: b.a, b: a.b}
			} else if b.a == a.b {
				return &Line{a: a.a, b: b.b}
			}
		}
		return nil
	}

	lines := make([]Line, len(p.lines))
	copy(lines, p.lines)

	for {
		hasMerged := false
		var newLines []Line
		used := make(map[int]bool)

		// Try to merge every line with every other
		for i := 0; i < len(lines); i++ {
			if used[i] {
				continue
			}

			currentLine := lines[i]
			mergedThisLine := false

			for j := i + 1; j < len(lines); j++ {
				if used[j] {
					continue
				}

				if merged := mergeLines(currentLine, lines[j]); merged != nil {
					// Merge found, mark indices as used
					used[i] = true
					used[j] = true
					// Keep the merge result for the next iteration
					newLines = append(newLines, *merged)
					hasMerged = true
					mergedThisLine = true
					break
				}
			}

			// If this line has not been merged, keep it unchanged
			if !mergedThisLine {
				newLines = append(newLines, currentLine)
			}
		}

		lines = newLines

		// If no merges were found, we are done
		if !hasMerged {
			break
		}
	}

	return len(lines)
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
					line: Line{a: lib.Vec2{X: pos.X + 1, Y: pos.Y + 1}, b: lib.Vec2{X: pos.X + 1, Y: pos.Y}},
				},
				{
					pos:  lib.Vec2{X: pos.X, Y: pos.Y - 1},
					line: Line{a: lib.Vec2{X: pos.X + 1, Y: pos.Y}, b: lib.Vec2{X: pos.X, Y: pos.Y}},
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
