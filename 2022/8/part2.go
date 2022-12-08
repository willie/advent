package main

import "github.com/willie/advent/aoc"

func part2(name string) {
	g := aoc.LoadGrid(name)

	var score int
	g.Iterate(func(x, y int, s string) bool {
		ss := aoc.AtoI(s)

		left := 0
		g.SlopeIterate(x, y, -1, 0, func(gx, gy int, v string) bool {
			n := aoc.AtoI(v)

			left++
			if n >= ss {
				return false
			}

			return true
		})

		right := 0
		g.SlopeIterate(x, y, 1, 0, func(gx, gy int, v string) bool {
			n := aoc.AtoI(v)

			right++
			if n >= ss {
				return false
			}

			return true
		})

		down := 0
		g.SlopeIterate(x, y, 0, -1, func(gx, gy int, v string) bool {
			n := aoc.AtoI(v)

			down++
			if n >= ss {
				return false
			}

			return true
		})

		// up
		up := 0
		g.SlopeIterate(x, y, 0, 1, func(gx, gy int, v string) bool {
			n := aoc.AtoI(v)

			up++
			if n >= ss {
				return false
			}
			return true
		})

		scenic := left * right * up * down
		if scenic > score {
			score = scenic
		}

		return true
	})

	println(score)
}
