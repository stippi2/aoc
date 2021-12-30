package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Position struct {
	x, y, z int
}

func (p *Position) max(other Position) Position {
	return Position{
		x: max(p.x, other.x),
		y: max(p.y, other.y),
		z: max(p.z, other.z),
	}
}

func (p *Position) min(other Position) Position {
	return Position{
		x: min(p.x, other.x),
		y: min(p.y, other.y),
		z: min(p.z, other.z),
	}
}

type Volume struct {
	min, max Position
	on bool
}

func (v Volume) isValid() bool {
	return v.min.x <= v.max.x && v.min.y <= v.max.y && v.min.z <= v.max.z
}

func (v Volume) intersect(other Volume) (Volume, bool) {
	result := Volume{
		min: v.min.max(other.min),
		max: v.max.min(other.max),
	}
	return result, result.isValid()
}

func main() {
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func parseDimensions(dimensions []string, part int) Position {
	dimensions[0] = strings.TrimPrefix(dimensions[0], "x=")
	dimensions[1] = strings.TrimPrefix(dimensions[0], "y=")
	dimensions[2] = strings.TrimPrefix(dimensions[0], "z=")
	xMinMax := strings.Split(dimensions[0], "..")
	yMinMax := strings.Split(dimensions[1], "..")
	zMinMax := strings.Split(dimensions[2], "..")
	x, _ := strconv.Atoi(xMinMax[part])
	y, _ := strconv.Atoi(yMinMax[part])
	z, _ := strconv.Atoi(zMinMax[part])
	return Position{x, y, z}
}

func parseInput(input string) (volumes []Volume) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}
		dimensions := strings.Split(parts[1], ",")
		if len(dimensions) != 3 {
			continue
		}
		volume := Volume{on: parts[0] == "on"}
		volume.min = parseDimensions(dimensions, 0)
		volume.max = parseDimensions(dimensions, 1)
		volumes = append(volumes, volume)
	}
	return
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
