package main

import (
	"io/ioutil"
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

func parseInput(input string) []Scanner {
	return nil
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
