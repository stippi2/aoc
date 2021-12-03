package main

import (
	"io/ioutil"
	"strings"
)


func main() {
}

func loadInput(filename string) []string {
	fileContents, err := ioutil.ReadFile(filename)
	exitIfError(err)

	return strings.Split(string(fileContents), "\n")
}

func exitIfError(err error) {
	if err != nil {
		panic(err)
	}
}
