package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	size int
}

type Dir struct {
	parent *Dir
	dirs   map[string]*Dir
	files  map[string]*File
}

func (d *Dir) size() int {
	totalSize := 0
	for _, dir := range d.dirs {
		totalSize += dir.size()
	}
	for _, file := range d.files {
		totalSize += file.size
	}
	return totalSize
}

func (d *Dir) visit(visitor func(dir *Dir)) {
	for _, dir := range d.dirs {
		dir.visit(visitor)
	}
	visitor(d)
}

func (d *Dir) appendDir(name string, dir *Dir) {
	dir.parent = d
	d.dirs[name] = dir
}

func (d *Dir) appendFile(name string, file *File) {
	d.files[name] = file
}

func (d *Dir) getDir(name string) *Dir {
	return d.dirs[name]
}

func sumDirectoryBelow(root *Dir, maxSize int) int {
	relevantDirs := map[*Dir]int{}
	root.visit(func(dir *Dir) {
		size := dir.size()
		if size <= maxSize {
			relevantDirs[dir] = size
		}
	})
	totalSize := 0
	for _, size := range relevantDirs {
		totalSize += size
	}
	return totalSize
}

func findSmallest(root *Dir, minSize int) int {
	var smallestDir *Dir
	root.visit(func(dir *Dir) {
		size := dir.size()
		if size >= minSize {
			if smallestDir == nil || smallestDir.size() > size {
				smallestDir = dir
			}
		}
	})
	return smallestDir.size()
}

func main() {
	root := parseInput(loadInput("puzzle-input.txt"))
	fmt.Printf("total size of all directories <= 100000: %v\n", sumDirectoryBelow(root, 100000))
	free := 70000000 - root.size()
	required := 30000000 - free
	fmt.Printf("smallest directory to free up %v: %v\n", required, findSmallest(root, required))
}

func newDir() *Dir {
	return &Dir{
		files: map[string]*File{},
		dirs:  map[string]*Dir{},
	}
}

func parseInput(input string) (root *Dir) {
	lines := strings.Split(input, "\n")
	root = newDir()
	current := root
	for _, line := range lines {
		if line == "$ cd .." {
			current = current.parent
		} else if line == "$ cd /" {
			current = root
		} else if line == "$ ls" {
			// Can be ignored
		} else if strings.HasPrefix(line, "$ cd ") {
			name := strings.TrimPrefix(line, "$ cd ")
			current = current.getDir(name)
		} else {
			if strings.HasPrefix(line, "$ ") {
				panic("unhandled command")
			}
			parts := strings.Split(line, " ")
			if parts[0] == "dir" {
				current.appendDir(parts[1], newDir())
			} else {
				size, _ := strconv.Atoi(parts[0])
				current.appendFile(parts[1], &File{size})
			}
		}
	}
	return root
}

func loadInput(filename string) string {
	fileContents, _ := os.ReadFile(filename)
	return string(fileContents)
}
