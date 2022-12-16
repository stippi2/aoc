package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Pos struct {
	x, y int
}

type Sensor struct {
	pos    Pos
	beacon Pos
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

type Line struct {
	x1, x2 int
}

func (l *Line) intersects(other *Line) bool {
	return !(other.x2 < l.x1-1 || other.x1 > l.x2+1)
}

func (l *Line) union(other *Line) {
	if l.intersects(other) {
		l.x1 = min(l.x1, other.x1)
		l.x2 = max(l.x2, other.x2)
	}
}

func (l *Line) intersect(other *Line) {
	if l.intersects(other) {
		l.x1 = max(l.x1, other.x1)
		l.x2 = min(l.x2, other.x2)
	}
}

func (s *Sensor) minMaxX(y int) (*Line, bool) {
	distance := abs(s.beacon.x-s.pos.x) + abs(s.beacon.y-s.pos.y)
	distance -= abs(y - s.pos.y)
	if distance < 0 {
		return nil, false
	}
	return &Line{s.pos.x - distance, s.pos.x + distance}, true
}

func intersect(lines []*Line, line Line) []*Line {
	var newLines []*Line
	for _, l := range lines {
		if line.intersects(l) {
			intersection := &Line{line.x1, line.x2}
			intersection.intersect(l)
			newLines = append(newLines, intersection)
		}
	}
	return newLines
}

func knownSections(y int, sensors []Sensor) []*Line {
	var lines []*Line
	for _, sensor := range sensors {
		newLine, insideRange := sensor.minMaxX(y)
		if insideRange {
			newLines := []*Line{newLine}
			for _, line := range lines {
				if newLine.intersects(line) {
					newLine.union(line)
				} else {
					newLines = append(newLines, line)
				}
			}
			lines = newLines
		}
	}
	return lines
}

func emptyPositionsOnLine(y int, sensors []Sensor) int {
	lines := knownSections(y, sensors)
	allBeacons := make(map[Pos]bool)
	for _, sensor := range sensors {
		allBeacons[sensor.beacon] = true
	}
	emptyPositions := 0
	for _, line := range lines {
		emptyPositions += line.x2 - line.x1 + 1
		for beacon := range allBeacons {
			if beacon.y == y {
				if line.x1 <= beacon.x && line.x2 >= beacon.x {
					emptyPositions--
				}
			}
		}
	}
	return emptyPositions
}

func main() {
	sensors := parseInput(loadInput("puzzle-input.txt"))
	start := time.Now()
	emptyPositions := emptyPositionsOnLine(2000000, sensors)
	durationPart1 := time.Now().Sub(start)
	fmt.Printf("empty positions on y = 2000000: %v (%v)\n", emptyPositions, durationPart1)

	start = time.Now()
	for y := 0; y <= 4000000; y++ {
		lines := intersect(knownSections(y, sensors), Line{0, 4000000})
		if len(lines) != 1 {
			sort.Slice(lines, func(i, j int) bool {
				return lines[i].x1 < lines[j].x1
			})
			x := lines[0].x2 + 1
			durationPart2 := time.Now().Sub(start)
			fmt.Printf("distress signal location: %vx%v, tuning frequency: %v (%v)\n", x, y, x*4000000+y, durationPart2)
			//			break
		}
	}
}

func parseInput(input string) []Sensor {
	lines := strings.Split(input, "\n")
	sensors := make([]Sensor, len(lines))
	for i, line := range lines {
		matches, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensors[i].pos.x, &sensors[i].pos.y, &sensors[i].beacon.x, &sensors[i].beacon.y)
		if matches != 4 || err != nil {
			panic(fmt.Sprintf("failed to parse sensor line '%s': %v", line, err))
		}
	}
	return sensors
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
