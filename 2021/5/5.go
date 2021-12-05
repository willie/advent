package main

import (
	"fmt"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (result int) {
	var maxX, maxY int

	for _, i := range in {
		var x1, y1, x2, y2 int
		fmt.Sscanf(i, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		// fmt.Println(x1, y1, x2, y2)

		if x1 > maxX {
			maxX = x1
		}
		if x2 > maxX {
			maxX = x2
		}
		if y1 > maxY {
			maxY = x1
		}
		if y2 > maxY {
			maxY = y2
		}
	}

	fmt.Println(maxX, maxY)
	g := aoc.NewBlankGrid(maxX+1, maxY+1, ".")

	for _, i := range in {
		var x1, y1, x2, y2 int
		fmt.Sscanf(i, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		if (x1 == x2) || (y1 == y2) {
			g.IterateLine(x1, y1, x2, y2, func(x, y int, s string) bool {
				prev := 0
				if s != "." {
					prev = aoc.AtoI(s)
				}
				prev++
				g.Set(x, y, fmt.Sprintf("%d", prev))

				return true
			})
		}

	}

	// g.Print()

	count := 0
	g.Iterate(func(x, y int, s string) bool {
		v := 0
		if s != "." {
			v = aoc.AtoI(s)
		}

		if v > 1 {
			count++
		}
		return true
	})

	return count
}

func part2(in []string) (result int) {
	var maxX, maxY int

	for _, i := range in {
		var x1, y1, x2, y2 int
		fmt.Sscanf(i, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		// fmt.Println(x1, y1, x2, y2)

		if x1 > maxX {
			maxX = x1
		}
		if x2 > maxX {
			maxX = x2
		}
		if y1 > maxY {
			maxY = x1
		}
		if y2 > maxY {
			maxY = y2
		}
	}

	fmt.Println(maxX, maxY)
	g := aoc.NewBlankGrid(maxX+1, maxY+1, ".")

	for _, i := range in {
		var x1, y1, x2, y2 int
		fmt.Sscanf(i, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		g.IterateLine(x1, y1, x2, y2, func(x, y int, s string) bool {
			prev := 0
			if s != "." {
				prev = aoc.AtoI(s)
			}
			prev++
			g.Set(x, y, fmt.Sprintf("%d", prev))

			return true
		})

	}

	// g.Print()

	count := 0
	g.Iterate(func(x, y int, s string) bool {
		v := 0
		if s != "." {
			v = aoc.AtoI(s)
		}

		if v > 1 {
			count++
		}
		return true
	})

	return count
}

const day = "https://adventofcode.com/2021/day/5"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 5)
	aoc.Test("test2", part2(aoc.Strings("test")), 12)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
