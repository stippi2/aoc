package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Position struct {
	x, y, z int
}

func (p *Position) distance(to Position) float64 {
	dx := to.x - p.x
	dy := to.y - p.y
	dz := to.z - p.z
	d1 := math.Sqrt(float64(dx * dx + dy * dy))
	return math.Sqrt(d1 * d1 + float64(dz * dz))
}

func (p *Position) rotations() []Position {
	return []Position{
		{p.x, p.y, p.z},
		{p.x, -p.y, -p.z},
		{p.x, p.z, -p.y},
		{p.x, -p.z, p.y},
		{-p.x, p.y, p.z},
		{-p.x, -p.y, -p.z},
		{-p.x, p.z, -p.y},
		{-p.x, -p.z, p.y},

		{p.x, p.y, p.z},
		{-p.x, p.y, -p.z},
		{p.z, p.y, -p.x},
		{-p.z, p.y, p.x},
		{p.x, -p.y, p.z},
		{-p.x, -p.y, -p.z},
		{p.z, -p.y, -p.x},
		{-p.z, -p.y, p.x},

		{p.x, p.y, p.z},
		{-p.x, -p.y, p.z},
		{p.y, -p.x, p.z},
		{-p.y, p.x, p.z},
		{p.x, p.y, -p.z},
		{-p.x, -p.y, -p.z},
		{p.y, -p.x, -p.z},
		{-p.y, p.x, -p.z},
	}
}

type Beacon struct {
	// position is relative to the owning Scanner
	position Position
	// distancesToNearest is formed by computing the distance to every other beacon of the same scanner,
	// sorting the resulting distances list, then using the 2 smallest distance values.
	// The result should be independent of the owning Scanner's location and rotation,
	// and can be used to find "same-looking" beacons in a different scanner.
	distancesToNearest string
}

type Scanner struct {
	beacons []Beacon
}

func (s *Scanner) appendBeacon(x, y, z int) {
	s.beacons = append(s.beacons, Beacon{position: Position{x, y, z}})
}

func (s *Scanner) setBeaconDistances() {
	for i := 0; i < len(s.beacons); i++ {
		var distances []float64
		a := &s.beacons[i]
		for j := 0; j < len(s.beacons); j++ {
			if i == j {
				continue
			}
			b := &s.beacons[j]
			distances = append(distances, a.position.distance(b.position))
		}
		sort.Float64s(distances)
		for d := 0; d < 2 && d < len(distances); d++ {
			a.distancesToNearest += fmt.Sprintf(",%.3f", distances[d])
		}
	}
}

func (s *Scanner) translateBy(x, y, z int) Scanner {
	scanner := Scanner{}
	for i := 0; i < len(s.beacons); i++ {
		p := s.beacons[i].position
		scanner.appendBeacon(p.x - x, p.y - y, p.z - z)
	}
	return scanner
}

// volume returns the minimum volume which contains all beacons of the Scanner
func (s *Scanner) volume() (min, max Position) {
	min.x = math.MaxInt32
	min.y = math.MaxInt32
	min.z = math.MaxInt32

	max.x = math.MinInt32
	max.y = math.MinInt32
	max.z = math.MinInt32

	for i := 0; i < len(s.beacons); i++ {
		p := s.beacons[i].position
		if p.x < min.x {
			min.x = p.x
		}
		if p.y < min.y {
			min.y = p.y
		}
		if p.z < min.z {
			min.z = p.z
		}
		if p.x > max.x {
			max.x = p.x
		}
		if p.y > max.y {
			max.y = p.y
		}
		if p.z > max.z {
			max.z = p.z
		}
	}
	return
}


/*
func (s *Scanner) rotations() []Scanner {
	rotations := make([]Scanner, 24)
	for _, beacon := range s.beacons {
		p := beacon.position
		for r := 0; r < 24; r++ {
			rotations[r].appendBeacon(???)
		}
	}
	return rotations
}
*/

func rotate90(a, b int) (int, int) {
	return -b, a
}

func main() {
}

func parseInput(input string) []Scanner {
	var scanners []Scanner
	var currentScanner *Scanner
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "---") {
			scanners = append(scanners, Scanner{})
			currentScanner = &scanners[len(scanners) - 1]
		} else if currentScanner != nil {
			coords := strings.Split(line, ",")
			if len(coords) != 3 {
				panic("invalid coords")
			}
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			z, _ := strconv.Atoi(coords[2])
			currentScanner.appendBeacon(x, y, z)
		}
	}
	return scanners
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
