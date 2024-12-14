package day10

import (
	"aoc/2024/go/lib"
)

func findTrails(grid *lib.Grid, start lib.Vec2) (int, int) {
	reachableNines := make(map[lib.Vec2]int)

	queue := []lib.Vec2{start}
	for len(queue) > 0 {
		pos := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		height := grid.Get(pos.X, pos.Y) - '0'

		if height == 9 {
			reachableNines[pos] += 1
			continue
		}
		neighbors := []lib.Vec2{
			{X: pos.X - 1, Y: pos.Y},
			{X: pos.X + 1, Y: pos.Y},
			{X: pos.X, Y: pos.Y - 1},
			{X: pos.X, Y: pos.Y + 1},
		}
		for _, next := range neighbors {
			if !grid.Contains(next.X, next.Y) {
				continue
			}
			nextHeight := grid.Get(next.X, next.Y) - '0'
			if nextHeight == height+1 {
				queue = append(queue, next)
			}
		}
	}

	sumTrails := 0
	for _, value := range reachableNines {
		sumTrails += value
	}
	return len(reachableNines), sumTrails
}

func sumTrailHeads(input string, countUniqueTrails bool) int {
	grid := lib.NewGrid(input)
	sum := 0
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			value := grid.Get(x, y) - '0'
			if value == 0 {
				reachableNines, uniqueTrails := findTrails(grid, lib.Vec2{X: x, Y: y})
				if countUniqueTrails {
					sum += uniqueTrails
				} else {
					sum += reachableNines
				}
			}
		}
	}
	return sum
}

func Part1() int {
	input, _ := lib.ReadInput(10)
	return sumTrailHeads(input, false)
}

func Part2() any {
	input, _ := lib.ReadInput(10)
	return sumTrailHeads(input, true)
}
