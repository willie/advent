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

	fmt.Println(potential)

	// iterate over the paths of adjacent points, finding ones that have are adjacent to the current point, but not traversed yet, starting with potential[0] and ending with potential[1] and avoiding s

	current, _ := potential[0], potential[1]
	visited := map[image.Point]int{current: 0}

	previous := s
	fmt.Println(s, current, nextPoint(grid, current, s))

	for next := nextPoint(grid, current, previous); next != s; next = nextPoint(grid, current, previous) {
		previous, current = current, next
		visited[current] = len(visited)
	}

	fmt.Println(visited)
	return aoc.Max(maps.Values(visited)...)/2 + 1
}

// You quickly reach the farthest point of the loop, but the animal never emerges. Maybe its nest is within the area enclosed by the loop?

// To determine whether it's even worth taking the time to search for such a nest, you should calculate how many tiles are contained within the loop. For example:

// ...........
// .S-------7.
// .|F-----7|.
// .||.....||.
// .||.....||.
// .|L-7.F-J|.
// .|..|.|..|.
// .L--J.L--J.
// ...........
// The above loop encloses merely four tiles - the two pairs of . in the southwest and southeast (marked I below). The middle . tiles (marked O below) are not in the loop. Here is the same loop again with those regions marked:

// ...........
// .S-------7.
// .|F-----7|.
// .||OOOOO||.
// .||OOOOO||.
// .|L-7OF-J|.
// .|II|O|II|.
// .L--JOL--J.
// .....O.....
// In fact, there doesn't even need to be a full tile path to the outside for tiles to count as outside the loop - squeezing between pipes is also allowed! Here, I is still within the loop and O is still outside the loop:

// ..........
// .S------7.
// .|F----7|.
// .||OOOO||.
// .||OOOO||.
// .|L-7F-J|.
// .|II||II|.
// .L--JL--J.
// ..........
// In both of the above examples, 4 tiles are enclosed by the loop.

// Here's a larger example:

// .F----7F7F7F7F-7....
// .|F--7||||||||FJ....
// .||.FJ||||||||L7....
// FJL7L7LJLJ||LJ.L-7..
// L--J.L7...LJS7F-7L7.
// ....F-J..F7FJ|L7L7L7
// ....L7.F7||L7|.L7L7|
// .....|FJLJ|FJ|F7|.LJ
// ....FJL-7.||.||||...
// ....L---J.LJ.LJLJ...
// The above sketch has many random bits of ground, some of which are in the loop (I) and some of which are outside it (O):

// OF----7F7F7F7F-7OOOO
// O|F--7||||||||FJOOOO
// O||OFJ||||||||L7OOOO
// FJL7L7LJLJ||LJIL-7OO
// L--JOL7IIILJS7F-7L7O
// OOOOF-JIIF7FJ|L7L7L7
// OOOOL7IF7||L7|IL7L7|
// OOOOO|FJLJ|FJ|F7|OLJ
// OOOOFJL-7O||O||||OOO
// OOOOL---JOLJOLJLJOOO
// In this larger example, 8 tiles are enclosed by the loop.

// Any tile that isn't part of the main loop can count as being enclosed by the loop. Here's another example with many bits of junk pipe lying around that aren't connected to the main loop at all:

// FF7FSF7F7F7F7F7F---7
// L|LJ||||||||||||F--J
// FL-7LJLJ||||||LJL-77
// F--JF--7||LJLJ7F7FJ-
// L---JF-JLJ.||-FJLJJ7
// |F|F-JF---7F7-L7L|7|
// |FFJF7L7F-JF7|JL---7
// 7-L-JL7||F7|L7F-7F7|
// L.L7LFJ|||||FJL7||LJ
// L7JLJL-JLJLJL--JLJ.L
// Here are just the tiles that are enclosed by the loop marked with I:

// FF7FSF7F7F7F7F7F---7
// L|LJ||||||||||||F--J
// FL-7LJLJ||||||LJL-77
// F--JF--7||LJLJIF7FJ-
// L---JF-JLJIIIIFJLJJ7
// |F|F-JF---7IIIL7L|7|
// |FFJF7L7F-JF7IIL---7
// 7-L-JL7||F7|L7F-7F7|
// L.L7LFJ|||||FJL7||LJ
// L7JLJL-JLJLJL--JLJ.L
// In this last example, 10 tiles are enclosed by the loop.

// Figure out whether you have time to search for the nest by calculating the area within the loop. How many tiles are enclosed by the loop?

func part2(in []string) (total int) {
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

	fmt.Println(potential)

	// iterate over the paths of adjacent points, finding ones that have are adjacent to the current point, but not traversed yet, starting with potential[0] and ending with potential[1] and avoiding s

	current, _ := potential[0], potential[1]
	visited := map[image.Point]int{current: 0}

	previous := s
	fmt.Println(s, current, nextPoint(grid, current, s))

	for next := nextPoint(grid, current, previous); next != s; next = nextPoint(grid, current, previous) {
		previous, current = current, next
		visited[current] = len(visited)
	}

	// what tiles are inside the loop?
	visited[s] = 0
	bounds := aoc.Bounds(maps.Keys(visited))

	// count the number of "." in each row that are preceded by an odd number of loop tiles
	for y := bounds.Max.Y; y >= bounds.Min.Y; y-- {
		// tileCount := 0

		// find all of the visited tiles in this row and get the min and max x values
		visitedInRow := []int{}
		for _, pt := range maps.Keys(visited) {
			if pt.Y == y {
				visitedInRow = append(visitedInRow, pt.X)
			}
		}

		// for x := bounds.Min.X; x <= bounds.Max.X; x++ {

		borderCount := 0
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
					borderCount++
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

	grid.PrintYFlipped(".")

	return
}

// check if a point is inside a shape defined by a list of points
func inside(pt image.Point, shape []image.Point) (inside bool) {
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

	// aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}

// The pipes are arranged in a two-dimensional grid of tiles:

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
// Based on the acoustics of the animal's scurrying, you're confident the pipe that contains the animal is one large, continuous loop.

// For example, here is a square loop of pipe:

// .....
// .F-7.
// .|.|.
// .L-J.
// .....
// If the animal had entered this loop in the northwest corner, the sketch would instead look like this:

// .....
// .S-7.
// .|.|.
// .L-J.
// .....
// In the above diagram, the S tile is still a 90-degree F bend: you can tell because of how the adjacent pipes connect to it.

// Unfortunately, there are also many pipes that aren't connected to the loop! This sketch shows the same loop as above:

// -L|F7
// 7S-7|
// L|7||
// -L-J|
// L|-JF
// In the above diagram, you can still figure out which pipes form the main loop: they're the ones connected to S, pipes those pipes connect to, pipes those pipes connect to, and so on. Every pipe in the main loop connects to its two neighbors (including S, which will have exactly two pipes connecting to it, and which is assumed to connect back to those two pipes).

// Here is a sketch that contains a slightly more complex main loop:

// ..F7.
// .FJ|.
// SJ.L7
// |F--J
// LJ...
// Here's the same example sketch with the extra, non-main-loop pipe tiles also shown:

// 7-F7-
// .FJ|7
// SJLL7
// |F--J
// LJ.LJ
// If you want to get out ahead of the animal, you should find the tile in the loop that is farthest from the starting position. Because the animal is in the pipe, it doesn't make sense to measure this by direct distance. Instead, you need to find the tile that would take the longest number of steps along the loop to reach from the starting point - regardless of which way around the loop the animal went.

// In the first example with the square loop:

// .....
// .S-7.
// .|.|.
// .L-J.
// .....
// You can count the distance each tile in the loop is from the starting point like this:

// .....
// .012.
// .1.3.
// .234.
// .....
// In this example, the farthest point from the start is 4 steps away.

// Here's the more complex loop again:

// ..F7.
// .FJ|.
// SJ.L7
// |F--J
// LJ...
// Here are the distances for each tile on that loop:

// ..45.
// .236.
// 01.78
// 14567
// 23...
// Find the single giant loop starting at S. How many steps along the loop does it take to get from the starting position to the point farthest from the starting position?
