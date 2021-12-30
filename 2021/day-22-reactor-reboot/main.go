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
	}
	return result, result.isValid()
}

func (v Volume) splitAlongX(x int) []Volume {
	if x <= v.min.x || x >= v.max.x {
		return []Volume{v}
	}
	return []Volume{
		{
			on: v.on,
			min: Position{v.min.x, v.min.y, v.min.z},
			max: Position{x, v.max.y, v.max.z},
		},
		{
			on: v.on,
			min: Position{x + 1, v.min.y, v.min.z},
			max: Position{v.max.x, v.max.y, v.max.z},
		},
	}
}

func (v Volume) splitAlongY(y int) []Volume {
	if y <= v.min.y || y >= v.max.y {
		return []Volume{v}
	}
	return []Volume{
		{
			on: v.on,
			min: Position{v.min.x, v.min.y, v.min.z},
			max: Position{v.max.x, y, v.max.z},
		},
		{
			on: v.on,
			min: Position{v.min.x, y + 1, v.min.z},
			max: Position{v.max.x, v.max.y, v.max.z},
		},
	}
}

func (v Volume) splitAlongZ(z int) []Volume {
	if z <= v.min.z || z >= v.max.z {
		return []Volume{v}
	}
	return []Volume{
		{
			on: v.on,
			min: Position{v.min.x, v.min.y, v.min.z},
			max: Position{v.max.x, v.max.y, z},
		},
		{
			on: v.on,
			min: Position{v.min.x, v.min.y, z + 1},
			max: Position{v.max.x, v.max.y, v.max.z},
		},
	}
}

func (v Volume) contains(other Volume) bool {
	return v.min.x <= other.min.x &&
		v.min.y <= other.min.y &&
		v.min.z <= other.min.z &&
		v.max.x >= other.max.x &&
		v.max.y >= other.max.y &&
		v.max.z >= other.max.z
}

func splitAtX(volumes []Volume, x int) []Volume {
	var result []Volume
	for _, volume := range volumes {
		result = append(result, volume.splitAlongX(x)...)
	}
	return result
}

func splitAtY(volumes []Volume, y int) []Volume {
	var result []Volume
	for _, volume := range volumes {
		result = append(result, volume.splitAlongY(y)...)
	}
	return result
}

func splitAtZ(volumes []Volume, z int) []Volume {
	var result []Volume
	for _, volume := range volumes {
		result = append(result, volume.splitAlongZ(z)...)
	}
	return result
}

func (v Volume) splitAt(intersection Volume) (volumes []Volume) {
	volumes = []Volume{v}
	volumes = splitAtX(volumes, intersection.min.x)
	volumes = splitAtX(volumes, intersection.max.x)
	volumes = splitAtY(volumes, intersection.min.y)
	volumes = splitAtY(volumes, intersection.max.y)
	volumes = splitAtZ(volumes, intersection.min.z)
	volumes = splitAtZ(volumes, intersection.max.z)
	return
}

func (v Volume) union(other Volume) (volumes []Volume) {
	if v.contains(other) {
		return []Volume{v}
	}
	if other.contains(v) {
		return []Volume{other}
	}
	intersection, valid := v.intersect(other)
	if !valid {
		return []Volume{v, other}
	}
	uniqueVolumes := make(map[Volume]bool)
	for _, volume := range v.splitAt(intersection) {
		uniqueVolumes[volume] = true
	}
	for _, volume := range other.splitAt(intersection) {
		uniqueVolumes[volume] = true
	}
	for volume := range uniqueVolumes {
		volumes = append(volumes, volume)
	}
	return
}

func (v Volume) subtract(other Volume) (volumes []Volume) {
	if !v.contains(other) {
		return []Volume{v}
	}
	if other.contains(v) {
		return []Volume{}
	}
	intersection, valid := v.intersect(other)
	if !valid {
		return []Volume{v, other}
	}
	uniqueVolumes := make(map[Volume]bool)
	for _, volume := range v.splitAt(intersection) {
		uniqueVolumes[volume] = true
	}
	for _, volume := range other.splitAt(intersection) {
		delete(uniqueVolumes, volume)
	}
	for volume := range uniqueVolumes {
		volumes = append(volumes, volume)
	}
	return
}

func countCubes(volumes []Volume, within Volume) int {
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

func rebootSequence(sequence []Volume) (volumes []Volume) {
	for _, step := range sequence {
		if len(volumes) == 0 && step.on {
			volumes = append(volumes, step)
			continue
		}
		uniqueVolumes := make(map[Volume]bool)
		for _, volume := range volumes {
			var parts []Volume
			if step.on {
				parts = volume.union(step)
			} else {
				parts = volume.subtract(step)
			}
			for _, part := range parts {
				uniqueVolumes[part] = true
			}
		}
		volumes = nil
		for volume := range uniqueVolumes {
			volumes = append(volumes, volume)
		}
	}
	return
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
