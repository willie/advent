package main

import (
	"fmt"

	"github.com/willie/advent/aoc"
)

const (
	empty    = "L"
	occupied = "#"
	floor    = "."
)

func countAdjacentOccupied(in aoc.Grid, x, y int) (count int) {
	delta := []struct {
		x, y int
	}{
		{-1, 1}, {0, 1}, {1, 1},
		{-1, 0} /*{0,0}*/, {1, 0},
		{-1, -1}, {0, -1}, {1, -1},
	}

	for _, d := range delta {
		if in.Get(x+d.x, y+d.y, "") == occupied {
			count++
		}
	}

	return
}

func nextRound(in aoc.Grid) (dst aoc.Grid) {
	dst = aoc.NewBlankGrid(in.Width(), in.Height(), "")

	in.Iterate(func(x, y int, s string) bool {
		switch s {
		case empty:
			if countAdjacentOccupied(in, x, y) == 0 {
				s = occupied
			}
		case occupied:
			if countAdjacentOccupied(in, x, y) >= 4 {
				s = empty
			}
		default:
		}

		dst.Set(x, y, s)
		return true
	})

	return
}

func count(in aoc.Grid, t string) (total int) {
	in.Iterate(func(x, y int, s string) bool {
		if s == t {
			total++
		}
		return true
	})
	return
}

func part1(in aoc.Grid) (first int) {
	grids := []aoc.Grid{in}

	prev := in
	for {
		next := nextRound(prev)
		if count(next, occupied) == count(prev, occupied) {
			break
		}

		grids = append(grids, next)
		prev = next
	}
	first = count(prev, occupied)

	for i, g := range grids {
		fmt.Println("-------", "grid", i, count(g, occupied))
		// g.Print()
	}

	return
}

func countVisibleOccupied(in aoc.Grid, x, y int) (count int) {
	delta := []struct {
		x, y int
	}{
		{-1, 1}, {0, 1}, {1, 1},
		{-1, 0} /*{0,0}*/, {1, 0},
		{-1, -1}, {0, -1}, {1, -1},
	}

	max := aoc.Max(in.Width(), in.Height())

	for _, d := range delta {
		for scale := 1; scale <= max; scale++ {
			s := in.Get(x+(d.x*scale), y+(d.y*scale), "")
			if s == occupied {
				count++
				break
			} else if s == empty {
				break
			}
		}
	}

	return
}

func nextRound2(in aoc.Grid) (dst aoc.Grid) {
	dst = aoc.NewBlankGrid(in.Width(), in.Height(), "")

	in.Iterate(func(x, y int, s string) bool {
		switch s {
		case empty:
			if countVisibleOccupied(in, x, y) == 0 {
				s = occupied
			}
		case occupied:
			if countVisibleOccupied(in, x, y) >= 5 {
				s = empty
			}
		}

		dst.Set(x, y, s)
		return true
	})

	return
}

func part2(in aoc.Grid) (second int) {
	grids := []aoc.Grid{in}

	prev := in
	for {
		next := nextRound2(prev)
		if count(next, occupied) == count(prev, occupied) {
			break
		}

		grids = append(grids, next)
		prev = next
	}
	second = count(prev, occupied)

	for i, g := range grids {
		fmt.Println("-------", "grid", i, count(g, occupied))
		g.Print()
	}

	return
}

const day = "https://adventofcode.com/2020/day/11"

func main() {
	println(day)
	aoc.Input(day)

	println("------- part1")

	aoc.Test("test", part1(aoc.LoadGrid("test")), 37)
	aoc.Run("run", part1(aoc.LoadGrid(day)))

	println("------- part2")

	g := aoc.LoadGrid("count")
	x, y := g.FindFirst(empty)
	aoc.Test("count", countVisibleOccupied(g, x, y), 8)

	g = aoc.LoadGrid("count2")
	x, y = g.FindFirst(empty)
	aoc.Test("count2", countVisibleOccupied(g, x, y), 0)

	g = aoc.LoadGrid("count3")
	x, y = g.FindFirst(empty)
	aoc.Test("count3", countVisibleOccupied(g, x, y), 0)

	aoc.Test("test", part2(aoc.LoadGrid("test")), 26)

	aoc.Run("run", part2(aoc.LoadGrid(day)))
}
