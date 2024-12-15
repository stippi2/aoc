package day15

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleSmall = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

const exampleLarge = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

func Test_part1(t *testing.T) {
	assert.Equal(t, 2028, predictRobotMovementsFromInput(exampleSmall, moveRobotSimple, 'O'))
	assert.Equal(t, 10092, predictRobotMovementsFromInput(exampleLarge, moveRobotSimple, 'O'))
}

func Test_part2Left(t *testing.T) {
	warehouse, movements := parseInput(`##########
##......##
##.[][]@##
##......##
##########

<`)
	predictRobotMovements(warehouse, movements, moveRobotWide)

	assert.Equal(t, `##########
##......##
##[][]@.##
##......##
##########`, strings.TrimSpace(warehouse.String()))
}

func Test_part2Right(t *testing.T) {
	warehouse, movements := parseInput(`##########
##......##
##@[][].##
##......##
##########

>`)
	predictRobotMovements(warehouse, movements, moveRobotWide)

	assert.Equal(t, `##########
##......##
##.@[][]##
##......##
##########`, strings.TrimSpace(warehouse.String()))
}

func Test_part2Up(t *testing.T) {
	warehouse, movements := parseInput(`##########
##......##
##..[]..##
##.[]...##
##..@...##
##########

^`)
	predictRobotMovements(warehouse, movements, moveRobotWide)

	assert.Equal(t, `##########
##..[]..##
##.[]...##
##..@...##
##......##
##########`, strings.TrimSpace(warehouse.String()))
}

func Test_part2Down(t *testing.T) {
	warehouse, movements := parseInput(`##########
##..@...##
##..[]..##
##.[]...##
##......##
##########

v`)
	predictRobotMovements(warehouse, movements, moveRobotWide)

	fmt.Print(warehouse.String())

	assert.Equal(t, `##########
##......##
##..@...##
##..[]..##
##.[]...##
##########`, strings.TrimSpace(warehouse.String()))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 9021, predictRobotMovementsFromInput(widenWarehouse(exampleLarge), moveRobotWide, '['))
}
