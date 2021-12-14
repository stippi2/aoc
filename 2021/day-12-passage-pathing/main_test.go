package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseInput(t *testing.T) {
	n := parseInput(examples[0]).start
	assert.Equal(t, "start", n.name)
	if assert.Equal(t, 2, len(n.next)) {
		assert.Equal(t, "A", n.next[0].name)
		assert.Equal(t, "b", n.next[1].name)
	}
}

func Test_findPathsPart1(t *testing.T) {
	assert.Len(t, findPathsPart1(parseInput(examples[0])), 10)
	assert.Len(t, findPathsPart1(parseInput(examples[1])), 19)
	assert.Len(t, findPathsPart1(parseInput(examples[2])), 226)
}

func Test_findPathsPart2(t *testing.T) {
	assert.Len(t, findPathsPart2(parseInput(examples[0])), 36)
	assert.Len(t, findPathsPart2(parseInput(examples[1])), 103)
	assert.Len(t, findPathsPart2(parseInput(examples[2])), 3509)
}

var examples = []string{
`start-A
start-b
A-c
A-b
b-d
A-end
b-end`,

`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`,

`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`,
}
