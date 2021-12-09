package main

import (
	"image"

	"github.com/willie/advent/aoc"
)

const (
	invalid = "9"
)

func part1(in []string) (result int) {
	// lowpoints := []int{}

	g := aoc.NewGrid(in)
	g.Iterate(func(x, y int, s string) (next bool) {
		next = true
		values := []int{}

		for _, diff := range adjacent {
			if i := g.Get(x+diff.X, y+diff.Y, invalid); i != invalid {
				values = append(values, aoc.AtoI(i))
			}
		}

		if len(values) == 0 {
			return
		}

		c := aoc.AtoI(s)
		if c < aoc.Min(values...) {
			result += (1 + c)
			// fmt.Println(x, y, " -> ", c)
		}

		return
	})

	return
}

var (
	adjacent = []image.Point{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
)

func part2(in []string) (result int) {
	// points := map[image.Point]int{}

	g := aoc.NewGrid(in)
	g.Iterate(func(x, y int, s string) (next bool) {
		next = true
		values := []int{}

		for _, diff := range adjacent {
			if i := g.Get(x+diff.X, y+diff.Y, invalid); i != invalid {
				values = append(values, aoc.AtoI(i))
			}
		}

		if len(values) == 0 {
			return
		}

		c := aoc.AtoI(s)
		if c < aoc.Min(values...) {
			result += (1 + c)
			// fmt.Println(x, y, " -> ", c)
		}

		return
	})

	return
}

const day = "https://adventofcode.com/2021/day/9"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 15)
	aoc.Test("test2", part2(aoc.Strings("test")), 15)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
