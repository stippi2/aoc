package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func Test_partOne(t *testing.T) {
	m := parseInput(input, false)
	assert.Equal(t, 16, partOne(m, 6))
}

func Test_partTwo(t *testing.T) {
	m := parseInput(input, true)

	//for {
	//	m.step()
	//
	//	var in string
	//	_, err := fmt.Scanln(&in)
	//	if err != nil {
	//		fmt.Println("Error reading input:", err)
	//		continue
	//	}
	//
	//	fmt.Print("\033[1A\033[K")
	//	fmt.Printf("%s\n", m)
	//}

	//	fmt.Printf("%s\n", m)
	//m.countTilesReachable(20)
	//fmt.Printf("%s\n", m)
	assert.Equal(t, 16733044, partTwo(m, 5000))
}
