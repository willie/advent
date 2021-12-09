package main

import (
	"fmt"

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

		if i := g.Get(x-1, y, invalid); i != invalid {
			values = append(values, aoc.AtoI(i))
		}

		if i := g.Get(x+1, y, invalid); i != invalid {
			values = append(values, aoc.AtoI(i))
		}

		if i := g.Get(x, y+1, invalid); i != invalid {
			values = append(values, aoc.AtoI(i))
		}

		if i := g.Get(x, y-1, invalid); i != invalid {
			values = append(values, aoc.AtoI(i))
		}

		if len(values) == 0 {
			return
		}

		c := aoc.AtoI(s)
		if c < aoc.Min(values...) {
			result += (1 + c)

			fmt.Println(x, y, " -> ", c)
		}

		return
	})

	return
}

func part2(in []string) (result int) {
	// lowpoints := []int{}

	g := aoc.NewGrid(in)
	g.Iterate(func(x, y int, s string) (next bool) {
		next = true
		values := []int{}

		if i := g.Get(x-1, y, invalid); i != invalid {
			values = append(values, aoc.AtoI(i))
		}
		if i := g.Get(x+1, y, invalid); i != invalid {
			values = append(values, aoc.AtoI(i))
		}
		if i := g.Get(x, y+1, invalid); i != invalid {
			values = append(values, aoc.AtoI(i))
		}
		if i := g.Get(x, y-1, invalid); i != invalid {
			values = append(values, aoc.AtoI(i))
		}
		if len(values) == 0 {
			return
		}

		c := aoc.AtoI(s)
		if c < aoc.Min(values...) {
			result += (1 + c)

			fmt.Println(x, y, " -> ", c)
		}

		return
	})

	return
}

const day = "https://adventofcode.com/2021/day/9"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 15)
	// aoc.Test("test2", part2(aoc.Strings("test")), 12)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
