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

func (v Volume) contains(p Position) bool {
	return v.min.x <= p.x && v.max.x >= p.x &&
		v.min.y <= p.y && v.max.y >= p.y &&
		v.min.z <= p.z && v.max.z >= p.z
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

// volume returns the minimum Volume which contains all beacons of the Scanner
func (s *Scanner) volume() Volume {
	v := Volume{
		min: Position{math.MaxInt32, math.MaxInt32, math.MaxInt32},
		max: Position{math.MinInt32, math.MinInt32, math.MinInt32},
	}

	for i := 0; i < len(s.beacons); i++ {
		p := s.beacons[i].position
		v.min.x = min(p.x, v.min.x)
		v.min.y = min(p.y, v.min.y)
		v.min.z = min(p.z, v.min.z)
		v.max.x = max(p.x, v.max.x)
		v.max.y = max(p.y, v.max.y)
		v.max.z = max(p.z, v.max.z)
	}
	return v
}

func (s *Scanner) getBeaconsInVolume(v Volume) []Beacon {
	var beacons []Beacon
	for i := 0; i < len(s.beacons); i++ {
		if v.contains(s.beacons[i].position) {
			beacons = append(beacons, s.beacons[i])
		}
	}
	return beacons
}

func (s *Scanner) rotations() []Scanner {
	rotatedScanners := make([]Scanner, 24)
	for _, beacon := range s.beacons {
		rotatedPositions := beacon.position.rotations()
		for r := 0; r < 24; r++ {
			p := rotatedPositions[r]
			rotatedScanners[r].appendBeacon(p.x, p.y, p.z)
		}
	}
	return rotatedScanners
}

type AlignmentInfo struct {
	rotationIndexA int
	rotationIndexB int
	offset         Position
}

func alignScanners(a, b *Scanner) (*AlignmentInfo, bool) {
	type match struct {
		beaconIndexA int
		beaconIndexB int
	}
	var matchingBeacons []match
	for beaconIndexA, beaconA := range a.beacons {
		for beaconIndexB, beaconB := range b.beacons {
			if beaconA.distancesToNearest == beaconB.distancesToNearest {
				matchingBeacons = append(matchingBeacons, match{beaconIndexA, beaconIndexB})
			}
		}
	}
	if len(matchingBeacons) < 12 {
		return nil, false
	}
	for rotationIndexA, rotationA := range a.rotations() {
		for rotationIndexB, rotationB := range b.rotations() {
			for _, matching := range matchingBeacons {
				// If we found an alignment, we can transform both scanners to have the matching beacon as origin,
				// then form the intersecting volume, and all beacons within the intersection need to match
				originA := rotationA.beacons[matching.beaconIndexA].position
				translatedA := rotationA.translateBy(originA.x, originA.y, originA.z)

				originB := rotationB.beacons[matching.beaconIndexB].position
				translatedB := rotationB.translateBy(originB.x, originB.y, originB.z)

				intersection, overlap := translatedB.volume().intersect(translatedA.volume())
				if !overlap {
					// Should not be possible when we translated both scanners to the same origin
					continue
				}
				beaconsInIntersectionA := translatedA.getBeaconsInVolume(intersection)
				if len(beaconsInIntersectionA) < 12 {
					continue
				}
				beaconsInIntersectionB := translatedB.getBeaconsInVolume(intersection)

				if containsSameBeacons(beaconsInIntersectionA, beaconsInIntersectionB) {
					return &AlignmentInfo{
						rotationIndexA,
						rotationIndexB,
						Position{
							x: originB.x - originA.x,
							y: originB.y - originA.y,
							z: originB.z - originA.z,
						},
					}, true
				}
			}
		}
	}
	return nil, false
}


func main() {
}

func sortBeacons(a []Beacon) {
	sort.Slice(a, func (i, j int) bool {
		if a[i].position.x < a[j].position.x {
			return true
		}
		if a[i].position.x == a[j].position.x {
			if a[i].position.y < a[j].position.y {
				return true
			}
			if a[i].position.y == a[j].position.y {
				if a[i].position.z < a[j].position.z {
					return true
				}
				return false
			}
			return false
		}
		return false
	})
}

func containsSameBeacons(a, b []Beacon) bool {
	if len(a) != len(b) {
		return false
	}
	sortBeacons(a)
	sortBeacons(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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
