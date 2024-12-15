package day15

import (
	"aoc/2024/go/lib"
	"fmt"
	"strings"
)

func parseInput(input string) (warehouse *lib.Grid, movements string) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	warehouse = lib.NewGrid(parts[0])
	movements = strings.ReplaceAll(parts[1], "\n", "")
	return
}

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

func moveThing(from, to lib.Vec2, grid *lib.Grid) {
	grid.Set(to.X, to.Y, grid.Get(from.X, from.Y))
	grid.Set(from.X, from.Y, '.')
}

func moveRobotSimple(robot, direction lib.Vec2, grid *lib.Grid) lib.Vec2 {
	test := robot.Add(direction)
	for {
		if grid.Get(test.X, test.Y) == '#' {
			// Hit a wall, robot cannot move
			return robot
		}
		if grid.Get(test.X, test.Y) == '.' {
			// Found a free space in the direction the robot wants to move
			break
		}
		test = test.Add(direction)
	}
	// Move whatever is before the robot to the empty space
	newRobot := robot.Add(direction)
	moveThing(newRobot, test, grid)
	// Move the robot
	moveThing(robot, newRobot, grid)
	return newRobot
}

func sumBoxLocations(warehouse *lib.Grid, box byte) int {
	sum := 0
	for y := 0; y < warehouse.Height(); y++ {
		for x := 0; x < warehouse.Width(); x++ {
			if warehouse.Get(x, y) == box {
				sum += 100*y + x
			}
		}
	}
	return sum
}

type MoveFunc func(robot, direction lib.Vec2, warehose *lib.Grid) lib.Vec2

func predictRobotMovements(input string, moveRobot MoveFunc, box byte) int {
	warehouse, movements := parseInput(input)
	robot := findRobot(warehouse)

	for _, movement := range movements {
		switch movement {
		case '^':
			robot = moveRobot(robot, lib.Vec2{X: 0, Y: -1}, warehouse)
		case '<':
			robot = moveRobot(robot, lib.Vec2{X: -1, Y: 0}, warehouse)
		case '>':
			robot = moveRobot(robot, lib.Vec2{X: 1, Y: 0}, warehouse)
		case 'v':
			robot = moveRobot(robot, lib.Vec2{X: 0, Y: 1}, warehouse)
		}
	}

	return sumBoxLocations(warehouse, box)
}

func Part1() any {
	input, _ := lib.ReadInput(15)
	return predictRobotMovements(input, moveRobotSimple, 'O')
}

func Part2() any {
	input, _ := lib.ReadInput(15)
	input = strings.ReplaceAll(input, "#", "##")
	input = strings.ReplaceAll(input, "O", "[]")
	input = strings.ReplaceAll(input, ".", "..")
	input = strings.ReplaceAll(input, "@", "@.")
	fmt.Print(input)
	return "Not implemented"
}
