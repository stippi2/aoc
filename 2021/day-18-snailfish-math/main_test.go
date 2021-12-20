package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func numberPair(left, right int) Node {
	return &Pair{ nil,&RegularNumber{left}, &RegularNumber{right}}
}

func Test_parseSnailfishNumber(t *testing.T) {
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

func Test_split(t *testing.T) {
	type test struct {
		example  string
		expected string
	}

	var explodeExamples = []test{
		{"[9,4]", "[9,4]"},
		{"[11,4]", "[[5,6],4]"},
		{"[4,12]", "[4,[6,6]]"},
	}

	for _, e := range explodeExamples {
		p := parseSnailfishNumber(e.example).(*Pair)
		p.Split()
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
		{"[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]", 3993},
	}

	for _, e := range explodeExamples {
		p := parseSnailfishNumber(e.example).(*Pair)
		assert.Equal(t, e.expected, p.Magnitude())
	}
}

func Test_reduce(t *testing.T) {
	type test struct {
		example  string
		expected string
	}

	var examples = []test{
		{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}

	for _, e := range examples {
		p := parseSnailfishNumber(e.example)
		assert.Equal(t, e.expected, fmt.Sprintf("%s", reduce(p)))
	}
}

func Test_reduceOnce(t *testing.T) {
	type test struct {
		example  string
		expected []string
	}

	var examples = []test{
		{
			example: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			expected: []string{
				"[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
				"[[[[0,7],4],[15,[0,13]]],[1,1]]",
				"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
				"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
				"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			},
		},
		{
			example: "[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[[[[0,7],4],[[7,8],[6,0]]],[8,1]]]",
			expected: []string{
				"[[[[4,0],[5,0]],[[[4,5],[2,6]],[9,5]]],[[[[0,7],4],[[7,8],[6,0]]],[8,1]]]",
				"[[[[4,0],[5,4]],[[0,[7,6]],[9,5]]],[[[[0,7],4],[[7,8],[6,0]]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,0],[15,5]]],[[[[0,7],4],[[7,8],[6,0]]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,0],[15,5]]],[[[0,11],[[7,8],[6,0]]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,0],[15,5]]],[[[0,18],[0,[14,0]]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,0],[15,5]]],[[[0,18],[14,0]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,0],[[7,8],5]]],[[[0,18],[14,0]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[0,13]]],[[[0,18],[14,0]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[0,[6,7]]]],[[[0,18],[14,0]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[[7,18],[14,0]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[[7,[9,9]],[14,0]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[[16,0],[23,0]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[[[8,8],0],[23,0]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,8]]],[[[0,8],[23,0]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,8]]],[[[0,8],[[11,12],0]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,8]]],[[[0,19],[0,12]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,8]]],[[[0,[9,10]],[0,12]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,8]]],[[[9,0],[10,12]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,8]]],[[[9,0],[[5,5],12]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,8]]],[[[9,5],[0,17]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,8]]],[[[9,5],[0,[8,9]]],[8,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,8]]],[[[9,5],[8,0]],[17,1]]]",
				"[[[[4,0],[5,4]],[[7,7],[6,8]]],[[[9,5],[8,0]],[[8,9],1]]]",
			},
		},
	}

	for _, e := range examples {
		p := parseSnailfishNumber(e.example)
		for _, expected := range e.expected {
			p, _ = reduceOnce(p)
			if !assert.Equal(t, expected, fmt.Sprintf("%s", p)) {
				break
			}
		}
	}
}

func Test_addingTwo(t *testing.T) {
	type test struct {
		left  string
		right  string
		expected string
	}

	var explodeExamples = []test{
		{
			left:     "[[[[4,3],4],4],[7,[[8,4],9]]]",
			right:    "[1,1]",
			expected: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			left:     "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			right:    "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			expected: "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
		{
			left:     "[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
			right:    "[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
			expected: "[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]",
		},
		{
			left:     "[[[7,9],[7,8]],[[8,3],8]]",
			right:    "[[[7,7],[7,7]],[[6,6],[6,6]]]",
			expected: "[[[[7,9],[7,8]],[[8,3],8]],[[[7,7],[7,7]],[[6,6],[6,6]]]]",
		},
	}

	for _, e := range explodeExamples {
		left := parseSnailfishNumber(e.left)
		right := parseSnailfishNumber(e.right)
		sum := add(left, right)
		fmt.Printf("before reduce: %s\n", sum)
		assert.Equal(t, e.expected, fmt.Sprintf("%s", reduce(sum)))
	}
}

func Test_adding(t *testing.T) {
	input := `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`
	numbers := parseInput(input)
	number := numbers[0]
	for i := 1; i < len(numbers); i++ {
		number = add(number, numbers[i])
		number = reduce(number)
	}
	assert.Equal(t, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", fmt.Sprintf("%s", number))
}

func Test_maxMagnitudeOfAnyTwo(t *testing.T) {
	input := `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`
	numbers := parseInput(input)
	maxMagnitude := maxMagnitudeOfAnyTwo(numbers)
	assert.Equal(t, 3993, maxMagnitude)
}
