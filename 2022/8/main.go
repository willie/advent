package main

import (
	"github.com/willie/advent/aoc"
)

func part1(name string) {
	g := aoc.LoadGrid(name)

	var count int
	g.Iterate(func(x, y int, s string) bool {
		if x == 0 || y == 0 || x == g.Width() || y == g.Height() {
			count++
			return true
		}

		var visible bool

		visible = true
		for i := 0; i < x; i++ {
			n := g.At(i, y)
			if aoc.AtoI(n) >= aoc.AtoI(s) {
				visible = false
			}
		}

		if visible {
			count++
			return true
		}

		visible = true
		for i := x + 1; i < g.Width(); i++ {
			n := g.At(i, y)
			if aoc.AtoI(n) >= aoc.AtoI(s) {
				visible = false
			}
		}

		if visible {
			count++
			return true
		}

		visible = true
		for i := 0; i < y; i++ {
			n := g.At(x, i)
			if aoc.AtoI(n) >= aoc.AtoI(s) {
				visible = false
			}
		}

		if visible {
			count++
			return true
		}

		visible = true
		for i := y + 1; i < g.Height(); i++ {
			n := g.At(x, i)
			if aoc.AtoI(n) >= aoc.AtoI(s) {
				visible = false
			}
		}

		if visible {
			count++
			return true
		}

		return true
	})

	println(count)
}

func main() {
	part1("test.txt")
	part1("input.txt")

	println("------")

	part2("test.txt")
	part2("input.txt")
}
