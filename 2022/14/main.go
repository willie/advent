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
	grid[image.Pt(500, 0)] = "+"

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

	grid.Print(".")
	fmt.Println(sand, grid.Bounds())
}

func main() {
	part1("test.txt")
	// part1("input.txt")
}
