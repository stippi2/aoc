package day06

import "aoc/2024/go/lib"

func findStart(grid *lib.Grid) lib.Vec2 {
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if grid.Get(x, y) == '^' {
				return lib.Vec2{X: x, Y: y}
			}
		}
	}
	panic("start position not found")
}

func sumVisitedFields(input string) int {
	grid := lib.NewGrid(input)
	current := findStart(grid)
	visited := make(map[lib.Vec2]bool)
	direction := lib.Vec2{X: 0, Y: -1}
	for {
		next := current.Add(direction)
		visited[current] = true
		if !grid.Contains(next.X, next.Y) {
			break
		}
		switch grid.Get(next.X, next.Y) {
		case '#':
			// Next char is obstacle, rotate right 90°
			direction = lib.Vec2{X: -direction.Y, Y: direction.X}
		default:
			// Continue in this direction
			current = next
		}
	}
	return len(visited)
}

func Part1() interface{} {
	input, _ := lib.ReadInput(6)
	return sumVisitedFields(input)
}

func Part2() interface{} {
	return "Not implemented"
}
