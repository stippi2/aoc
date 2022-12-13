package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

func Test_parse(t *testing.T) {
	assert.Equal(t, []any{1, []any{21, []any{3, []any{4, []any{5, 621, 7}}}}, 8, 9}, parseItem("[1,[21,[3,[4,[5,621,7]]]],8,9]"))
}

func Test_part1(t *testing.T) {
	pairs := parseInput(exampleInput)

	assert.Equal(t, []Pair{
		{
			[]any{1, 1, 3, 1, 1},
			[]any{1, 1, 5, 1, 1},
		},
		{
			[]any{[]any{1}, []any{2, 3, 4}},
			[]any{[]any{1}, 4},
		},
		{
			[]any{9},
			[]any{[]any{8, 7, 6}},
		},
		{
			[]any{[]any{4, 4}, 4, 4},
			[]any{[]any{4, 4}, 4, 4, 4},
		},
		{
			[]any{7, 7, 7, 7},
			[]any{7, 7, 7},
		},
		{
			[]any{},
			[]any{3},
		},
		{
			[]any{[]any{[]any{}}},
			[]any{[]any{}},
		},
		{
			[]any{1, []any{2, []any{3, []any{4, []any{5, 6, 7}}}}, 8, 9},
			[]any{1, []any{2, []any{3, []any{4, []any{5, 6, 0}}}}, 8, 9},
		},
	}, pairs)

	assert.Equal(t, 13, sumIndicesOrderedPairs(pairs))
}

func Test_compare(t *testing.T) {
	assert.Equal(t, Sorted, compare([]any{1, 1, 3, 1, 1}, []any{1, 1, 5, 1, 1}))

	assert.Equal(t, Sorted, compare([]any{[]any{1}, []any{2, 3, 4}}, []any{[]any{1}, 4}))

	assert.Equal(t, Unsorted, compare([]any{9}, []any{[]any{8, 7, 6}}))

	assert.Equal(t, Sorted, compare([]any{[]any{4, 4}, 4, 4}, []any{[]any{4, 4}, 4, 4, 4}))

	assert.Equal(t, Unsorted, compare([]any{7, 7, 7, 7}, []any{7, 7, 7}))

	assert.Equal(t, Sorted, compare([]any{}, []any{3}))

	assert.Equal(t, Unsorted, compare([]any{[]any{[]any{}}}, []any{[]any{}}))

	assert.Equal(t, Unsorted, compare([]any{1, []any{2, []any{3, []any{4, []any{5, 6, 7}}}}, 8, 9}, []any{1, []any{2, []any{3, []any{4, []any{5, 6, 0}}}}, 8, 9}))
}
