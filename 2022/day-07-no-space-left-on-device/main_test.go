package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func Test_Part1(t *testing.T) {
	root := parseInput(exampleInput)
	assert.Equal(t, 94853, root.getDir("a").size())
	assert.Equal(t, 24933642, root.getDir("d").size())
	assert.Equal(t, 95437, sumDirectoryBelow(root, 100000))
	assert.Equal(t, 48381165, root.size())
}

func Test_Part2(t *testing.T) {
	root := parseInput(exampleInput)
	assert.Equal(t, 24933642, findSmallest(root, 8381165))
}
