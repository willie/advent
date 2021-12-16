package main

import (
	"fmt"
	"image"

	"github.com/beefsack/go-astar"
	"github.com/willie/advent/aoc"
)

const (
	invalid = "-1"
)

type location struct {
	x, y int
	g    *aoc.Grid
}

func Location(g *aoc.Grid, x, y int) location {
	return location{g: g, x: x, y: y}
}

func (l *location) String() string {
	return fmt.Sprintf("{%d,%d}", l.x, l.y)
}

var (
	adjacent = []image.Point{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} // {1, 1}, {-1, -1}, {1, -1}, {-1, 1},
)

func (l location) PathNeighbors() (neighbors []astar.Pather) {
	for _, diff := range adjacent {
		if i := l.g.Get(l.x+diff.X, l.y+diff.Y, invalid); i != invalid {
			neighbors = append(neighbors, Location(l.g, l.x+diff.X, l.y+diff.Y))
		}
	}
	return
}

func (l location) PathNeighborCost(to astar.Pather) (cost float64) {
	t := to.(location)
	v := l.g.Get(t.x, t.y, invalid)
	return float64(aoc.AtoI(v))
}

func (l location) PathEstimatedCost(to astar.Pather) (cost float64) {
	t := to.(location)
	dist := aoc.ManhattanDistance(l.x, l.y, t.x, t.y)
	return float64(dist)
}

func part1(in []string) (result int) {
	g := aoc.NewGrid(in)

	start := Location(&g, 0, 0)
	end := Location(&g, g.Width()-1, g.Height()-1)

	_, distance, _ := astar.Path(start, end)

	return int(distance)
}

func part2(in string) (result int) {
	return
}

const day = "https://adventofcode.com/2021/day/15"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 40)
	// aoc.Test("test2", part2(aoc.String("test")), 195)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	// aoc.Run("part2", part2(aoc.String(day)))
}
