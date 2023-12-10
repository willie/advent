package main

import (
	"fmt"
	"image"
	"slices"

	"github.com/willie/advent/aoc"
	"golang.org/x/exp/maps"
)

func nextPoints(grid aoc.Grid2[string], c image.Point) (next []image.Point) {
	var (
		north = c.Add(image.Point{0, -1})
		south = c.Add(image.Point{0, 1})
		east  = c.Add(image.Point{1, 0})
		west  = c.Add(image.Point{-1, 0})
	)

	switch grid[c] {
	case "|":
		return []image.Point{north, south}
	case "-":
		return []image.Point{east, west}
	case "L":
		return []image.Point{north, east}
	case "J":
		return []image.Point{north, west}
	case "7":
		return []image.Point{south, west}
	case "F":
		return []image.Point{south, east}
	}

	return
}

func nextPoint(grid aoc.Grid2[string], c image.Point, previous image.Point) (next image.Point) {
	for _, n := range nextPoints(grid, c) {
		if n != previous {
			return n
		}
	}

	return
}

func part1(in []string) (total int) {
	grid := aoc.LoadStringGrid(in)

	s := grid.Contains("S")[0]

	grid.PrintYFlipped(".")
	fmt.Println(s)

	// pathDistance := map[image.Point]int{}
	potential := []image.Point{}
	for _, p := range grid.FourWayAdjacent(s) {
		if slices.Contains(nextPoints(grid, p), s) {
			potential = append(potential, p)
		}
	}

	// fmt.Println(potential)

	// iterate over the paths of adjacent points, finding ones that have are adjacent to the current point, but not traversed yet, starting with potential[0] and ending with potential[1] and avoiding s

	current, _ := potential[0], potential[1]
	visited := map[image.Point]int{current: 0}

	previous := s

	for next := nextPoint(grid, current, previous); next != s; next = nextPoint(grid, current, previous) {
		previous, current = current, next
		visited[current] = len(visited)
	}

	// fmt.Println(visited)
	return aoc.Max(maps.Values(visited)...)/2 + 1
}

func part2(in []string) (total int) {
	grid := aoc.LoadStringGrid(in)

	s := grid.Contains("S")[0]

	// grid.PrintYFlipped(".")
	// fmt.Println(s)

	// pathDistance := map[image.Point]int{}
	potential := []image.Point{}
	for _, p := range grid.FourWayAdjacent(s) {
		if slices.Contains(nextPoints(grid, p), s) {
			potential = append(potential, p)
		}
	}

	fmt.Println(potential)

	// iterate over the paths of adjacent points, finding ones that have are adjacent to the current point, but not traversed yet, starting with potential[0] and ending with potential[1] and avoiding s

	current, _ := potential[0], potential[1]
	visited := map[image.Point]int{current: 0}

	previous := s
	// fmt.Println(s, current, nextPoint(grid, current, s))

	for next := nextPoint(grid, current, previous); next != s; next = nextPoint(grid, current, previous) {
		previous, current = current, next
		visited[current] = len(visited)
	}

	// what tiles are inside the loop?
	visited[s] = 0
	bounds := aoc.Bounds(maps.Keys(visited))

	for y := bounds.Max.Y; y >= bounds.Min.Y; y-- {

		// find all of the visited tiles in this row and get the min and max x values
		visitedInRow := []int{}
		for _, pt := range maps.Keys(visited) {
			if pt.Y == y {
				visitedInRow = append(visitedInRow, pt.X)
			}
		}

		inside := false
		prev := ""

		for x := aoc.Min(visitedInRow...); x < aoc.Max(visitedInRow...); x++ {
			pt := image.Pt(x, y)
			v := grid[pt]

			if _, ok := visited[pt]; ok {
				if v == "|" ||
					(prev == "F" && v == "J") ||
					(prev == "L" && v == "7") {
					inside = !inside
				} else if v == "-" {
					continue
				}

				prev = v

			} else {
				if inside {
					total++
					grid[pt] = " "
				} else {
					grid[pt] = "."
				}
			}
		}
	}

	// grid.PrintYFlipped(".")

	return
}

const day = "https://adventofcode.com/2023/day/10"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 8)
	aoc.Test("test2", part2(aoc.Strings("test2")), 4)
	aoc.Test("test3", part2(aoc.Strings("test3")), 8)
	aoc.Test("test4", part2(aoc.Strings("test4")), 10)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
