package main

import (
	"fmt"
	"image"
	"math"

	"github.com/willie/advent/aoc"
)

func part1(name string) {
	grid := aoc.LoadRuneGrid(aoc.Strings(name))
	var start, end image.Point

	for pt, v := range grid {
		if v == 'S' {
			start = pt
			grid[pt] = 'a'
		} else if v == 'E' {
			end = pt
			grid[pt] = '{'
		}
	}

	path := aoc.BFS(start, end, func(in image.Point) (neighbors []image.Point) {
		return aoc.Filter(func(dest image.Point) bool { return grid[dest] <= grid[in]+1 }, grid.FourWayAdjacent(in))
	})

	shortest := math.MaxInt
	for _, a := range aoc.Contains(grid, 'a') {
		path := aoc.BFS(a, end, func(in image.Point) (neighbors []image.Point) {
			return aoc.Filter(func(dest image.Point) bool { return grid[dest] <= grid[in]+1 }, grid.FourWayAdjacent(in))
		})

		length := len(path)
		if length != 1 && length < shortest {
			shortest = length
		}
	}

	fmt.Println(len(path)-1, shortest-1)
}

func main() {
	part1("test.txt")
	part1("input.txt")
}
