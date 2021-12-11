package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/willie/advent/aoc"
)

const (
	invalid = "-1"
)

func flash(g aoc.Grid, x, y int, points map[image.Point]int) (result int) {
	if _, ok := points[image.Point{x, y}]; ok {
		return
	}

	s := g.Get(x, y, invalid)
	if s == invalid {
		return
	}

	v := aoc.AtoI(s)
	if v > 9 {
		g.Set(x, y, "0")
		points[image.Point{x, y}] = 1
		result++

		for _, diff := range adjacent {
			if i := g.Get(x+diff.X, y+diff.Y, invalid); i != invalid {
				iv := aoc.AtoI(i)
				if iv == 0 {
					continue
				}

				g.Set(x+diff.X, y+diff.Y, fmt.Sprintf("%d", iv+1))
				result += flash(g, x+diff.X, y+diff.Y, points)
			}
		}
	}

	return
}

func part1(in []string, iterations int) (result int) {
	g := aoc.NewGrid(in)
	// g.Print()

	for i := 0; i < iterations; i++ {
		// increment
		g.Iterate(func(x, y int, s string) (next bool) {
			v := aoc.AtoI(s) + 1
			g.Set(x, y, fmt.Sprintf("%d", v))

			return true
		})

		// flashes
		points := map[image.Point]int{}

		g.Iterate(func(x, y int, s string) (next bool) {
			result += flash(g, x, y, points)
			return true
		})

		// println()
		// g.Print()
	}

	return
}

var (
	adjacent = []image.Point{{-1, 0}, {1, 0}, {0, 1}, {0, -1},
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
)

func fill(g aoc.Grid, x, y int, points map[image.Point]int) {
	if _, ok := points[image.Point{x, y}]; ok {
		return
	}

	if i := g.Get(x, y, invalid); i == invalid {
		return
	}

	points[image.Point{x, y}] = 1

	for _, diff := range adjacent {
		if i := g.Get(x+diff.X, y+diff.Y, invalid); i != invalid {
			fill(g, x+diff.X, y+diff.Y, points)
		}
	}
}

func saveImage(filename string, g aoc.Grid) {
	colormap := map[string]color.Color{"9": color.RGBA{0, 0, 0, 255}}
	for i := 0; i < 9; i++ {
		colormap[fmt.Sprintf("%d", i)] = color.RGBA{0, 0, uint8(255 - (i * (150 / 8))), 255}
	}

	img := g.NewRGBAImage(2)
	g.DrawImage(img, 2, colormap)

	aoc.SavePNG(filename, img)

}

func part2(in []string) (iterations int) {
	g := aoc.NewGrid(in)
	// g.Print()

	for {
		iterations++

		// increment
		g.Iterate(func(x, y int, s string) (next bool) {
			v := aoc.AtoI(s) + 1
			g.Set(x, y, fmt.Sprintf("%d", v))

			return true
		})

		// flashes
		points := map[image.Point]int{}

		g.Iterate(func(x, y int, s string) (next bool) {
			flash(g, x, y, points)
			return true
		})

		flashed := 0
		g.Iterate(func(x, y int, s string) (next bool) {
			if s == "0" {
				flashed++
			}
			return true
		})

		if flashed == 100 {
			break
		}

		// println()
		// g.Print()
	}

	return
}

const day = "https://adventofcode.com/2021/day/11"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test"), 100), 1656)
	aoc.Test("test2", part2(aoc.Strings("test")), 195)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day), 100))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
