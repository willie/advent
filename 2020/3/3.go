package main

import (
	"fmt"

	"github.com/willie/advent/aoc"
)

func part1(in []string, right, down int) (trees int) {
	grid := aoc.NewGrid(in)

	// work down
	x := 0
	for y := 0; y < grid.Height(); y += down {
		nx := x % grid.Width()

		if grid.At(nx, y) == "#" {
			trees++
		}

		x += right
	}

	return
}

func part2(in []string) int {
	slopes := []struct {
		right, down int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	var trees []int
	for _, i := range slopes {
		trees = append(trees, part1(in, i.right, i.down))
	}

	fmt.Println(trees)

	return aoc.Product(trees...)
}

const day = "https://adventofcode.com/2020/day/3"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test"), 3, 1), 7)
	aoc.Run("part1", part1(aoc.Strings(day), 3, 1))

	println("-------")

	aoc.Test("test2", part2(aoc.Strings("test")), 336)
	aoc.Run("part2", part2(aoc.Strings(day)))
}
