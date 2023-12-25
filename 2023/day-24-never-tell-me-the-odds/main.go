package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Vector struct {
	x, y, z float64
}

type Hailstone struct {
	position Vector
	velocity Vector
}

func calculateIntersection2D(a, b Hailstone) (Vector, bool) {
	var m1, m2, b1, b2 float64
	if a.velocity.x != 0 {
		m1 = a.velocity.y / a.velocity.x
		b1 = a.position.y - m1*a.position.x
	} else {
		m1 = math.MaxFloat64 // Big number to represent a vertical line
		b1 = a.position.x
	}
	if b.velocity.x != 0 {
		m2 = b.velocity.y / b.velocity.x
		b2 = b.position.y - m2*b.position.x
	} else {
		m2 = math.MaxFloat64 // Big number to represent a vertical line
		b2 = b.position.x
	}
	// Check if the two paths a re parallel
	if math.Abs(m1-m2) < 0.0000001 {
		return Vector{}, false
	}
	intersectionX := (b2 - b1) / (m1 - m2)
	intersectionY := m1*intersectionX + b1
	intersection := Vector{intersectionX, intersectionY, 0}

	if !isFutureIntersection(a, intersection) || !isFutureIntersection(b, intersection) {
		return Vector{}, false
	}

	return intersection, true
}

func isFutureIntersection(h Hailstone, intersection Vector) bool {
	directionToIntersection := Vector{x: intersection.x - h.position.x, y: intersection.y - h.position.y}
	scalarProduct := directionToIntersection.x*h.velocity.x + directionToIntersection.y*h.velocity.y

	return scalarProduct > 0
}

func partOne(hailstones []Hailstone, minX, maxX float64) int {
	// Find the intersections of all pairs of hailstones
	count := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if intersection, ok := calculateIntersection2D(hailstones[i], hailstones[j]); ok {
				if intersection.x >= minX && intersection.x <= maxX && intersection.y >= minX && intersection.y <= maxX {
					count++
				}
			}
		}
	}
	return count
}

func partTwo() int {
	return 0
}

func main() {
	now := time.Now()
	hailstones := parseInput(loadInput("puzzle-input.txt"))
	part1 := partOne(hailstones, 200000000000000, 400000000000000)
	part2 := partTwo()
	duration := time.Since(now)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", duration)
}

func parseInput(input string) []Hailstone {
	lines := strings.Split(input, "\n")
	hailstones := make([]Hailstone, len(lines))
	for i, line := range lines {
		var x, y, z, vx, vy, vz int
		_, _ = fmt.Sscanf(line, "%d, %d, %d @ %d, %d, %d", &x, &y, &z, &vx, &vy, &vz)
		hailstones[i] = Hailstone{
			position: Vector{float64(x), float64(y), float64(z)},
			velocity: Vector{float64(vx), float64(vy), float64(vz)},
		}
	}
	return hailstones
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
