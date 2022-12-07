package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/willie/advent/aoc"
)

type node struct {
	name  string
	files map[string]int64
	dirs  []*node
}

func NewNode(name string) *node {
	return &node{
		name:  name,
		files: map[string]int64{},
		dirs:  []*node{},
	}
}

func (n *node) AddNode(name string) (child *node) {
	child = NewNode(name)
	n.dirs = append(n.dirs, child)
	return
}

func (n *node) Size() (size int64) {
	for _, s := range n.files {
		size += s
	}

	for _, d := range n.dirs {
		size += d.Size()
	}

	return
}

func (n *node) Dirs() (nodes []*node) {
	nodes = append(nodes, n.dirs...)

	for _, c := range n.dirs {
		nodes = append(nodes, c.Dirs()...)
	}

	return
}

func part1(name string) {
	input := aoc.Strings(name)

	// dirs := map[string]*node{}
	root := NewNode("/")
	// dirs["/"] = root

	workingDir := aoc.Stack[*node]{}
	for i := 0; i < len(input); i++ {
		s := input[i]

		switch {
		case strings.Index(s, "$ cd /") == 0:
			workingDir.Push(root)

		case strings.Index(s, "$ cd ..") == 0:
			workingDir.Pop()

		case strings.Index(s, "$ cd ") == 0:
			dirName := s[len("$ cd "):]

			n := NewNode(dirName)
			// dirs[dirName] = n

			workingDir.Top().dirs = append(workingDir.Top().dirs, n)
			workingDir.Push(n)

		case strings.Index(s, "$ ls") == 0:
		case strings.Index(s, "dir") == 0:

		default:
			var size int
			var name string
			fmt.Sscanf(s, "%d %s", &size, &name)

			workingDir.Top().files[name] = int64(size)
		}
	}

	free := 70000000 - root.Size()

	p1, p2 := int64(0), int64(math.MaxInt64)
	for _, dir := range root.Dirs() {
		size := dir.Size()

		if size <= 100000 {
			p1 += size
		}

		if size < p2 && (size+free >= 30000000) {
			p2 = size
		}
	}

	println(p1, p2)
}

func main() {
	part1("test.txt")
	part1("input.txt")
}
