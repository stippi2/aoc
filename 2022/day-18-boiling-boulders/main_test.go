package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func Test_part1(t *testing.T) {
	droplet := parseInput(exampleInput)
	assert.Equal(t, 64, droplet.surfaceArea())
}

func Test_getEnclosingVolume(t *testing.T) {
	droplet := parseInput(`1,1,1`)
	enclosing := droplet.getEnclosingVolume()
	assert.Equal(t, 0, enclosing.minX)
	assert.Equal(t, 0, enclosing.minY)
	assert.Equal(t, 0, enclosing.minZ)
	assert.Equal(t, 2, enclosing.maxX)
	assert.Equal(t, 2, enclosing.maxY)
	assert.Equal(t, 2, enclosing.maxZ)
	assert.Len(t, enclosing.voxels, 9+9+3+3+2)
}

func Test_part2(t *testing.T) {
	droplet := parseInput(exampleInput)
	droplet.fillAllPockets()
	assert.Equal(t, 58, droplet.surfaceArea())
}

func Test_part2_enclosingVolume(t *testing.T) {
	droplet := parseInput(exampleInput)

	enclosing := droplet.getEnclosingVolume()
	exteriorArea := enclosing.surfaceArea()
	exteriorArea -= 2 * (enclosing.maxX - enclosing.minX + 1) * (enclosing.maxY - enclosing.minY + 1)
	exteriorArea -= 2 * (enclosing.maxX - enclosing.minX + 1) * (enclosing.maxZ - enclosing.minZ + 1)
	exteriorArea -= 2 * (enclosing.maxY - enclosing.minY + 1) * (enclosing.maxZ - enclosing.minZ + 1)

	assert.Equal(t, 58, exteriorArea)
}
