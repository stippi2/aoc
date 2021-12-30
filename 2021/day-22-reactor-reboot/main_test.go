package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var smallExample = `on x=10..12,y=10..12,z=10..12
on x=11..13,y=11..13,z=11..13
off x=9..11,y=9..11,z=9..11
on x=10..10,y=10..10,z=10..10`

func Test_parseInput(t *testing.T) {
	expected := []Volume{
		{on: true, min: Position{10, 10, 10}, max: Position{12, 12, 12}},
		{on: true, min: Position{11, 11, 11}, max: Position{13, 13, 13}},
		{on: false, min: Position{9, 9, 9}, max: Position{11, 11, 11}},
		{on: true, min: Position{10, 10, 10}, max: Position{10, 10, 10}},
	}
	actual := parseInput(smallExample)
	assert.Equal(t, expected, actual)
}

var largerExample = `on x=-20..26,y=-36..17,z=-47..7
on x=-20..33,y=-21..23,z=-26..28
on x=-22..28,y=-29..23,z=-38..16
on x=-46..7,y=-6..46,z=-50..-1
on x=-49..1,y=-3..46,z=-24..28
on x=2..47,y=-22..22,z=-23..27
on x=-27..23,y=-28..26,z=-21..29
on x=-39..5,y=-6..47,z=-3..44
on x=-30..21,y=-8..43,z=-13..34
on x=-22..26,y=-27..20,z=-29..19
off x=-48..-32,y=26..41,z=-47..-37
on x=-12..35,y=6..50,z=-50..-2
off x=-48..-32,y=-32..-16,z=-15..-5
on x=-18..26,y=-33..15,z=-7..46
off x=-40..-22,y=-38..-28,z=23..41
on x=-16..35,y=-41..10,z=-47..6
off x=-32..-23,y=11..30,z=-14..3
on x=-49..-5,y=-3..45,z=-29..18
off x=18..30,y=-20..-8,z=-3..13
on x=-41..9,y=-7..43,z=-33..15
on x=-54112..-39298,y=-85059..-49293,z=-27449..7877
on x=967..23432,y=45373..81175,z=27513..53682`

func Test_largeExamplePartOne(t *testing.T) {
	sequence := parseInput(largerExample)
	volumes := rebootSequence(sequence)
	cubes := countCubesInVolume(volumes, Volume{
		min: Position{-50, -50, -50},
		max: Position{50, 50, 50},
	})
	assert.Equal(t, 590784, cubes)
}

func Test_smallExamplePartOne(t *testing.T) {
	sequence := parseInput(smallExample)
	volumes := rebootSequence(sequence)
	cubes := countCubes(volumes)
	assert.Equal(t, 39, cubes)
}

func Test_smallExamplePartOneSteps(t *testing.T) {
	sequence := parseInput(smallExample)
	volumes := applyRebootStep(nil, sequence[0])
	assert.Equal(t, 27, countCubes(volumes))
	volumes = applyRebootStep(volumes, sequence[1])
	assert.Equal(t, 27 + 19, countCubes(volumes))
	volumes = applyRebootStep(volumes, sequence[2])
	assert.Equal(t, 27 + 19 - 8, countCubes(volumes))
	volumes = applyRebootStep(volumes, sequence[3])
	assert.Equal(t, 27 + 19 - 8 + 1, countCubes(volumes))
}

func Test_intersect(t *testing.T) {
	tests := []struct {
		v1, v2 Volume
		expectedIntersection Volume
		expectedValid bool
	}{
		{
			v1: Volume{
				min: Position{0, 0, 0},
				max: Position{30, 30, 30},
			},
			v2: Volume{
				min: Position{10, 10, 10},
				max: Position{20, 20, 20},
			},
			expectedIntersection: Volume{
				min: Position{10, 10, 10},
				max: Position{20, 20, 20},
			},
			expectedValid: true,
		},
		{
			v1: Volume{
				min: Position{0, 0, 0},
				max: Position{0, 0, 0},
			},
			v2: Volume{
				min: Position{0, 0, 0},
				max: Position{0, 0, 0},
			},
			expectedIntersection: Volume{
				min: Position{0, 0, 0},
				max: Position{0, 0, 0},
			},
			expectedValid: true,
		},
		{
			v1: Volume{
				min: Position{0, 0, 0},
				max: Position{10, 10, 10},
			},
			v2: Volume{
				min: Position{10, 10, 10},
				max: Position{20, 20, 20},
			},
			expectedIntersection: Volume{
				min: Position{10, 10, 10},
				max: Position{10, 10, 10},
			},
			expectedValid: true,
		},
		{
			v1: Volume{
				min: Position{0, 0, 0},
				max: Position{10, 11, 12},
			},
			v2: Volume{
				min: Position{11, 12, 13},
				max: Position{30, 30, 30},
			},
			expectedValid: false,
		},
	}
	for _, test := range tests {
		actualIntersection, actualValid := test.v1.intersect(test.v2)
		assert.Equal(t, test.expectedValid, actualValid)
		if actualValid {
			assert.Equal(t, test.expectedIntersection, actualIntersection)
		}
	}
}

func Test_subtract(t *testing.T) {
	volume := Volume{
		min: Position{0, 0, 0},
		max: Position{30, 30, 30},
	}
	tests := []struct {
		subtract Volume
		expected []Volume
	}{
		{
			Volume{
				min: Position{20, 0, 0},
				max: Position{30, 30, 30},
			},
			[]Volume{
				{
					min: Position{0, 0, 0},
					max: Position{19, 30, 30},
				},
			},
		},
		{
			Volume{
				min: Position{0, 0, 0},
				max: Position{10, 30, 30},
			},
			[]Volume{
				{
					min: Position{11, 0, 0},
					max: Position{30, 30, 30},
				},
			},
		},
		{
			Volume{
				min: Position{0, 0, 0},
				max: Position{30, 30, 30},
			},
			[]Volume{},
		},
	}
	for _, test := range tests {
		actual := volume.subtract(test.subtract)
		assert.Equal(t, test.expected, actual)
	}
}

func Test_union(t *testing.T) {
	volume := Volume{
		min: Position{0, 0, 0},
		max: Position{30, 30, 30},
	}
	tests := []struct {
		other Volume
		expected []Volume
	}{
		{
			Volume{
				min: Position{10, 10, 10},
				max: Position{30, 30, 30},
			},
			[]Volume{
				{
					min: Position{0, 0, 0},
					max: Position{30, 30, 30},
				},
			},
		},
		{
			Volume{
				min: Position{30, 30, 30},
				max: Position{30, 30, 30},
			},
			[]Volume{
				{
					min: Position{0, 0, 0},
					max: Position{30, 30, 30},
				},
			},
		},
		{
			Volume{
				min: Position{31, 31, 31},
				max: Position{31, 31, 31},
			},
			[]Volume{
				{
					min: Position{0, 0, 0},
					max: Position{30, 30, 30},
				},
				{
					min: Position{31, 31, 31},
					max: Position{31, 31, 31},
				},
			},
		},
	}
	for _, test := range tests {
		actual := volume.union(test.other)
		assert.Equal(t, test.expected, actual)
	}
}

func Test_subtractVolume(t *testing.T) {
	volume := Volume{
		min: Position{0, 0, 0},
		max: Position{30, 30, 30},
	}
	intersection := Volume{
		min: Position{10, 10, 10},
		max: Position{20, 20, 20},
	}
	actual := volume.subtract(intersection)
	for i, v1 := range actual {
		for j, v2 := range actual {
			if i == j {
				continue
			}
			_, intersects := v1.intersect(v2)
			assert.False(t, intersects)
		}
	}
	assert.Equal(t, 31*31*31 - 11*11*11, countCubes(actual))
}
