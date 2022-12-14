package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(name string) {
	sand := 0

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

	for simulate := true; simulate; sand++ {
		grain := image.Pt(500, 0)
		for dropping := true; dropping && simulate; {
			paths := grid.Exists(aoc.Map(grain.Add, []image.Point{{0, 1}, {-1, 1}, {1, 1}}))

			if len(paths) == 0 {
				if grain == image.Pt(500, 0) || grain.Y > abyss.Y {
					simulate, dropping = false, false
					continue
				}

				grid[grain] = "o"
				dropping = false
				break
			}

			grain = paths[0]
		}
		// grid.PrintYFlipped(".")
	}

	grid.PrintYFlipped(" ")
	fmt.Println(sand, grid.Bounds())
}

func main() {
	part1("test.txt")
	// part1("input.txt")
}
