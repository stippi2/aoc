package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

func Test_parseInput(t *testing.T) {
	root := parseInput(exampleInput).tip[0]
	assert.Equal(t, root.label, "AA")
	if assert.Len(t, root.connectedNodes, 3) {
		assert.Equal(t, root.connectedNodes[0].label, "DD")
		assert.Equal(t, root.connectedNodes[1].label, "II")
		assert.Equal(t, root.connectedNodes[2].label, "BB")
		assert.Equal(t, root.distance["DD"], 1)
		assert.Equal(t, root.distance["II"], 1)
		assert.Equal(t, root.distance["CC"], 2)
	}
}

func Test_part1(t *testing.T) {
	start := parseInput(exampleInput)
	assert.Equal(t, 1651, maximumPressureRelease(start, 30, 1))
}

func Test_part2(t *testing.T) {
	start := parseInput(exampleInput)
	// Elephant: DD(20), HH(22), EE(3)
	// Me: JJ(21), BB(13), CC(2)
	assert.Equal(t, 1707, maximumPressureRelease(start, 26, 2))
}
