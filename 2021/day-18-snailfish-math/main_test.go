package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func numberPair(left, right int) Node {
	return &Pair{ nil,&RegularNumber{left}, &RegularNumber{right}}
}

func Test_parseInput(t *testing.T) {
	type test struct {
		example      string
		expectedPair Node
	}

	var examples = []test{
		{"[1,2]", numberPair(1, 2)},
		{"[[1,2],3]", add(numberPair(1, 2), &RegularNumber{3})},
		{"[9,[8,7]]", add(&RegularNumber{9}, numberPair(8, 7)) },
		{"[[1,9],[8,5]]", add(numberPair(1, 9), numberPair(8, 5)) },
		{
			example: "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
			expectedPair: add(
				add(
					add(numberPair(1, 2), numberPair(3, 4)),
					add(numberPair(5, 6), numberPair(7, 8)),
				),
				&RegularNumber{9},
			),
		},
	}

	for _, e := range examples {
		assert.Equal(t, e.expectedPair, parseSnailfishNumber(e.example))
	}
}

func Test_explode(t *testing.T) {
	type test struct {
		example  string
		expected string
	}

	var explodeExamples = []test{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}

	for _, e := range explodeExamples {
		p := parseSnailfishNumber(e.example).(*Pair)
		p.Explode(0)
		assert.Equal(t, e.expected, fmt.Sprintf("%s", p))
	}
}

func Test_magnitude(t *testing.T) {
	type test struct {
		example  string
		expected int
	}

	var explodeExamples = []test{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}

	for _, e := range explodeExamples {
		p := parseSnailfishNumber(e.example).(*Pair)
		assert.Equal(t, e.expected, p.Magnitude())
	}
}

