package main

import (
	"fmt"
	"image"
	"image/color"
	"sort"

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

func part2(in []string) (result int) {
	g := aoc.NewGrid(in)

	// visualize
	saveImage("9.png", g)

	basins := []int{}

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
			// result += (1 + c)
			points := map[image.Point]int{}
			fill(g, x, y, points)

			basins = append(basins, len(points))
		}

		return
	})

	sort.Ints(basins)
	basins = basins[len(basins)-3:]

	return aoc.Product(basins...)
}

const day = "https://adventofcode.com/2021/day/9"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 15)
	aoc.Test("test2", part2(aoc.Strings("test")), 1134)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
