package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node interface {
	AddLeft(number int, previous Node) bool
	AddRight(number int, previous Node) bool
	Magnitude() int
	Explode(level int) (Node, bool)
	Split() (Node, bool)
}

type RegularNumber struct {
	value int
}

func (n *RegularNumber) AddLeft(number int, _ Node) bool {
	n.value += number
	return true
}

func (n *RegularNumber) AddRight(number int, _ Node) bool {
	n.value += number
	return true
}

func (n *RegularNumber) Magnitude() int {
	return n.value
}

func (n *RegularNumber) Explode(_ int) (Node, bool) {
	return n, false
}

func (n *RegularNumber) Split() (Node, bool) {
	if n.value > 9 {
		return &Pair{
			parent: nil,
			left:   &RegularNumber{n.value / 2},
			right:  &RegularNumber{(n.value + 1) / 2},
		}, true
	}
	return n, false
}

func (n *RegularNumber) String() string {
	return fmt.Sprintf("%v", n.value)
}

type Pair struct {
	parent Node
	left   Node
	right  Node
}

func (p *Pair) AddLeft(number int, previous Node) bool {
	if previous == p.right {
		return p.left.AddLeft(number, p)
	}
	if previous == p.parent {
		return p.right.AddLeft(number, p)
	}
	if p.parent == nil {
		return false
	}
	return p.parent.AddLeft(number, p)
}

func (p *Pair) AddRight(number int, previous Node) bool {
	if previous == p.left {
		return p.right.AddRight(number, p)
	}
	if previous == p.parent {
		return p.left.AddRight(number, p)
	}
	if p.parent == nil {
		return false
	}
	return p.parent.AddRight(number, p)
}

func (p *Pair) Magnitude() int {
	return p.left.Magnitude() * 3 + p.right.Magnitude() * 2
}

func (p *Pair) Explode(level int) (Node, bool) {
	if level == 4 {
		return &RegularNumber{0}, true
	}
	newLeft, explodedLeft := p.left.Explode(level + 1)
	if explodedLeft {
		if p.left != newLeft {
			pair := p.left.(*Pair)
			left := pair.left.Magnitude()
			p.parent.AddLeft(left, p)
			right := pair.right.Magnitude()
			p.right.AddRight(right, p)
			p.left = newLeft
		}
		return p, true
	}
	newRight, explodedRight := p.right.Explode(level + 1)
	if explodedRight {
		if p.right != newRight {
			pair := p.right.(*Pair)
			left := pair.left.Magnitude()
			p.left.AddLeft(left, p)
			right := pair.right.Magnitude()
			p.parent.AddRight(right, p)
			p.right = newRight
		}
		return p, true
	}
	return p, false
}

func (p *Pair) Split() (Node, bool) {
	if newLeft, split := p.left.Split(); split {
		newLeft.(*Pair).parent = p
		p.left = newLeft
		return p, true
	}
	if newRight, split := p.right.Split(); split {
		newRight.(*Pair).parent = p
		p.right = newRight
		return p, true
	}
	return p, false
}

func (p *Pair) String() string {
	return fmt.Sprintf("[%v,%v]", p.left, p.right)
}

func main() {
	numbers := parseInput(loadInput("puzzle-input.txt"))
	// Part 1
	sum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		sum = reduce(add(sum, numbers[i]))
	}
	fmt.Printf("sum: %s, mangitude: %v\n", sum, sum.Magnitude())
	// Part 2
	// Need to parse again, since the nodes were modified by summing
	numbers = parseInput(loadInput("puzzle-input.txt"))
	maxMagnitude := maxMagnitudeOfAnyTwo(numbers)
	fmt.Printf("largest magnitude of any two numbers: %v\n", maxMagnitude)
}

func maxMagnitudeOfAnyTwo(numbers []Node) int {
	maxMagnitude := 0
	for i, numberA := range numbers {
		for j, numberB := range numbers {
			if i == j {
				continue
			}
			number := reduce(add(numberA, numberB))
			magnitude := number.Magnitude()
			fmt.Printf("%s + %s = %s (magnitude: %v)\n", numberA, numberB, number, magnitude)
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}
		}
	}
	return maxMagnitude
}

func reduceOnce(node Node) bool {
	_, anyExploded := node.Explode(0)
	if anyExploded {
		return true
	}
	_, anySplit := node.Split()
	if anySplit {
		return true
	}
	return false
}

func reduce(node Node) Node {
	clone := parseSnailfishNumber(fmt.Sprintf("%s", node))
	for reduceOnce(clone) {}
	return clone
}

func add(left, right Node) Node {
	result := &Pair{}
	if leftPair, ok := left.(*Pair); ok {
		leftPair.parent = result
	}
	if rightPair, ok := right.(*Pair); ok {
		rightPair.parent = result
	}
	result.left = left
	result.right = right
	return result
}

func parseSnailfishNumber(line string) Node {
	// trim the outer brackets (if any)
	line = strings.TrimPrefix(line, "[")
	line = strings.TrimSuffix(line, "]")

	level := 0
	for i := 0; i < len(line); i++ {
		char := line[i:i+1]
		if char == "[" {
			level++
		} else if char == "]" {
			level--
		} else if char == "," {
			if level == 0 {
				// This is a pair, and we found the "," on the pair's level
				return add(parseSnailfishNumber(line[:i]), parseSnailfishNumber(line[i+1:]))
			}
		}
	}
	number, err := strconv.Atoi(line)
	if err != nil {
		panic(fmt.Sprintf("failed to parse '%s': %s", line, err))
	}
	return &RegularNumber{number}
}

func parseInput(input string) (numbers []Node) {
	for _, line := range strings.Split(input, "\n") {
		numbers = append(numbers, parseSnailfishNumber(line))
	}
	return
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
