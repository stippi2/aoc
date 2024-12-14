package day10

import (
	"aoc/2024/go/lib"
)

func findTrails(grid *lib.Grid, start lib.Vec2) int {
	reachableNines := make(map[lib.Vec2]bool)
	visited := make(map[lib.Vec2]bool)

	queue := []lib.Vec2{start}
	for len(queue) > 0 {
		pos := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		visited[pos] = true
		height := grid.Get(pos.X, pos.Y) - '0'

		if height == 9 {
			reachableNines[pos] = true
			continue
		}
		neighbors := []lib.Vec2{
			{X: pos.X - 1, Y: pos.Y},
			{X: pos.X + 1, Y: pos.Y},
			{X: pos.X, Y: pos.Y - 1},
			{X: pos.X, Y: pos.Y + 1},
		}
		for _, next := range neighbors {
			if !grid.Contains(next.X, next.Y) || visited[next] {
				continue
			}
			nextHeight := grid.Get(next.X, next.Y) - '0'
			if nextHeight == height+1 {
				queue = append(queue, next)
			}
		}
	}

	return len(reachableNines)
}

func sumTrailHeads(input string) int {
	grid := lib.NewGrid(input)
	sum := 0
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			value := grid.Get(x, y) - '0'
			if value == 0 {
				sum += findTrails(grid, lib.Vec2{X: x, Y: y})
			}
		}
	}
	return sum
}

func Part1() int {
	input, _ := lib.ReadInput(10)
	return sumTrailHeads(input)
}

func Part2() any {
	return "Not implemented"
}
