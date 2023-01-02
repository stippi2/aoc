package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Pos struct {
	x, y, z int
}

func (p *Pos) add(v Pos) Pos {
	return Pos{p.x + v.x, p.y + v.y, p.z + v.z}
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

func (d *Droplet) countSurfacesAlongVector(p Pos, vector Pos) int {
	inside := false
	inOutFlips := 0
	for d.isWithinBounds(p) {
		if inside != d.voxels[p] {
			inside = !inside
			inOutFlips++
		}
		p = p.add(vector)
	}
	if inside {
		inOutFlips++
	}
	return inOutFlips
}

func (d *Droplet) fillPocketsAlongVector(p Pos, vector Pos) {
	inside := false
	for d.isWithinBounds(p) {
		if !inside && d.voxels[p] {
			inside = true
		} else if !d.voxels[p] && inside {
			d.fill(p)
		}
		p = p.add(vector)
	}
}

func (d *Droplet) fill(start Pos) {
	if d.voxels[start] {
		return
	}
	visited := make(map[Pos]bool)
	queue := []Pos{start}
	reachedOutside := false
	for len(queue) > 0 {
		tip := queue[0]
		if !d.isWithinBounds(tip) {
			reachedOutside = true
			break
		}
		visited[tip] = true
		queue = queue[1:]
		nextPositions := []Pos{
			tip.add(Pos{1, 0, 0}),
			tip.add(Pos{-1, 0, 0}),
			tip.add(Pos{0, 1, 0}),
			tip.add(Pos{0, -1, 0}),
			tip.add(Pos{0, 0, 1}),
			tip.add(Pos{0, 0, -1}),
		}
		for _, pos := range nextPositions {
			if !visited[pos] && !d.voxels[pos] {
				queue = append(queue, pos)
			}
		}
	}
	if !reachedOutside {
		// Fill the interior
		for pos := range visited {
			d.voxels[pos] = true
		}
	}
}

func (d *Droplet) getEnclosingVolume() *Droplet {
	result := &Droplet{
		voxels: make(map[Pos]bool),
		minX:   d.minX - 1,
		maxX:   d.maxX + 1,
		minY:   d.minY - 1,
		maxY:   d.maxY + 1,
		minZ:   d.minZ - 1,
		maxZ:   d.maxZ + 1,
	}
	queue := []Pos{{result.minX, result.minY, result.minZ}}
	for len(queue) > 0 {
		tip := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		result.voxels[tip] = true
		nextPositions := []Pos{
			tip.add(Pos{1, 0, 0}),
			tip.add(Pos{-1, 0, 0}),
			tip.add(Pos{0, 1, 0}),
			tip.add(Pos{0, -1, 0}),
			tip.add(Pos{0, 0, 1}),
			tip.add(Pos{0, 0, -1}),
		}
		for _, pos := range nextPositions {
			if result.isWithinBounds(pos) && !result.voxels[pos] && !d.voxels[pos] {
				queue = append(queue, pos)
			}
		}
	}
	return result
}

func (d *Droplet) sweepVolume(callback func(start, vector Pos)) {
	for x := d.minX; x <= d.maxX; x++ {
		for y := d.minY; y <= d.maxY; y++ {
			callback(Pos{x, y, d.minZ}, Pos{0, 0, 1})
		}
	}
	for x := d.minX; x <= d.maxX; x++ {
		for z := d.minZ; z <= d.maxZ; z++ {
			callback(Pos{x, d.minY, z}, Pos{0, 1, 0})
		}
	}
	for y := d.minY; y <= d.maxY; y++ {
		for z := d.minZ; z <= d.maxZ; z++ {
			callback(Pos{d.minX, y, z}, Pos{1, 0, 0})
		}
	}
}

func (d *Droplet) fillAllPockets() {
	for x := d.minX; x <= d.maxX; x++ {
		for y := d.minY; y <= d.maxY; y++ {
			d.fillPocketsAlongVector(Pos{x, y, d.minZ}, Pos{0, 0, 1})
		}
	}
}

func (d *Droplet) surfaceArea() int {
	area := 0
	d.sweepVolume(func(start, vector Pos) {
		area += d.countSurfacesAlongVector(start, vector)
	})
	return area
}

func main() {
	droplet := parseInput(loadInput("puzzle-input.txt"))
	start := time.Now()
	fmt.Printf("part 1, surface area including trapped air: %v (%v)\n", droplet.surfaceArea(), time.Since(start))

	start = time.Now()
	droplet.fillAllPockets()
	fmt.Printf("part 2, exterior surface area: %v (%v)\n", droplet.surfaceArea(), time.Since(start))

	droplet = parseInput(loadInput("puzzle-input.txt"))
	start = time.Now()
	enclosing := droplet.getEnclosingVolume()
	exteriorArea := enclosing.surfaceArea()
	exteriorArea -= 2 * (enclosing.maxX - enclosing.minX + 1) * (enclosing.maxY - enclosing.minY + 1)
	exteriorArea -= 2 * (enclosing.maxX - enclosing.minX + 1) * (enclosing.maxZ - enclosing.minZ + 1)
	exteriorArea -= 2 * (enclosing.maxY - enclosing.minY + 1) * (enclosing.maxZ - enclosing.minZ + 1)
	fmt.Printf("part 2, exterior surface area via enclosing surface: %v (%v)\n", exteriorArea, time.Since(start))
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
