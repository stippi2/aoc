package day12

import (
	"aoc/2024/go/lib"
	"strings"
)

func calculateRegionPrice(input string) int {
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
		perimeter := 0
		for len(queue) > 0 {
			pos := queue[len(queue)-1]
			queue = queue[:len(queue)-1]

			area++

			neighbors := []lib.Vec2{
				{X: pos.X - 1, Y: pos.Y},
				{X: pos.X + 1, Y: pos.Y},
				{X: pos.X, Y: pos.Y - 1},
				{X: pos.X, Y: pos.Y + 1},
			}
			for _, neighbor := range neighbors {
				if garden.Contains(neighbor.X, neighbor.Y) {
					neighborPlant := garden.Get(neighbor.X, neighbor.Y)
					if neighborPlant != plant {
						perimeter++
					} else {
						if !visited[neighbor] {
							queue = append(queue, neighbor)
							visited[neighbor] = true
						}
					}
				} else {
					perimeter++
				}
			}
		}
		price += area * perimeter
	}

	return price
}

func Part1() any {
	input, _ := lib.ReadInput(12)
	return calculateRegionPrice(strings.TrimSpace(input))
}

func Part2() any {
	return "Not implemented"
}
