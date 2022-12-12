package main

import (
	"fmt"
	"image"
	"log"

	"github.com/dominikbraun/graph"
	"github.com/willie/advent/aoc"
)

var adjacent = []image.Point{{0, 1}, {-1, 0}, {1, 0}, {0, -1}}

func part1(name string) {
	grid := aoc.LoadStringGrid(aoc.Strings(name))

	g := graph.New(graph.StringHash)

	var S, E image.Point

	// set up the points
	for pt, v := range grid { // iterate grid
		g.AddVertex(pt.String())
		if v == "S" {
			S = pt
			grid[pt] = "a"
		} else if v == "E" {
			E = pt
			grid[pt] = "z"
		}
	}

	for pt1, val1 := range grid { // iterate grid
		for _, pt2 := range aoc.Map(pt1.Add, adjacent) { // get all adjacent
			if val2, ok := grid[pt2]; ok { // draw an edge
				if val2[0] <= val1[0]+1 {
					fmt.Println(val1, val2)

					g.AddEdge(pt1.String(), pt2.String())
				}
			}
		}
	}

	path, err := graph.ShortestPath(g, S.String(), E.String())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(path), path)
}

func main() {
	part1("test.txt")
	// part1("input.txt")

	// println("------")

	// part2("test.txt")
	// part2("input.txt")
}
