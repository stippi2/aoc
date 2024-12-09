package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = "2333133121414131402"

func Test_calculateChecksum(t *testing.T) {
	assert.Equal(t, 1928, calculateChecksum(example))
}
