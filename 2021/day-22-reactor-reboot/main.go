package main

import (
	"fmt"
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

func (v Volume) width() int {
	return v.max.x - v.min.x + 1
}

func (v Volume) height() int {
	return v.max.y - v.min.y + 1
}

func (v Volume) depth() int {
	return v.max.z - v.min.z + 1
}
func (v Volume) intersect(other Volume) (Volume, bool) {
	result := Volume{
		min: v.min.max(other.min),
		max: v.max.min(other.max),
		on: v.on,
	}
	return result, result.isValid()
}

func (v Volume) contains(other Volume) bool {
	return v.min.x <= other.min.x &&
		v.min.y <= other.min.y &&
		v.min.z <= other.min.z &&
		v.max.x >= other.max.x &&
		v.max.y >= other.max.y &&
		v.max.z >= other.max.z
}

func (v Volume) subtract(other Volume) (volumes []Volume) {
	i, intersects := v.intersect(other)
	if !intersects {
		return []Volume{v}
	}
	if other.contains(v) {
		return []Volume{}
	}
	front := Volume{
		min: v.min,
		max: Position{v.max.x, v.max.y, i.min.z - 1},
		on:  v.on,
	}
	back := Volume{
		min: Position{v.min.x, v.min.y, i.max.z + 1},
		max: v.max,
		on:  v.on,
	}
	top := Volume{
		min: Position{v.min.x, i.max.y + 1, i.min.z},
		max: Position{v.max.x, v.max.y, i.max.z},
		on:  v.on,
	}
	bottom := Volume{
		min: Position{v.min.x, v.min.y, i.min.z},
		max: Position{v.max.x, i.min.y - 1, i.max.z},
		on:  v.on,
	}
	left := Volume{
		min: Position{v.min.x, i.min.y, i.min.z},
		max: Position{i.min.x - 1, i.max.y, i.max.z},
		on:  v.on,
	}
	right := Volume{
		min: Position{i.max.x + 1, i.min.y, i.min.z},
		max: Position{v.max.x, i.max.y, i.max.z},
		on:  v.on,
	}
	for _, volume := range []Volume{front, back, top, bottom, left, right} {
		if volume.isValid() {
			volumes = append(volumes, volume)
		}
	}
	return
}

func countCubes(volumes []Volume) int {
	cubes := 0
	for _, volume := range volumes {
		cubes += volume.width() * volume.height() * volume.depth()
	}
	return cubes
}

func countCubesInVolume(volumes []Volume, within Volume) int {
	cubes := 0
	for _, volume := range volumes {
		intersection, valid := within.intersect(volume)
		if !valid {
			continue
		}
		cubes += intersection.width() * intersection.height() * intersection.depth()
	}
	return cubes
}

func applyRebootStep(volumes []Volume, step Volume) []Volume {
	if len(volumes) == 0 {
		if step.on {
			return []Volume{step}
		}
		return nil
	}
	var newVolumes []Volume
	for _, volume := range volumes {
		newVolumes = append(newVolumes, volume.subtract(step)...)
	}
	if step.on {
		newVolumes = append(newVolumes, step)
	}
	return newVolumes
}


func rebootSequence(sequence []Volume) (volumes []Volume) {
	for _, step := range sequence {
		volumes = applyRebootStep(volumes, step)
	}
	return
}

func main() {
	sequence := parseInput(loadInput("puzzle-input.txt"))
	volumes := rebootSequence(sequence)
	cubes := countCubesInVolume(volumes, Volume{
		min: Position{-50, -50, -50},
		max: Position{50, 50, 50},
	})
	fmt.Printf("cubes turned on in init volume: %v\n", cubes)
	fmt.Printf("cubes turned on in total: %v\n", countCubes(volumes))
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

func parseDimensions(dimensions []string, part int) Position {
	dimensions[0] = strings.TrimPrefix(dimensions[0], "x=")
	dimensions[1] = strings.TrimPrefix(dimensions[1], "y=")
	dimensions[2] = strings.TrimPrefix(dimensions[2], "z=")
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
