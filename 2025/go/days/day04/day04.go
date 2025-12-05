package day04

import (
	"aoc/2025/go/lib"
)

func countNeighbors(grid *lib.Grid, x, y int) int {
	count := 0

	neighbors := []lib.Vec2{
		{X: -1, Y: -1},
		{X: 0, Y: -1},
		{X: 1, Y: -1},

		{X: -1, Y: 0},
		{X: 1, Y: 0},

		{X: -1, Y: 1},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
	}

	for _, n := range neighbors {
		if grid.Get(x+n.X, y+n.Y) == '@' {
			count++
		}
	}

	return count
}

func countAndRemoveAccessiblePaperRolls(grid *lib.Grid) (int, *lib.Grid) {
	result := lib.NewGridFilled(grid.Width(), grid.Height(), '.')

	count := 0

	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if grid.Get(x, y) == '@' {
				if countNeighbors(grid, x, y) < 4 {
					count++
				} else {
					result.Set(x, y, '@')
				}
			}
		}
	}

	return count, result
}

func countAccessiblePaperRolls(input string) int {
	grid := lib.NewGrid(input)
	count, _ := countAndRemoveAccessiblePaperRolls(grid)
	return count
}

func countRemovablePaperRolls(input string) int {
	grid := lib.NewGrid(input)
	total := 0

	for {
		count, newGrid := countAndRemoveAccessiblePaperRolls(grid)
		if count == 0 {
			break
		}
		total += count
		grid = newGrid
	}

	return total
}

func Part1() any {
	input, _ := lib.ReadInput(4)
	return countAccessiblePaperRolls(input)
}

func Part2() any {
	input, _ := lib.ReadInput(4)
	return countRemovablePaperRolls(input)
}
