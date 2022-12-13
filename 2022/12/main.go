package main

import (
	"fmt"
	"image"
	"math"

	"github.com/beefsack/go-astar"
	"github.com/willie/advent/aoc"
)

var adjacent = []image.Point{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}

type location struct {
	pt   image.Point
	grid aoc.Grid2[string]
}

func (src location) PathNeighbors() (neighbors []astar.Pather) {
	for _, adj := range aoc.Map(src.pt.Add, adjacent) {
		if b, ok := src.grid[adj]; ok {
			if b[0] <= src.grid[src.pt][0]+1 {
				neighbors = append(neighbors, location{pt: adj, grid: src.grid})
			}
		}
	}
	return
}

func (l location) PathNeighborCost(to astar.Pather) (cost float64) {
	t := to.(location)
	if _, ok := l.grid[t.pt]; !ok {
		return 1
	}
	return
}

func (l location) PathEstimatedCost(to astar.Pather) (cost float64) {
	t := to.(location)
	dist := aoc.ManhattanDistance(l.pt.X, l.pt.Y, t.pt.X, t.pt.Y)
	return float64(dist)
}

func part1(name string) {
	grid := aoc.LoadStringGrid(aoc.Strings(name))

	var start, end image.Point

	// set up the points
	for pt, v := range grid { // iterate grid
		if v == "S" {
			start = pt
			grid[pt] = "a"
		} else if v == "E" {
			end = pt
			grid[pt] = "{"
		}
	}

	path := aoc.BFS(start, end, func(in image.Point) (neighbors []image.Point) {
		for _, dest := range aoc.Map(in.Add, adjacent) {
			if destVal, ok := grid[dest]; ok {
				if destVal[0] <= grid[in][0]+1 {
					neighbors = append(neighbors, dest)
				}
			}
		}

		return
	})

	shortest := math.MaxInt
	for _, a := range aoc.Contains(grid, "a") {
		path := aoc.BFS(a, end, func(in image.Point) (neighbors []image.Point) {
			for _, dest := range aoc.Map(in.Add, adjacent) {
				if destVal, ok := grid[dest]; ok {
					if destVal[0] <= grid[in][0]+1 {
						neighbors = append(neighbors, dest)
					}
				}
			}

			return
		})

		length := len(path) - 1
		if length != 0 && length < shortest {
			shortest = length
		}
	}

	// for src, v1 := range grid { // iterate grid
	// 	for _, dest := range aoc.Map(src.Add, adjacent) { // get all adjacent
	// 		if v2, ok := grid[dest]; ok { // draw an edge
	// 			if v2[0] <= v1[0]+1 {
	// 				// fmt.Println(val1, val2)

	// 				// g.AddEdge(pt1.String(), pt2.String())
	// 				g.AddBoth(pt1.String(), pt2.String())
	// 			}
	// 		}
	// 	}
	// }

	// g, _ = graph.TransitiveReduction(g)

	// path, err := graph.ShortestPath(g, S.String(), E.String())
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// file, _ := os.Create("./mygraph.gv")
	// _ = draw.DOT(g, file)

	fmt.Println(len(path)-1, shortest)
}

func main() {
	part1("test.txt")
	part1("input.txt")

	// println("------")

	// part2("test.txt")
	// part2("input.txt")
}
