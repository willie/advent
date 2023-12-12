package main

import (
	"fmt"
	"image"
	"slices"

	"github.com/willie/advent/aoc"
	"golang.org/x/exp/maps"
)

func part1(in []string, expansion int) (total int) {
	// load it
	g := aoc.LoadStringGrid(in)
	g.PrintYFlipped(" ")
	fmt.Println()
	fmt.Println("number of galaxies", len(g))

	// keep only galaxies
	g = aoc.FilterMap(g, func(pt image.Point, v string) bool { return v == "#" })
	g.PrintYFlipped(" ")
	fmt.Println()
	fmt.Println("number of galaxies", len(g))

	// number the galaxies
	counter := 0
	bounds := g.Bounds()

	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			if _, ok := g[image.Pt(x, y)]; ok {
				counter++
				g[image.Pt(x, y)] = fmt.Sprintf("%d", counter)
			}
		}
	}
	g.PrintYFlipped(" ")

	// find all the empty rows and columns
	emptyRows := []int{}
	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		rowEmpty := true
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			if _, ok := g[image.Pt(x, y)]; ok {
				rowEmpty = false
				break
			}
		}
		if rowEmpty {
			emptyRows = append(emptyRows, y)
		}
	}

	emptyColumns := []int{}
	for x := bounds.Min.X; x <= bounds.Max.X; x++ {
		colEmpty := true
		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
			if _, ok := g[image.Pt(x, y)]; ok {
				colEmpty = false
				break
			}
		}
		if colEmpty {
			emptyColumns = append(emptyColumns, x)
		}
	}

	fmt.Println("number of galaxies", len(g))

	// iterate over the galaxies, find the manhattan distance to all the others
	for _, path := range uniquePointPairs(maps.Keys(g)) {
		src, dest := path.start, path.end

		diff := 0

		for _, row := range emptyRows {
			if NumIsBetween(row, src.Y, dest.Y) {
				diff += expansion
			}
		}
		for _, col := range emptyColumns {
			if NumIsBetween(col, src.X, dest.X) {
				diff += expansion
			}
		}

		dist := aoc.ManhattanDistance(src.X, src.Y, dest.X, dest.Y)
		total += (dist + diff)
	}

	return
}

const day = "https://adventofcode.com/2023/day/11"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test"), 1), 374)
	aoc.Test("test1", part1(aoc.Strings("test"), 100-1), 8410)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day), 1))
	aoc.Run("part2", part1(aoc.Strings(day), 1000000-1))
}

type Points struct {
	start, end image.Point
}

func uniquePointPairs(points []image.Point) (pairs []Points) {
	rects := map[Points]struct{}{}

	for _, p := range points {
		for _, q := range points {
			if aoc.ComparePoints(p, q) == -1 {
				rects[Points{p, q}] = struct{}{}
			}
		}
	}

	pairs = maps.Keys(rects)

	slices.SortFunc(pairs, func(a, b Points) int {
		if a.start == b.start {
			return aoc.ComparePoints(a.end, b.end)
		}
		return aoc.ComparePoints(a.start, b.start)
	})

	return
}

func NumIsBetween(num, a, b int) bool {
	low, high := a, b
	if a > b {
		low, high = b, a
	}
	return num >= low && num <= high
}
