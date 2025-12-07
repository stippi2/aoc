package day07

import (
	"aoc/2025/go/lib"
)

func countBeamSplits(input string) int {
	grid := lib.NewGrid(input)
	beamSplits := 0
	for y := 1; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			above := grid.Get(x, y-1)
			if above == '.' {
				continue
			}
			cell := grid.Get(x, y)
			if above == 'S' || above == '|' {
				switch cell {
				case '.':
					grid.Set(x, y, '|')
				case '^':
					if x > 0 && grid.Get(x-1, y) == '.' {
						grid.Set(x-1, y, '|')
					}
					if x < grid.Width()-1 && grid.Get(x+1, y) == '.' {
						grid.Set(x+1, y, '|')
					}
					beamSplits++
				}
			}
		}
	}

	return beamSplits
}

func Part1() any {
	input, _ := lib.ReadInput(7)
	return countBeamSplits(input)
}

func countTachyonTimelines(input string) int {
	grid := lib.NewGrid(input)

	timelines := map[int]int{}
	for x := 0; x < grid.Width(); x++ {
		if grid.Get(x, 0) == 'S' {
			timelines[x] = 1
			break
		}
	}

	for y := 1; y < grid.Height(); y++ {
		newTimelines := map[int]int{}
		for x, count := range timelines {
			cell := grid.Get(x, y)
			switch cell {
			case '.':
				newTimelines[x] += count
			case '^':
				if x > 0 {
					newTimelines[x-1] += count
				}
				if x < grid.Width()-1 {
					newTimelines[x+1] += count
				}
			}
		}
		timelines = newTimelines
	}

	sum := 0
	for _, value := range timelines {
		sum += value
	}
	return sum
}

func Part2() any {
	input, _ := lib.ReadInput(7)
	return countTachyonTimelines(input)
}
