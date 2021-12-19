package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type test struct {
		example      string
		expectedPair Node
	}

	var examples = []test{
		{"[1,2]", &Pair{ nil,&RegularNumber{1}, &RegularNumber{2}}},
		{"[[1,2],3]", add(&Pair{ nil, &RegularNumber{1}, &RegularNumber{2}}, &RegularNumber{3})},
		{"[9,[8,7]]", add(&RegularNumber{9}, &Pair{ nil, &RegularNumber{8}, &RegularNumber{7}}) },
		{"[[1,9],[8,5]]", add(&Pair{nil, &RegularNumber{1}, &RegularNumber{9}}, &Pair{nil, &RegularNumber{8}, &RegularNumber{5}}) },
		{
			example: "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
			expectedPair: add(
				add(
					add(
						&Pair{nil, &RegularNumber{1}, &RegularNumber{2}},
						&Pair{nil, &RegularNumber{3}, &RegularNumber{4}}),
					add(
						&Pair{nil, &RegularNumber{5}, &RegularNumber{6}},
						&Pair{nil, &RegularNumber{7}, &RegularNumber{8}}),
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
