package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

type Sensor struct {
	pos    Pos
	beacon Pos
}

func main() {
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
