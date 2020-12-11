package main

import (
	"image"
	"image/color"

	"github.com/willie/advent/aoc"
)

const (
	empty    = "L"
	occupied = "#"
	floor    = "."
)

var slopes = []image.Point{
	{-1, 1}, {0, 1}, {1, 1},
	{-1, 0} /*{0,0}*/, {1, 0},
	{-1, -1}, {0, -1}, {1, -1},
}

func countAdjacent(in aoc.Grid, x, y int, s string) (count int) {
	for _, d := range slopes {
		if in.Get(x+d.X, y+d.Y, "") == s {
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
			if countAdjacent(in, x, y, occupied) == 0 {
				s = occupied
			}
		case occupied:
			if countAdjacent(in, x, y, occupied) >= 4 {
				s = empty
			}
		}

		dst.Set(x, y, s)
		return true
	})

	return
}

func part1(in aoc.Grid) (first int) {
	grids := []aoc.Grid{in}

	prev := in
	for {
		next := nextRound(prev)
		if next.Count(occupied) == prev.Count(occupied) {
			break
		}

		grids = append(grids, next)
		prev = next
	}
	first = prev.Count(occupied)

	images := []*image.Paletted{}
	for i, g := range grids {
		if i%2 != 0 {
			continue
		}

		img := g.NewRGBAImage(5)
		g.DrawImage(img, 5, map[string]color.Color{
			occupied: blue,
			empty:    black,
			floor:    red,
		})
		images = append(images, aoc.PaletteImageFromImage(img))
	}
	aoc.SaveGIFs("part1.gif", images, 1)

	return
}

func countVisibleOccupied(in aoc.Grid, x, y int) (count int) {
	for _, d := range slopes {
		in.SlopeIterate(x, y, d.X, d.Y, func(gx, gy int, seat string) bool {
			if seat == occupied {
				count++
				return false
			} else if seat == empty {
				return false
			}

			return true
		})
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
		next := aoc.NewBlankGrid(in.Width(), in.Height(), "")
		prev.Iterate(func(x, y int, s string) bool {
			switch s {
			case empty:
				if countVisibleOccupied(prev, x, y) == 0 {
					s = occupied
				}
			case occupied:
				if countVisibleOccupied(prev, x, y) >= 5 {
					s = empty
				}
			}

			next.Set(x, y, s)
			return true
		})

		if next.Count(occupied) == prev.Count(occupied) {
			break
		}

		grids = append(grids, next)
		prev = next
	}
	second = prev.Count(occupied)

	images := []*image.Paletted{}
	for i, g := range grids {
		if i%2 != 0 {
			continue
		}

		img := g.NewRGBAImage(5)
		g.DrawImage(img, 5, map[string]color.Color{
			occupied: blue,
			empty:    black,
			floor:    red,
		})
		images = append(images, aoc.PaletteImageFromImage(img))
	}
	aoc.SaveGIFs("part2.gif", images, 1)

	return
}

var (
	black = color.RGBA{0, 0, 0, 255}
	red   = color.RGBA{255, 0, 0, 255}
	gray  = color.RGBA{128, 128, 128, 255}
	green = color.RGBA{0, 255, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	white = color.RGBA{255, 255, 255, 255}
)

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
