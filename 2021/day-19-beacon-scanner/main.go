package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Position struct {
	x, y, z int
}

type Beacon struct {
	// position is relative to the owning Scanner
	position Position
	// distancesTo12NearestBeacons is formed by computing the distance to every other beacon of the same scanner,
	// sorting the resulting distances list, then using the 12 smallest distance values.
	// The result should be independent of the owning Scanner's location and rotation,
	// and can be used to find the same beacon in a different scanner.
	distancesTo12NearestBeacons string
}

type Scanner struct {
	beacons []Beacon
}

func main() {
}

func parseInput(input string) []*Scanner {
	var scanners []*Scanner
	var currentScanner *Scanner
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "---") {
			currentScanner = &Scanner{}
			scanners = append(scanners, currentScanner)
		} else if currentScanner != nil {
			coords := strings.Split(line, ",")
			if len(coords) != 3 {
				panic("invalid coords")
			}
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			z, _ := strconv.Atoi(coords[2])
			currentScanner.beacons = append(currentScanner.beacons, Beacon{position: Position{x, y, z}})
		}
	}
	return scanners
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
