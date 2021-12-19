package main

import (
	"testing"
)

func Test_aim(t *testing.T) {
	target := Target{
		minX: 20,
		maxX: 30,
		minY: -10,
		maxY: -5,
	}
	aim(target)
}
