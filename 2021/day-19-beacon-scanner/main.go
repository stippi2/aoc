package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Position struct {
	x, y, z int
}

func (p *Position) distance(to Position) int {
	dx := to.x - p.x
	dy := to.y - p.y
	dz := to.z - p.z
	d1 := dx * dx + dy * dy
	return d1 * d1 + dz * dz
}

type Beacon struct {
	// position is relative to the owning Scanner
	position Position
	// distancesTo12NearestBeacons is formed by computing the distance to every other beacon of the same scanner,
	// sorting the resulting distances list, then using the 12 smallest distance values.
	// The result should be independent of the owning Scanner's location and rotation,
	// and can be used to find the same beacon in a different scanner.
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
		var distances []int
		a := &s.beacons[i]
		for j := 0; j < len(s.beacons); j++ {
			if i == j {
				continue
			}
			b := &s.beacons[j]
			distances = append(distances, a.position.distance(b.position))
		}
		sort.Ints(distances)
		for d := 0; d < 12 && d < len(distances); d++ {
			a.distancesToNearest += ","+strconv.Itoa(d)
		}
	}
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
