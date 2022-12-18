package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Pos struct {
	x, y, z int
}

func (p *Pos) add(v Pos) {
	p.x += v.x
	p.y += v.y
	p.z += v.z
}

type Droplet struct {
	voxels     map[Pos]bool
	minX, maxX int
	minY, maxY int
	minZ, maxZ int
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

func (d *Droplet) set(p Pos) {
	d.voxels[p] = true
	d.minX = min(d.minX, p.x)
	d.minY = min(d.minY, p.y)
	d.minZ = min(d.minZ, p.z)
	d.maxX = max(d.maxX, p.x)
	d.maxY = max(d.maxY, p.y)
	d.maxZ = max(d.maxZ, p.z)
}

func (d *Droplet) isWithinBounds(p Pos) bool {
	return p.x >= d.minX && p.x <= d.maxX && p.y >= d.minY && p.y <= d.maxY && p.z >= d.minZ && p.z <= d.maxZ
}

func (d *Droplet) castRay(p Pos, v Pos) int {
	inside := false
	inOutFlips := 0
	for d.isWithinBounds(p) {
		if inside != d.voxels[p] {
			inside = !inside
			inOutFlips++
		}
		p.add(v)
	}
	if inside {
		inOutFlips++
	}
	return inOutFlips
}

func (d *Droplet) surfaceArea() int {
	area := 0
	for x := d.minX; x <= d.maxX; x++ {
		for y := d.minY; y <= d.maxY; y++ {
			area += d.castRay(Pos{x, y, d.minZ}, Pos{0, 0, 1})
		}
	}
	for x := d.minX; x <= d.maxX; x++ {
		for z := d.minZ; z <= d.maxZ; z++ {
			area += d.castRay(Pos{x, d.minY, z}, Pos{0, 1, 0})
		}
	}
	for y := d.minY; y <= d.maxY; y++ {
		for z := d.minZ; z <= d.maxZ; z++ {
			area += d.castRay(Pos{d.minX, y, z}, Pos{1, 0, 0})
		}
	}
	return area
}

func main() {
	droplet := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("surface area: %v\n", droplet.surfaceArea())
}

func parseInput(input string) *Droplet {
	droplet := &Droplet{
		voxels: make(map[Pos]bool),
		minX:   math.MaxInt,
		maxX:   math.MinInt,
		minY:   math.MaxInt,
		maxY:   math.MinInt,
		minZ:   math.MaxInt,
		maxZ:   math.MinInt,
	}
	for _, line := range strings.Split(input, "\n") {
		p := Pos{}
		matches, err := fmt.Sscanf(line, "%d,%d,%d", &p.x, &p.y, &p.z)
		if matches != 3 || err != nil {
			panic(fmt.Sprintf("failed to parse line '%s': %v", line, err))
		}
		droplet.set(p)
	}
	return droplet
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
