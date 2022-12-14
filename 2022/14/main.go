package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(name string) {
	grid := aoc.Grid2[string]{}
	grid[image.Pt(500, -1)] = "+"

	cursor := image.Point{}
	for _, s := range aoc.Strings(name) {
		for i, coord := range strings.Split(s, " -> ") {
			pt := image.Point{}
			fmt.Sscanf(coord, "%d,%d", &pt.X, &pt.Y)

			if i == 0 {
				cursor = pt
				continue
			}

			grid.IterateLine(cursor, pt, func(pt image.Point, v string) bool {
				grid[pt] = "#"
				return true
			})
			cursor = pt
		}
	}

	abyss := image.Pt(500, grid.Bounds().Max.Y+1)

	for simulate := true; simulate; {
		grain := image.Pt(500, 0)
		for dropping := true; dropping && simulate; {
			paths := aoc.Filter(func(in image.Point) bool {
				return grid.Get(in, " ") == " "
			}, aoc.Map(grain.Add, []image.Point{{0, 1}, {-1, 1}, {1, 1}}))

			if len(paths) == 0 {
				if grain == image.Pt(500, 0) {
					simulate = false
					continue
				}

				grid[grain] = "o"
				dropping = false

				continue
			}

			grain = paths[0]

			if grain.Y > abyss.Y {
				simulate, dropping = false, false
				continue
			}
		}
	}

	// grid.PrintYFlipped(" ")
	fmt.Println(len(grid.Contains("o")), grid.Bounds())
}

func part2(name string) {
	grid := aoc.Grid2[string]{}
	grid[image.Pt(500, -1)] = "+"

	cursor := image.Point{}
	for _, s := range aoc.Strings(name) {
		for i, coord := range strings.Split(s, " -> ") {
			pt := image.Point{}
			fmt.Sscanf(coord, "%d,%d", &pt.X, &pt.Y)

			if i == 0 {
				cursor = pt
				continue
			}

			grid.IterateLine(cursor, pt, func(pt image.Point, v string) bool {
				grid[pt] = "#"
				return true
			})
			cursor = pt
		}
	}

	abyss := image.Pt(500, grid.Bounds().Max.Y+1)

	for simulate := true; simulate; {
		grain := image.Pt(500, -1)
		for dropping := true; dropping && simulate; {
			paths := aoc.Filter(func(in image.Point) bool {
				return grid.Get(in, " ") == " "
			}, aoc.Map(grain.Add, []image.Point{{0, 1}, {-1, 1}, {1, 1}}))

			if len(paths) == 0 {
				if grain == image.Pt(500, 0) {
					grid[grain] = "o"
					simulate = false
					continue
				}

				grid[grain] = "o"
				dropping = false

				continue
			}

			grain = paths[0]

			if grain.Y == abyss.Y {
				grid[grain] = "o"
				dropping = false
				continue
			}
		}
	}

	// grid.PrintYFlipped(" ")
	fmt.Println(len(grid.Contains("o")), grid.Bounds())
}

func main() {
	part1("test.txt")
	part1("input.txt")

	println("------")

	part2("test.txt")
	part2("input.txt")

}
