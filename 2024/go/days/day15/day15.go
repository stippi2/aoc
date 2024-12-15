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

func moveThing(from, to lib.Vec2, grid *lib.Grid) {
	grid.Set(to.X, to.Y, grid.Get(from.X, from.Y))
	grid.Set(from.X, from.Y, '.')
}

func moveRobot(robot, direction lib.Vec2, grid *lib.Grid) lib.Vec2 {
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
	input, _ := lib.ReadInput(15)
	return predictRobotMovements(input)
}

func Part2() any {
	return "Not implemented"
}
