package day05

import (
	"aoc/2025/go/lib"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func (r *Range) contains(v int) bool {
	return v >= r.start && v <= r.end
}

func (r *Range) length() int {
	return r.end - r.start + 1
}

func parseInput(input string) ([]Range, []int) {
	parts := strings.Split(input, "\n\n")

	rangeLines := strings.Split(parts[0], "\n")
	idLines := strings.Split(parts[1], "\n")

	ranges := make([]Range, len(rangeLines))
	ids := make([]int, len(idLines))

	for i, rangeLine := range rangeLines {
		fmt.Sscanf(rangeLine, "%d-%d", &ranges[i].start, &ranges[i].end)
	}

	for i, idLine := range idLines {
		ids[i], _ = strconv.Atoi(idLine)
	}

	return ranges, ids
}

func countFreshIds(input string) int {
	ranges, ids := parseInput(input)

	freshIds := 0

	for _, id := range ids {
		for i := range ranges {
			if ranges[i].contains(id) {
				freshIds++
				break
			}
		}
	}

	return freshIds
}

func countTotalFreshIds(input string) int {
	ranges, _ := parseInput(input)

	if len(ranges) == 0 {
		return 0
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return cmp.Compare(a.start, b.start)
	})

	count := 0

	merged := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if ranges[i].start <= merged.end {
			merged.end = lib.Max(ranges[i].end, merged.end)
		} else {
			count += merged.length()
			merged = ranges[i]
		}
	}

	count += merged.length()

	return count
}

func Part1() any {
	input, _ := lib.ReadInput(5)
	return countFreshIds(input)
}

func Part2() any {
	input, _ := lib.ReadInput(5)
	return countTotalFreshIds(input)
}
