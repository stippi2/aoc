package day07

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	calibrations := parseInput([]string{"42: 40 2", "3: 1 2", "79325232924618: 907 466 857 34 3 618"})
	expected := []Calibration{
		{result: 42, sequence: []int64{40, 2}},
		{result: 3, sequence: []int64{1, 2}},
		{result: int64(79325232924618), sequence: []int64{907, 466, 857, 34, 3, 618}},
	}
	assert.Equal(t, expected, calibrations)
}

const example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func Test_sumValidCalibrations(t *testing.T) {
	inputLines := strings.Split(example, "\n")
	assert.Equal(t, int64(3749), sumValidCalibrations(inputLines, false))
}

func Test_sumValidCalibrationsPart2(t *testing.T) {
	inputLines := strings.Split(example, "\n")
	assert.Equal(t, int64(11387), sumValidCalibrations(inputLines, true))
}
