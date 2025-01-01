package day12

import (
	"aoc/2024/go/lib"
	"fmt"
	"strings"
)

type Perimeter interface {
	extend(line Line, direction string)
	getLength() int
}

type RealPerimeter struct {
	length int
}

func (p *RealPerimeter) extend(_ Line, _ string) {
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
	left  []Line
	right []Line
	up    []Line
	down  []Line
}

func (p *DiscountedPerimeter) extend(line Line, direction string) {
	switch direction {
	case "left":
		p.left = append(p.left, line)
	case "right":
		p.right = append(p.right, line)
	case "up":
		p.up = append(p.up, line)
	case "down":
		p.down = append(p.down, line)
	}
}

func (p *DiscountedPerimeter) getLength() int {
	return 0
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
				pos       lib.Vec2
				line      Line
				direction string
			}

			neighbors := []Neighbor{
				{
					pos:       lib.Vec2{X: pos.X - 1, Y: pos.Y},
					line:      Line{a: lib.Vec2{X: pos.X, Y: pos.Y}, b: lib.Vec2{X: pos.X, Y: pos.Y + 1}},
					direction: "down",
				},
				{
					pos:       lib.Vec2{X: pos.X + 1, Y: pos.Y},
					line:      Line{a: lib.Vec2{X: pos.X + 1, Y: pos.Y + 1}, b: lib.Vec2{X: pos.X + 1, Y: pos.Y}},
					direction: "up",
				},
				{
					pos:       lib.Vec2{X: pos.X, Y: pos.Y - 1},
					line:      Line{a: lib.Vec2{X: pos.X + 1, Y: pos.Y}, b: lib.Vec2{X: pos.X, Y: pos.Y}},
					direction: "left",
				},
				{
					pos:       lib.Vec2{X: pos.X, Y: pos.Y + 1},
					line:      Line{a: lib.Vec2{X: pos.X, Y: pos.Y + 1}, b: lib.Vec2{X: pos.X + 1, Y: pos.Y + 1}},
					direction: "right",
				},
			}
			for _, neighbor := range neighbors {
				if garden.Contains(neighbor.pos.X, neighbor.pos.Y) {
					neighborPlant := garden.Get(neighbor.pos.X, neighbor.pos.Y)
					if neighborPlant != plant {
						perimeter.extend(neighbor.line, neighbor.direction)
					} else {
						if !visited[neighbor.pos] {
							queue = append(queue, neighbor.pos)
							// Set visited here, to avoid adding neighbors already in the queue
							visited[neighbor.pos] = true
						}
					}
				} else {
					perimeter.extend(neighbor.line, neighbor.direction)
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
