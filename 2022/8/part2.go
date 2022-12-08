package main

import "github.com/willie/advent/aoc"

func part2(name string) {
	g := aoc.LoadGrid(name)

	var score int
	g.Iterate(func(x, y int, s string) bool {
		ss := aoc.AtoI(s)

		var left, right, up, down int

		g.SlopeIterate(x, y, -1, 0, func(gx, gy int, v string) bool {
			left++
			return aoc.AtoI(v) < ss
		})

		g.SlopeIterate(x, y, 1, 0, func(gx, gy int, v string) bool {
			right++
			return aoc.AtoI(v) < ss
		})

		g.SlopeIterate(x, y, 0, -1, func(gx, gy int, v string) bool {
			down++
			return aoc.AtoI(v) < ss
		})

		g.SlopeIterate(x, y, 0, 1, func(gx, gy int, v string) bool {
			up++
			return aoc.AtoI(v) < ss
		})

		scenic := left * right * up * down
		if scenic > score {
			score = scenic
		}

		return true
	})

	println(score)
}
