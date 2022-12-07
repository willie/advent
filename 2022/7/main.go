package main

import (
	"fmt"
	"math"
	"path"
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

func useTree(name string) {
	workingDir := aoc.Stack[*node]{}

	root := NewNode("/")
	workingDir.Push(root)

	for _, s := range aoc.Strings(name) {

		switch {
		case strings.Index(s, "$ cd /") == 0:

		case strings.Index(s, "$ cd ..") == 0:
			workingDir.Pop()

		case strings.Index(s, "$ cd ") == 0:
			dirName := s[len("$ cd "):]
			n := workingDir.Top().AddNode(dirName)
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

func usePath(name string) {
	workingDir := ""
	dirSizes := map[string]int{}

	for _, s := range aoc.Strings(name) {
		switch {
		case strings.Index(s, "$ cd ") == 0:
			dirName := s[len("$ cd "):]
			workingDir = path.Join(workingDir, dirName)

		default:
			var size int
			if parsed, _ := fmt.Sscanf(s, "%d", &size); parsed == 0 {
				continue
			}

			for dir := workingDir; dir != "/"; dir = path.Dir(dir) {
				dirSizes[dir] += size
			}

			dirSizes["/"] += size
		}
	}

	free := 70000000 - dirSizes["/"]

	p1, p2 := 0, dirSizes["/"]
	for _, size := range dirSizes {
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
	useTree("test.txt")
	useTree("input.txt")

	fmt.Println("------")

	usePath("test.txt")
	usePath("input.txt")
}
