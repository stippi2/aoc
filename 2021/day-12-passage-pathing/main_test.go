package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseInput(t *testing.T) {
	n := parseInput(examples[0])
	assert.Equal(t, "start", n.name)
	if assert.Equal(t, 2, len(n.next)) {
		assert.Equal(t, "A", n.next[0].name)
		assert.Equal(t, "b", n.next[1].name)
	}
}

func Test_findEnd(t *testing.T) {
	assert.Equal(t, 10, findEnd(parseInput(examples[0]), ""))
	assert.Equal(t, 19, findEnd(parseInput(examples[1]), ""))
	assert.Equal(t, 226, findEnd(parseInput(examples[2]), ""))
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
