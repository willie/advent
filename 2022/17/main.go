package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/willie/advent/aoc"
)

var pieces = [][]int{
	{},
	{},
	{},
	{},
}

// ####

// .#.
// ###
// .#.

// ..#
// ..#
// ###

// #
// #
// #
// #

// ##
// ##

type loop[T any] struct {
	l     []T
	count int
}

func (l *loop) next() (n T) {
	n = l.l[count%len(l.l)]
	count++
	return
}

func loopGet[T any](in []T, i int) T { return in[i%len(in)] }

func part1(name string, rocks int) {
	grid := aoc.Grid2[string]{}
	for x := 0; x < 7; x++ {
		grid[image.Pt(x, -1)] = "-"
	}

	movement := &loop{l: strings.Split(aoc.String(name), "")}

	floor, height := 0, 3
	for i := 0; i < rocks; i++ {
		start := image.Pt(2, height)

	}

	fmt.Println()
}

func main() {
	r := []int{0, 1, 2, 3, 4}
	for i := 0; i < 20; i++ {
		fmt.Println(get(r, i))
	}

	part1("test.txt", 10)
	// part1("input.txt", 2000000)

	// part1draw("test.txt")
	// part1geometry("test.txt", 10)
	// part1geometry("input.txt", 2000000)

	// println("------")

	// part2("test.txt")
	// part2("input.txt")

}
