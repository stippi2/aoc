package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	name string
	next []Node
}

func main() {
}

func parseInput(input string) *Node {
	lines := strings.Split(input, "\n")
	return nil
}

func loadInput(filename string) string {
	fileContents, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(fileContents))
}
