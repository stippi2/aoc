package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = "2333133121414131402"

func Test_calculateChecksum(t *testing.T) {
	disk := make([]int, calculateDiskSize(example))
	expandBlocks(example, disk)
	compactDisk(disk)
	assert.Equal(t, int64(1928), calculateChecksum(disk))
}

func Test_calculateDiskSize(t *testing.T) {
	assert.Equal(t, int64(6), calculateDiskSize("123"))
}

func Test_expandBlocks(t *testing.T) {
	disk := make([]int, 6)
	expandBlocks("123", disk)
	assert.Equal(t, []int{0, -1, -1, 1, 1, 1}, disk)
}

func Test_expandBlocksExample(t *testing.T) {
	disk := make([]int, calculateDiskSize(example))
	expandBlocks(example, disk)
	assert.Equal(t, []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9}, disk)
}

func Test_compactDisk(t *testing.T) {
	disk := []int{0, -1, -1, 1, 1, 1, -1, 2, 2, -1, 3, -1}
	compactDisk(disk)
	assert.Equal(t, []int{0, 3, 2, 1, 1, 1, 2, -1, -1, -1, -1, -1}, disk)
}

func Test_defragmentDisk(t *testing.T) {
	disk := make([]int, calculateDiskSize(example))
	expandBlocks(example, disk)
	defragmentDisk(disk)
	assert.Equal(t, []int{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1}, disk)
}
