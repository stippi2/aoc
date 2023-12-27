package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type SandSlab struct {
	x1, y1, z1 int
	x2, y2, z2 int
}

func (s *SandSlab) lower() SandSlab {
	return SandSlab{
		x1: s.x1,
		y1: s.y1,
		z1: s.z1 - 1,
		x2: s.x2,
		y2: s.y2,
		z2: s.z2 - 1,
	}
}

func (s *SandSlab) raise() SandSlab {
	return SandSlab{
		x1: s.x1,
		y1: s.y1,
		z1: s.z1 + 1,
		x2: s.x2,
		y2: s.y2,
		z2: s.z2 + 1,
	}
}

func (s *SandSlab) intersects(other SandSlab) bool {
	return s.x1 <= other.x2 && s.x2 >= other.x1 &&
		s.y1 <= other.y2 && s.y2 >= other.y1 &&
		s.z1 <= other.z2 && s.z2 >= other.z1
}

func sort(slabs []SandSlab) {
	slices.SortFunc(slabs, func(a, b SandSlab) int {
		return a.z1 - b.z1
	})
}

func settle(slabs []SandSlab, index int) {
	for i := index; i < len(slabs); i++ {
		for {
			if slabs[i].z1 == 1 {
				break
			}
			// We could also find the first slab below the slab at i where the x-y areas overlap
			slabs[i] = slabs[i].lower()
			foundIntersection := false
			for j := i - 1; j >= 0; j-- {
				if slabs[i].intersects(slabs[j]) {
					foundIntersection = true
					break
				}
			}
			if foundIntersection {
				slabs[i] = slabs[i].raise()
				break
			}
		}
	}
}

func testRemoval(slabs []SandSlab, index int) bool {
	slabsCopy := make([]SandSlab, len(slabs)-1)
	copy(slabsCopy, slabs[:index])
	copy(slabsCopy[index:], slabs[index+1:])
	height := slabsCopy[len(slabsCopy)-1].z2
	settle(slabsCopy, index)
	return slabsCopy[len(slabsCopy)-1].z2 == height
}

func partOne(slabs []SandSlab) int {
	settle(slabs, 0)
	count := 0
	for i := 0; i < len(slabs); i++ {
		if testRemoval(slabs, i) {
			count++
		}
	}
	return count
}

func partTwo() int {
	return 0
}

func main() {
	now := time.Now()
	slabs := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(slabs)
	part2 := partTwo()
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseLine(line string) SandSlab {
	var slab SandSlab
	_, _ = fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &slab.x1, &slab.y1, &slab.z1, &slab.x2, &slab.y2, &slab.z2)
	return slab
}

func parseInput(input string) []SandSlab {
	lines := strings.Split(input, "\n")
	slabs := make([]SandSlab, len(lines))
	for i, line := range lines {
		slabs[i] = parseLine(line)
	}
	sort(slabs)
	return slabs
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
