package main

import (
	"image"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(grid aoc.Grid2[string]) (total int) {
	bounds := grid.Bounds()
	for y := bounds.Max.Y; y >= bounds.Min.Y; y-- {
		var part string
		adjacent := false

		checkPart := func() {
			if len(part) > 0 {
				if adjacent {
					total += aoc.AtoI(part)
				}
				part, adjacent = "", false
			}
		}

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
				checkPart()
			}
		}

		checkPart()
	}

	return
}

type Part struct {
	value int
	pts   []image.Point
}

func part2(grid aoc.Grid2[string]) (total int) {
	parts := []Part{}

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
					// fmt.Println(part)

					parts = append(parts, Part{aoc.AtoI(part), pts})
					part = ""
					pts = nil
				}
			}
		}

		if len(part) > 0 {
			// fmt.Println(part)

			parts = append(parts, Part{aoc.AtoI(part), pts})
			part = ""
			pts = nil
		}

	}

	// fmt.Println(parts)

	gears := grid.Contains("*")
	// fmt.Println(gears)
	for _, gear := range gears {
		adjacent := aoc.NewSet(grid.EightWayAdjacent(gear)...)

		var intersect []int
		for _, part := range parts {
			if adjacent.ContainsAny(part.pts) {
				intersect = append(intersect, part.value)
			}
		}

		// fmt.Println(intersect)
		if len(intersect) == 2 {
			total += aoc.Product(intersect...)
		}

	}

	return
}

const day = "https://adventofcode.com/2023/day/4"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.LoadStringGrid(aoc.Strings("test"))), 4361)
	aoc.Test("test2", part2(aoc.LoadStringGrid(aoc.Strings("test2"))), 467835)

	println("-------")

	aoc.Run("part1", part1(aoc.LoadStringGrid(aoc.Strings(day))))
	aoc.Run("part2", part2(aoc.LoadStringGrid(aoc.Strings(day))))
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
