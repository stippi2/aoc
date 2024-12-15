package day15

import (
	"aoc/2024/go/lib"
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

func moveRobotSimple(robot, direction lib.Vec2, warehouse *lib.Grid) lib.Vec2 {
	test := robot.Add(direction)
	for {
		if warehouse.Get(test.X, test.Y) == '#' {
			// Hit a wall, robot cannot move
			return robot
		}
		if warehouse.Get(test.X, test.Y) == '.' {
			// Found a free space in the direction the robot wants to move
			break
		}
		test = test.Add(direction)
	}
	// Move whatever is before the robot to the empty space
	newRobot := robot.Add(direction)
	moveThing(newRobot, test, warehouse)
	// Move the robot
	moveThing(robot, newRobot, warehouse)
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

func predictRobotMovements(warehouse *lib.Grid, movements string, moveRobot MoveFunc) {
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
}

func predictRobotMovementsFromInput(input string, moveRobot MoveFunc, box byte) int {
	warehouse, movements := parseInput(input)
	predictRobotMovements(warehouse, movements, moveRobot)
	return sumBoxLocations(warehouse, box)
}

func Part1() any {
	input, _ := lib.ReadInput(15)
	return predictRobotMovementsFromInput(input, moveRobotSimple, 'O')
}

func canMove(pos, direction lib.Vec2, warehouse *lib.Grid) bool {
	pos = pos.Add(direction)
	slot := warehouse.Get(pos.X, pos.Y)
	if slot == '.' {
		return true
	}
	if direction.Y != 0 {
		// up/down
		if slot == '[' {
			return canMove(pos, direction, warehouse) &&
				canMove(pos.Add(lib.Vec2{X: 1, Y: 0}), direction, warehouse)
		}
		if slot == ']' {
			return canMove(pos.Add(lib.Vec2{X: -1, Y: 0}), direction, warehouse) &&
				canMove(pos, direction, warehouse)
		}
	} else {
		// left/right
		for {
			pos = pos.Add(direction)
			if warehouse.Get(pos.X, pos.Y) == '#' {
				// Hit a wall, robot cannot move
				return false
			}
			if warehouse.Get(pos.X, pos.Y) == '.' {
				// Found a free space in the direction the robot wants to move
				return true
			}
		}
	}
	return false
}

func move(pos, direction lib.Vec2, warehouse *lib.Grid) {
	slot := warehouse.Get(pos.X, pos.Y)
	if slot == '[' {
		if direction.Y != 0 {
			otherHalf := pos.Add(lib.Vec2{X: 1, Y: 0})
			move(otherHalf.Add(direction), direction, warehouse)
			moveThing(otherHalf, otherHalf.Add(direction), warehouse)
		}
		move(pos.Add(direction), direction, warehouse)
		moveThing(pos, pos.Add(direction), warehouse)
	}
	if slot == ']' {
		if direction.Y != 0 {
			// Redirect to first half
			otherHalf := pos.Add(lib.Vec2{X: -1, Y: 0})
			move(otherHalf, direction, warehouse)
		} else {
			move(pos.Add(direction), direction, warehouse)
			moveThing(pos, pos.Add(direction), warehouse)
		}
	}
}

func moveRobotWide(robot, direction lib.Vec2, warehouse *lib.Grid) lib.Vec2 {
	if !canMove(robot, direction, warehouse) {
		return robot
	}
	// Move all the boxes recursively
	newRobot := robot.Add(direction)
	move(newRobot, direction, warehouse)
	// Move the robot
	moveThing(robot, newRobot, warehouse)
	return newRobot
}

func widenWarehouse(input string) string {
	input = strings.ReplaceAll(input, "#", "##")
	input = strings.ReplaceAll(input, "O", "[]")
	input = strings.ReplaceAll(input, ".", "..")
	input = strings.ReplaceAll(input, "@", "@.")
	return input
}

func Part2() any {
	input, _ := lib.ReadInput(15)
	input = widenWarehouse(input)
	return predictRobotMovementsFromInput(input, moveRobotWide, '[')
}
