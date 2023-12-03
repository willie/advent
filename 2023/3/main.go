package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/willie/advent/aoc"
)

// includes boundaries
func PtInRect(p image.Point, r image.Rectangle) bool {
	return r.Min.X <= p.X && p.X <= r.Max.X &&
		r.Min.Y <= p.Y && p.Y <= r.Max.Y
}

func part1(grid aoc.Grid2[string]) (total int) {
	bounds := grid.Bounds()
	for y := bounds.Max.Y; y >= bounds.Min.Y; y-- {
		var part string
		adjacent := false

		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			if value, ok := grid[image.Pt(x, y)]; ok && strings.ContainsAny(value, "0123456789") {
				part += value

				for _, adj := range grid.EightWayAdjacent(image.Pt(x, y)) {
					if !strings.ContainsAny(grid[adj], ".0123456789") {
						adjacent = true
						break
					}
				}

			} else {
				if len(part) > 0 {

					if adjacent {
						// fmt.Println(part)
						total += aoc.AtoI(part)
					}
					part = ""
					adjacent = false
				}
			}
		}

		if len(part) > 0 {

			if adjacent {
				// fmt.Println(part)

				total += aoc.AtoI(part)
			}
			part = ""
			adjacent = false
		}

	}

	return
}

func part2(grid aoc.Grid2[string]) (total int) {
	parts := make(map[image.Rectangle]int)

	bounds := grid.Bounds()
	for y := bounds.Max.Y; y >= bounds.Min.Y; y-- {
		var part string
		var pts []image.Point

		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			if value, ok := grid[image.Pt(x, y)]; ok && strings.ContainsAny(value, "0123456789") {
				part += value
				pts = append(pts, image.Pt(x, y))

			} else {
				if len(part) > 0 {
					fmt.Println(part)
					b := aoc.Bounds(pts)
					parts[b] = aoc.AtoI(part)
					part = ""
					pts = nil
				}
			}
		}

		if len(part) > 0 {
			fmt.Println(part)
			b := aoc.Bounds(pts)
			parts[b] = aoc.AtoI(part)
			part = ""
			pts = nil
		}

	}

	fmt.Println(parts)

	gears := grid.Contains("*")
	fmt.Println(gears)
	for _, gear := range gears {
		var intersect []int

		for bounds, part := range parts {
			// if slices.Contains[image.Point](grid.EightWayAdjacent(gear), bounds) {
			// 	intersect = append(intersect, part)
			// }

			for _, adj := range grid.EightWayAdjacent(gear) {
				fmt.Println(grid.Get(adj, ""))

				if PtInRect(adj, bounds) {
					intersect = append(intersect, part)
					break
				}
			}
		}

		fmt.Println(intersect)
		if len(intersect) == 2 {
			total += aoc.Product(intersect...)
		}

	}

	return
}

const day = "https://adventofcode.com/2023/day/3"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.LoadStringGrid(aoc.Strings("test"))), 4361)
	aoc.Test("test2", part2(aoc.LoadStringGrid(aoc.Strings("test2"))), 467835)

	println("-------")

	// aoc.Run("part1", part1(aoc.LoadStringGrid(aoc.Strings(day))))
	// aoc.Run("part2", part2(aoc.LoadStringGrid(aoc.Strings(day))))
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
