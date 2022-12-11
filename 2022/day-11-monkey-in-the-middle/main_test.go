package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_add(t *testing.T) {
	fn := toFactorizedNumber(45)
	fn.add(4)
	assert.True(t, fn.isDivisibleBy(7))
}
