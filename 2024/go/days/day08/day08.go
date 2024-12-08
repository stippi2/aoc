package day08

import (
	"aoc/2024/go/lib"
)

func countAntiNodes(input string, repeat bool) int {
	grid := lib.NewGrid(input)
	antennaKindToPositions := make(map[byte][]lib.Vec2)
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			kind := grid.Get(x, y)
			if kind != '.' {
				positions := antennaKindToPositions[kind]
				antennaKindToPositions[kind] = append(positions, lib.Vec2{X: x, Y: y})
			}
		}
	}

	antiNodePositions := make(map[lib.Vec2]bool)
	for _, positions := range antennaKindToPositions {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				xVec := positions[i].X - positions[j].X
				yVec := positions[i].Y - positions[j].Y

				x1 := positions[i].X + xVec
				y1 := positions[i].Y + yVec
				x2 := positions[j].X - xVec
				y2 := positions[j].Y - yVec

				if !repeat {
					if x1 >= 0 && x1 < grid.Width() && y1 >= 0 && y1 < grid.Height() {
						antiNodePositions[lib.Vec2{X: x1, Y: y1}] = true
					}
					if x2 >= 0 && x2 < grid.Width() && y2 >= 0 && y2 < grid.Height() {
						antiNodePositions[lib.Vec2{X: x2, Y: y2}] = true
					}
				} else {
					antiNodePositions[positions[i]] = true
					antiNodePositions[positions[j]] = true
					for {
						if x1 < 0 || x1 >= grid.Width() || y1 < 0 || y1 >= grid.Height() {
							break
						}
						antiNodePositions[lib.Vec2{X: x1, Y: y1}] = true
						x1 += xVec
						y1 += yVec
					}
					for {
						if x2 < 0 || x2 >= grid.Width() || y2 < 0 || y2 >= grid.Height() {
							break
						}
						antiNodePositions[lib.Vec2{X: x2, Y: y2}] = true
						x2 -= xVec
						y2 -= yVec
					}
				}
			}
		}
	}

	return len(antiNodePositions)
}

func Part1() interface{} {
	input, _ := lib.ReadInput(8)
	return countAntiNodes(input, false)
}

func Part2() interface{} {
	input, _ := lib.ReadInput(8)
	return countAntiNodes(input, true)
}
