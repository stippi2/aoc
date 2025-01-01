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

func trackGuardFields(grid *lib.Grid, start lib.Vec2, obstacle *lib.Vec2) (visited map[lib.Vec2]bool, trapped bool) {
	current := start
	visited = make(map[lib.Vec2]bool)

	visitedByDirection := make(map[lib.Vec2]map[lib.Vec2]bool)
	directions := []lib.Vec2{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
	}
	for _, direction := range directions {
		visitedByDirection[direction] = make(map[lib.Vec2]bool)
	}

	direction := lib.Vec2{X: 0, Y: -1}
	for {
		next := current.Add(direction)
		visited[current] = true
		if visitedByDirection[direction][current] {
			return visited, true
		}
		visitedByDirection[direction][current] = true
		if !grid.Contains(next.X, next.Y) {
			break
		}
		if grid.Get(next.X, next.Y) == '#' || (obstacle != nil && *obstacle == next) {
			// Next char is obstacle, rotate right 90°
			direction = lib.Vec2{X: -direction.Y, Y: direction.X}
		} else {
			// Continue in this direction
			current = next
		}
	}
	return visited, false
}

func Part1() interface{} {
	input, _ := lib.ReadInput(6)
	grid := lib.NewGrid(input)
	start := findStart(grid)
	fieldsVisited, _ := trackGuardFields(grid, start, nil)
	return len(fieldsVisited)
}

func countLoopCausingObstacles(grid *lib.Grid, start lib.Vec2, visited map[lib.Vec2]bool) int {
	loopCausingObstacles := 0
	for position := range visited {
		if position == start {
			continue
		}
		_, isLoop := trackGuardFields(grid, start, &position)
		if isLoop {
			loopCausingObstacles++
		}
	}
	return loopCausingObstacles
}

func Part2() interface{} {
	input, _ := lib.ReadInput(6)
	grid := lib.NewGrid(input)
	start := findStart(grid)
	fieldsVisited, _ := trackGuardFields(grid, start, nil)
	return countLoopCausingObstacles(grid, start, fieldsVisited)
}
