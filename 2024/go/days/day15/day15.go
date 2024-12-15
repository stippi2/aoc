package day15

import (
	"aoc/2024/go/lib"
	"strings"
)

func findRobot(grid *lib.Grid) lib.Vec2 {
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if grid.Get(x, y) == '@' {
				return lib.Vec2{X: x, Y: y}
			}
		}
	}
	panic("did not find robot")
}

func moveRobot(pos, direction lib.Vec2, grid *lib.Grid) lib.Vec2 {
	return pos
}

func sumBoxLocations(grid *lib.Grid) int {
	sum := 0
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if grid.Get(x, y) == 'O' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func predictRobotMovements(input string) int {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	grid := lib.NewGrid(parts[0])
	movements := strings.ReplaceAll(parts[1], "\n", "")

	robot := findRobot(grid)

	for _, movement := range movements {
		switch movement {
		case '^':
			robot = moveRobot(robot, lib.Vec2{X: 0, Y: -1}, grid)
		case '<':
			robot = moveRobot(robot, lib.Vec2{X: -1, Y: 0}, grid)
		case '>':
			robot = moveRobot(robot, lib.Vec2{X: 1, Y: 0}, grid)
		case 'v':
			robot = moveRobot(robot, lib.Vec2{X: 0, Y: 1}, grid)
		}
	}

	return sumBoxLocations(grid)
}

func Part1() any {
	return "Not implemented"
}

func Part2() any {
	return "Not implemented"
}
