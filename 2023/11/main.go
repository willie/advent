package main

import (
	"fmt"
	"image"
	"slices"

	"github.com/willie/advent/aoc"
	"golang.org/x/exp/maps"
)

// The researcher has collected a bunch of data and compiled the data into a single giant image (your puzzle input). The image includes empty space (.) and galaxies (#). For example:

// ...#......
// .......#..
// #.........
// ..........
// ......#...
// .#........
// .........#
// ..........
// .......#..
// #...#.....
// The researcher is trying to figure out the sum of the lengths of the shortest path between every pair of galaxies. However, there's a catch: the universe expanded in the time it took the light from those galaxies to reach the observatory.

// Due to something involving gravitational effects, only some space expands. In fact, the result is that any rows or columns that contain no galaxies should all actually be twice as big.

// In the above example, three columns and two rows contain no galaxies:

//    v  v  v
//  ...#......
//  .......#..
//  #.........
// >..........<
//  ......#...
//  .#........
//  .........#
// >..........<
//  .......#..
//  #...#.....
//    ^  ^  ^
// These rows and columns need to be twice as big; the result of cosmic expansion therefore looks like this:

// ....#........
// .........#...
// #............
// .............
// .............
// ........#....
// .#...........
// ............#
// .............
// .............
// .........#...
// #....#.......
// Equipped with this expanded universe, the shortest path between every pair of galaxies can be found. It can help to assign every galaxy a unique number:

// ....1........
// .........2...
// 3............
// .............
// .............
// ........4....
// .5...........
// ............6
// .............
// .............
// .........7...
// 8....9.......
// In these 9 galaxies, there are 36 pairs. Only count each pair once; order within the pair doesn't matter. For each pair, find any shortest path between the two galaxies using only steps that move up, down, left, or right exactly one . or # at a time. (The shortest path between two galaxies is allowed to pass through another galaxy.)

// For example, here is one of the shortest paths between galaxies 5 and 9:

// ....1........
// .........2...
// 3............
// .............
// .............
// ........4....
// .5...........
// .##.........6
// ..##.........
// ...##........
// ....##...7...
// 8....9.......
// This path has length 9 because it takes a minimum of nine steps to get from galaxy 5 to galaxy 9 (the eight locations marked # plus the step onto galaxy 9 itself). Here are some other example shortest path lengths:

// Between galaxy 1 and galaxy 7: 15
// Between galaxy 3 and galaxy 6: 17
// Between galaxy 8 and galaxy 9: 5
// In this example, after expanding the universe, the sum of the shortest path between all 36 pairs of galaxies is 374.

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

	// fmt.Println(emptyRows, emptyColumns)
	fmt.Println("number of galaxies", len(g))

	galaxies := maps.Keys(g)
	// fmt.Println(galaxies)
	slices.SortFunc(galaxies, aoc.ComparePoints)
	// fmt.Println(galaxies)

	// paths :=
	paths2 := uniquePointPairs(galaxies)
	slices.SortFunc(paths2, func(a, b Points) int {
		if a.start == b.start {
			return aoc.ComparePoints(a.end, b.end)
		}
		return aoc.ComparePoints(a.start, b.start)
	})

	// fmt.Println("unique:", uniquePairs(galaxies), len(paths2))

	// iterate over the galaxies, find the manhattan distance to all the others
	for _, path := range paths2 {
		src, dest := path.start, path.end

		diff := 0

		for _, row := range emptyRows {
			// if row >= src.Y && row <= dest.Y {
			if NumIsBetween(row, src.Y, dest.Y) {
				diff += expansion
			}
		}
		for _, col := range emptyColumns {
			// if col >= src.X && col <= dest.X {
			if NumIsBetween(col, src.X, dest.X) {
				diff += expansion
			}
		}

		dist := aoc.ManhattanDistance(src.X, src.Y, dest.X, dest.Y)

		// fmt.Println(src, dest, g[src], "->", g[dest], dist, dist+diff)
		total += (dist + diff)
	}

	// for _, path := range paths {
	// 	src, dest := path.First, path.Second

	// 	dist := aoc.ManhattanDistance(src.X, src.Y, dest.X, dest.Y)
	// 	fmt.Println(src, dest, g[src], g[dest], dist)

	// 	for _, row := range emptyRows {
	// 		if row >= src.Y && row <= dest.Y {
	// 			fmt.Println("row", row, src, dest)
	// 			dist++
	// 		}
	// 	}
	// 	for _, col := range emptyColumns {
	// 		if col >= src.X && col <= dest.X {
	// 			fmt.Println("col", col, src, dest)
	// 			dist++
	// 		}
	// 	}
	// 	fmt.Println(src, dest, g[src], "->", g[dest], dist)
	// 	total += dist
	// }

	return
}

func part2(in []string) (total int) {
	return
}

const day = "https://adventofcode.com/2023/day/11"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test"), 1), 374)
	aoc.Test("test1", part1(aoc.Strings("test"), 100-1), 8410)
	// aoc.Test("test3", part2(aoc.Strings("test3")), 8)
	// aoc.Test("test4", part2(aoc.Strings("test4")), 10)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day), 1))
	aoc.Run("part2", part1(aoc.Strings(day), 1000000-1))
}

type PointPair struct {
	First, Second image.Point
}

func uniquePairs(points []image.Point) (pairs []PointPair) {
	n := len(points)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, PointPair{First: points[i], Second: points[j]})
		}
	}

	return
}

type Points struct {
	start, end image.Point
}

func uniquePointPairs(points []image.Point) (pairs []Points) {
	rects := map[Points]struct{}{}

	slices.SortFunc(points, aoc.ComparePoints)

	for _, p := range points {
		for _, q := range points {
			if aoc.ComparePoints(p, q) == -1 {
				rects[Points{p, q}] = struct{}{}

			}

		}
	}

	return maps.Keys(rects)
}

func NumIsBetween(num, a, b int) bool {
	low, high := a, b
	if a > b {
		low, high = b, a
	}
	return num >= low && num <= high
}
