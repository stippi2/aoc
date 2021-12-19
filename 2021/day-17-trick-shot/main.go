package main

import (
	"fmt"
)

type Probe struct {
	x         int
	y         int
	velocityX int
	velocityY int
}

type Target struct {
	minX int
	maxX int
	minY int
	maxY int
}

func (p *Probe) step() {
	p.x += p.velocityX
	p.y += p.velocityY
	if p.velocityX > 0 {
		p.velocityX--
	} else if p.velocityX < 0 {
		p.velocityX++
	}
	p.velocityY--
}

func (p* Probe) onTrackTo(t Target) bool {
	// TODO: Implement properly, currently checks if we are past the target already
	if p.x > t.maxX || p.y < t.minY {
		return false
	}
	return true
}

func (p* Probe) isWithin(t Target) bool {
	return p.x >= t.minX && p.x <= t.maxX && p.y >= t.minY && p.y <= t.maxY
}

func solveDistX(dist int) (velocity int, steps int) {
	for dist - velocity > 0 {
		if dist > 0 {
			velocity++
		} else if dist < 0 {
			velocity--
		}
		dist -= velocity
		steps++
	}
	return
}

func solveDistY(dist int) (velocity, steps int) {
	if dist < 0 {
		dist = -dist
	}
	for velocity < dist - 1 {
		velocity++
		steps += 2
	}
	steps += 2
	return
}

func aim(t Target) (vX int, vY int) {
	minVelocityX, minSteps := solveDistX(t.minX)
	maxVelocityX, maxSteps := solveDistX(t.maxX)

	fmt.Printf("min velocity X: %v, steps: %v\n", minVelocityX, minSteps)
	fmt.Printf("max velocity X: %v, steps: %v\n", maxVelocityX, maxSteps)

	maxVelocityY, maxStepsY := solveDistY(t.minY)
	minVelocityY, minStepsY := solveDistY(t.maxY)

	fmt.Printf("min velocity Y: %v, steps: %v\n", minVelocityY, minStepsY)
	fmt.Printf("max velocity Y: %v, steps: %v\n", maxVelocityY, maxStepsY)

	p := Probe{
		velocityX: (minVelocityX + maxVelocityX) / 2,
		velocityY: maxVelocityY,
	}

	for step := 0; step < maxStepsY; step++ {
		p.step()
		fmt.Printf("step: %2d, x: %2d, y: %2d\n", step, p.x, p.y)
	}

	fmt.Printf("distinct aim vectors: %v\n", (maxVelocityX - minVelocityX) + (maxVelocityY - minVelocityY))

	return
}

type Vector struct {
	x, y int
}

func possibleVectors(t Target) map[Vector]bool {
	vectorMap := map[Vector]bool{}

	minVelocityX, minSteps := solveDistX(t.minX)
	maxVelocityX := t.maxX

	fmt.Printf("min velocity X: %v, steps: %v\n", minVelocityX, minSteps)
	fmt.Printf("max velocity X: %v\n", maxVelocityX)

	minVelocityY := t.minY
	maxVelocityY, _ := solveDistY(t.minY)

	fmt.Printf("min velocity Y: %v\n", minVelocityY)
	fmt.Printf("max velocity Y: %v\n", maxVelocityY)

	// Brute-forcing the solutions space is really lame...
	for x := minVelocityX; x <= maxVelocityX; x++ {
		for y := minVelocityY; y <= maxVelocityY; y++ {
			probe := Probe{
				velocityX: x,
				velocityY: y,
			}
			for probe.onTrackTo(t) {
				probe.step()
				if probe.isWithin(t) {
					vectorMap[Vector{x, y}] = true
					break
				}
			}
		}
	}

	return vectorMap
}

func main() {
	// target area: x=137..171, y=-98..-73
	target := Target{
		minX: 137,
		maxX: 171,
		minY: -98,
		maxY: -73,
	}
	aim(target)

	vectors := possibleVectors(target)
	fmt.Printf("possible vectors: %v\n", len(vectors))
}
