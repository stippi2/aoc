package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

func Test_part1(t *testing.T) {
	computers := parseInput(example)
	assert.Equal(t, 7, countSetsOfThreeStartingWithT(computers))
}
